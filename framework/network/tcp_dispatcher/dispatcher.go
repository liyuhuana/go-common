package tcp_dispatcher

import (
	"github.com/liyuhuana/go-common/definition"
	"github.com/liyuhuana/go-common/framework/network/network_mapping"
	"github.com/liyuhuana/go-common/framework/network/network_tcp"
	"sync"
)

type DispatchFunc func(*network_tcp.Session, definition.PlayerId, []byte) []byte

type Dispatcher struct {
	mux sync.Mutex
	nMsgIdMap map[int32]int32
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
			nMsgIdMap:nil,
			nDispatchFunc: nil,
		}
	})
	return inst
}

func (d *Dispatcher) Init(f DispatchFunc, msgIdMap map[int32]int32) {
	d.nDispatchFunc = f
	d.nMsgIdMap = msgIdMap
}

func (d *Dispatcher) OnOpen(session *network_tcp.Session) {

}

func (d *Dispatcher) OnClose(session *network_tcp.Session, isForce bool) {

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

	rspMsgId := d.nMsgIdMap[msgId]
	rspData := d.nDispatchFunc(session, playerId, msgData)
	return rspMsgId, rspData
}
