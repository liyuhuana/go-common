package network_tcp

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"sync"
	"time"

	"sync/atomic"

	"github.com/liyuhuana/go-common/common_logger"
	"github.com/liyuhuana/go-common/definition"
	"github.com/liyuhuana/go-common/recover"
)

type Session struct {
	id     int32
	conn   net.Conn
	server *Server

	bodyLen uint32
	closed  int32 // 0:open 1:closed
	rspTime int64

	reqSeed int32
	reqPool sync.Map
}

func newSession(id int32, server *Server, conn net.Conn) *Session {
	s := &Session{
		id:      id,
		server:  server,
		conn:    conn,
		rspTime: time.Now().Unix(),
	}
	return s
}

func (this *Session) ID() int32 {
	return this.id
}

func (this *Session) GetRemoteIp() string {
	return this.conn.RemoteAddr().String()
}

func (this *Session) GetServer() *Server {
	return this.server
}

func (this *Session) Start() {
	common_logger.LogInfo("session connection established. sessionId:", this.ID())

	this.GetServer().OnOpen(this)

	go this.scan()
}

func (this *Session) scan() {
	defer recover.Recover()

	input := bufio.NewScanner(this.conn)
	input.Split(this.split)

	for input.Scan() {
		this.dispatch(input.Bytes())
	}

	this.Close(false)
}

func (this *Session) split(data []byte, atEOF bool) (advance int, token []byte, err error) {
	dataLen := len(data)
	offset := 0
	if dataLen == 0 {
		return 0, nil, nil
	}

	if atEOF {
		return dataLen, nil, nil
	}

	if this.bodyLen == 0 {
		uint32Size := definition.UInt32ByteLen.Int()
		if dataLen < uint32Size {
			return 0, nil, nil
		}

		this.bodyLen = binary.LittleEndian.Uint32(data[offset:uint32Size])
		dataLen -= uint32Size
		offset += uint32Size

		if dataLen < int(this.bodyLen) {
			if offset <= uint32Size {
				return uint32Size, nil, nil
			} else {
				return 2, nil, nil
			}
		}
	} else if dataLen < int(this.bodyLen) {
		return 0, nil, nil
	}
	advance = int(this.bodyLen) + offset
	this.bodyLen = 0
	return advance, data[offset:advance], nil
}

func (this *Session) dispatch(data []byte) {
	atomic.StoreInt64(&this.rspTime, time.Now().Unix())

	if len(data) < 1 {
		return
	}

	reader := bytes.NewBuffer(data)
	// read pattern
	pattern, err := reader.ReadByte()
	if err != nil {
		this.Close(false)
		common_logger.LogError(err)
		return
	}

	MonitorInst().IncrRead(1)

	left := len(data) - 1
	switch Pattern(pattern) {
	case Push:
		this.onPush(reader, left)
	case Request:
		this.onRequest(reader, left)
	case Response:
		this.onResponse(reader, left)
	case Ping:
		this.onPing(reader)
	case Pong:
		this.onPong(reader)
	case Sub:
	case Unsub:
	case Pub:
	}
}

func (this *Session) onPush(reader *bytes.Buffer, left int) {
	var msgId int32
	err := binary.Read(reader, binary.LittleEndian, &msgId)
	if err != nil {
		common_logger.LogError(err)
		this.Close(false)
		return
	}

	left -= 2
	body := make([]byte, left)
	n, err := reader.Read(body)
	if n != left || err != nil {
		this.Close(false)
		common_logger.LogError("session onPush exception, readByteLength:", n, "leftBuffLength:", left,
			"error:", err)
		return
	}

	this.server.OnPush(this, msgId, body)
}

func (this *Session) onRequest(reader *bytes.Buffer, left int) {
	var msgId int32
	err := binary.Read(reader, binary.LittleEndian, &msgId)
	if err != nil {
		common_logger.LogError(err)
		this.Close(false)
		return
	}

	left -= definition.Int32ByteLen.Int()
	body := make([]byte, left)
	n, err := reader.Read(body)
	if n != left || err != nil {
		this.Close(false)
		common_logger.LogError("session onRequest exception, readByteLength:", n, "leftBuffLength:", left,
			"error:", err)
		return
	}

	rspMsgId, result, rspData := this.server.OnRequest(this, msgId, body)
	// start response
	err = this.response(rspMsgId, result, rspData)
	if err != nil {
		common_logger.LogError("Session response error:", err)
	}
}

