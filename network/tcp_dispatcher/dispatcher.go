package tcp_dispatcher

import (
	"sync"

	"github.com/liyuhuana/go-common/definition"
	"github.com/liyuhuana/go-common/network/network_mapping"
	"github.com/liyuhuana/go-common/network/network_tcp"
)


type IHandler interface {
	OnOpen(*network_tcp.Session)
	OnClose(definition.PlayerId)
	OnMessage(*network_tcp.Session, int32, definition.PlayerId, []byte) (int32, int32, []byte)
}

type Dispatcher struct {
	mux sync.Mutex
	nHandler IHandler
}

var (
	once sync.Once
	inst *Dispatcher
)

func Inst() *Dispatcher {
	once.Do(func() {
		inst = &Dispatcher{
			mux:          sync.Mutex{},
		}
	})
	return inst
}

func (d *Dispatcher) Init(h IHandler) {
	d.nHandler = h
}

func (d *Dispatcher) OnOpen(session *network_tcp.Session) {
	d.nHandler.OnOpen(session)
}

func (d *Dispatcher) OnClose(session *network_tcp.Session, isForce bool) {
	playerId := network_mapping.Inst().Get(session.ID())
	d.nHandler.OnClose(playerId)
	network_mapping.Inst().Remove(session.ID())
}

func (d *Dispatcher) OnPush(session *network_tcp.Session, msgId int32, msgData []byte) {
	d.onMessage(session, msgId, msgData)
}

func (d *Dispatcher) OnRequest(session *network_tcp.Session, msgId int32, msgData []byte) (int32, int32, []byte) {
	msgId, result, rsp := d.onMessage(session, msgId, msgData)
	return msgId, result, rsp
}

func (d *Dispatcher) onMessage(session *network_tcp.Session, msgId int32, msgData[] byte) (int32, int32, []byte) {
	playerId := network_mapping.Inst().Get(session.ID())
	rspMsgId, result, rspData := d.nHandler.OnMessage(session, msgId, playerId, msgData)
	return rspMsgId, result, rspData
}
