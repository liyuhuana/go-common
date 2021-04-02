package network_tcp

import (
	"net"
	"sync"
	"time"

	"go.uber.org/atomic"

	"github.com/liyuhuana/go-common/logger"
	"github.com/liyuhuana/go-common/recover"
)

const (
	broadcastConcurrentCapacity = 10
)

type Server struct {
	serverInfo ServerInfo

	sessions   sync.Map
	sessionSeq atomic.Int32
	sessionCnt atomic.Int32
	handler    IHandler
	listener   net.Listener

	limit           chan int
	keepAliveSignal chan int
}

func NewServer(host string, handler IHandler) *Server {
	server := &Server{
		serverInfo: ServerInfo{
			host: host,
		},
		sessions: sync.Map{},
		handler:  handler,
		limit:    make(chan int, broadcastConcurrentCapacity),
	}
	return server
}

func (this *Server) Start() {
	host := this.serverInfo.GetHost()
	listener, err := net.Listen("tcp4", host)
	if err != nil {
		logger.Error("server listen failed:", err)
		return
	}
	logger.Info("server start running, tcp address: [", host, "]")

	this.listener = listener

	go this.keepAlive()

	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Error(err)
			continue
		}

		this.handleConn(conn)
	}
}

func (this *Server) keepAlive() {
	defer recover.Recover()

	d := time.Second * 5
	t := time.NewTimer(d)

	for {
		select {
		case <-t.C:
			t.Reset(d)
			this.sessions.Range(func(k, v interface{}) bool {
				s := v.(*Session)
				lastRspElpase := s.elapsedSinceLastResponse()
				if lastRspElpase > 60 {
					s.Close(true)
				} else if lastRspElpase > 30 {
					s.Ping()
				}
				return true
			})
		case <-this.keepAliveSignal:
			logger.Info("Server keep alive stopped!")
		}
	}
}

func (this *Server) handleConn(conn net.Conn) {
	this.sessionSeq.Add(1)
	session := newSession(this.sessionSeq.Load(), this, conn)

	session.Start()
}

func (this *Server) OnOpen(session *Session) {
	this.sessions.Store(session.ID(), session)
	this.sessionCnt.Add(1)

	this.handler.OnOpen(session)

	logger.Info("Session connection established. sessionId:", session.ID(), " remoteIp:", session.GetRemoteIp(), " totalSession:", this.sessionCnt.Load())
}

func (this *Server) OnClose(session *Session, force bool) {
	this.sessions.Delete(session.ID())
	this.sessionCnt.Add(-1)

	this.handler.OnClose(session, force)

	logger.Info("Session connection closed. sessionId:", session.ID(), " remoteIp:", session.GetRemoteIp(), " totalSession:", this.sessionCnt.Load())
}

func (this *Server) OnRequest(session *Session, msgId int32, data []byte) (int32, int32, []byte) {
	return this.handler.OnRequest(session, msgId, data)
}

func (this *Server) OnPush(session *Session, msgId int32, data []byte) {
	this.handler.OnPush(session, msgId, data)
}

func (this *Server) Stop() {
	this.sessions.Range(func(k, v interface{}) bool {
		session := v.(*Session)
		session.Close(true)
		this.sessions.Delete(k)
		return true
	})
	this.sessionSeq.Store(0)
	this.sessionCnt.Store(0)

	err := this.listener.Close()
	if err != nil {
		logger.Error("Server listener close error:", err)
	}

	logger.Info("Server stopped!")
}