func (this *Session) onResponse(reader *bytes.Buffer, left int) {
	var serial uint16
	var en int16
	err := binary.Read(reader, binary.LittleEndian, &serial)
	if err != nil {
		common_logger.LogError(err)
		this.Close(false)
		return
	}

	left -= 2
	err = binary.Read(reader, binary.LittleEndian, &en)
	if err != nil {
		common_logger.LogError(err)
		this.Close(false)
		return
	}

	left -= 2
	body := make([]byte, left)
	n, err := reader.Read(body)
	if n != left || err != nil {
		this.Close(false)
		common_logger.LogError("session onResponse exception, readByteLength:", n, "leftBuffLength:", left,
			"error:", err)
		return
	}
}

func (this *Session) onPing(reader *bytes.Buffer) {
	serial, err := reader.ReadByte()
	if err != nil {
		this.Close(false)
		common_logger.LogError(err)
		return
	}

	err = this.pong(serial)
	if err != nil {
		common_logger.LogError(err)
	}
}

func (this *Session) onPong(reader *bytes.Buffer) {

}

func (this *Session) Write(data []byte) error {
	if this.IsClosed() {
		return fmt.Errorf("session %d already closed", this.ID())
	}

	writeTimeout := time.Second * 3
	err := this.conn.SetWriteDeadline(time.Now().Add(writeTimeout))
	if err != nil {
		this.Close(false)
		return err
	}

	n, err := this.conn.Write(data)
	dataLen := len(data)
	if n != dataLen {
		return fmt.Errorf("session write error => write:%d expected:%d", n, dataLen)
	}
	if err != nil {
		this.Close(false)
		return err
	}

	MonitorInst().IncrWrite(1)
	return err
}

func (this *Session) response(msgId, result int32, msgData []byte) error {
	buf := new(bytes.Buffer)
	int32Size := definition.UInt32ByteLen.Int()
	err := binary.Write(buf, binary.LittleEndian, uint32(1+int32Size+int32Size+len(msgData)))
	if err != nil {
		return err
	}

	//write msg type
	err = buf.WriteByte(byte(Response))
	if err != nil {
		return err
	}

	// write msgId
	err = binary.Write(buf, binary.LittleEndian, msgId)
	if err != nil {
		return err
	}

	// write result
	err = binary.Write(buf, binary.LittleEndian, result)
	if err != nil {
		return err
	}

	// write msgData
	if len(msgData) > 0 {
		buf.Write(msgData)
	}

	rspData := buf.Bytes()
	_, err = this.conn.Write(rspData)
	if err != nil {
		return err
	}
	return nil
}

func (this *Session) Ping() error {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, uint32(1+1))
	if err != nil {
		return err
	}
	buf.WriteByte(byte(Ping))
	buf.WriteByte(0)

	err = this.Write(buf.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func (this *Session) pong(serial byte) error {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, uint32(1+1))
	if err != nil {
		return err
	}
	buf.WriteByte(byte(Pong))
	buf.WriteByte(serial)

	err = this.Write(buf.Bytes())
	return err
}

func (this *Session) elapsedSinceLastResponse() int64 {
	return time.Now().Unix() - atomic.LoadInt64(&this.rspTime)
}

func (this *Session) Stop() {
	this.Close(true)
}

func (this *Session) IsClosed() bool {
	return this == nil || atomic.LoadInt32(&this.closed) != 0
}

func (this *Session) Close(force bool) {
	if !atomic.CompareAndSwapInt32(&this.closed, 0, 1) {
		return
	}

	this.conn.Close()

	common_logger.LogInfo("session closed. sessionId:", this.id)
	this.server.OnClose(this, force)
}
