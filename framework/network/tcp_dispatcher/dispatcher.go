package tcp_dispatcher

import (
	"github.com/liyuhuana/go-common/definition"
	"github.com/liyuhuana/go-common/framework/network/network_mapping"
	"github.com/liyuhuana/go-common/framework/network/network_tcp"
	"sync"
)

type DispatchFunc func(*network_tcp.Session, int32, definition.PlayerId, []byte) (int32, []byte)

type Dispatcher struct {
	mux sync.Mutex
	nDispatchFunc DispatchFunc
}

var (
	once sync.Once
	inst *Dispatcher
)

func Inst() *Dispatcher {
	once.Do(func() {
		inst = &Dispatcher{
			mux:          sync.Mutex{},
			nDispatchFunc: nil,
		}
	})
	return inst
}

func (d *Dispatcher) Init(f DispatchFunc) {
	d.nDispatchFunc = f
}

func (d *Dispatcher) OnOpen(session *network_tcp.Session) {

}

func (d *Dispatcher) OnClose(session *network_tcp.Session, isForce bool) {
	network_mapping.Inst().Remove(session.ID())
}

func (d *Dispatcher) OnPush(session *network_tcp.Session, msgId int32, msgData []byte) {
	d.onMessage(session, msgId, msgData)
}

func (d *Dispatcher) OnRequest(session *network_tcp.Session, msgId int32, msgData []byte) (int32, []byte) {
	msgId, rsp := d.onMessage(session, msgId, msgData)
	return msgId, rsp
}

func (d *Dispatcher) onMessage(session *network_tcp.Session, msgId int32, msgData[] byte) (int32, []byte) {
	playerId := network_mapping.Inst().Get(session.ID())
	if playerId.IsEmpty() {
		return 0, nil
	}

	rspMsgId, rspData := d.nDispatchFunc(session, msgId, playerId, msgData)
	return rspMsgId, rspData
}
