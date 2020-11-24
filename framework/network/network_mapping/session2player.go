package network_mapping

import (
	"github.com/liyuhuana/go-common/definition"
	"sync"
)

type SessionToPlayer struct {
	mux sync.RWMutex
	items map[int32]definition.PlayerId
}

var (
	once sync.Once
	inst *SessionToPlayer
)

func Inst() *SessionToPlayer {
	once.Do(func() {
		inst = &SessionToPlayer{
			mux:   sync.RWMutex{},
			items: make(map[int32]definition.PlayerId),
		}
	})
	return inst
}

func (this *SessionToPlayer) Get(sessionId int32) definition.PlayerId {
	this.mux.RLock()
	defer this.mux.RUnlock()

	if playerId, ok := this.items[sessionId]; ok {
		return playerId
	}
	return definition.EmptyPlayerId
}

func (this *SessionToPlayer) Add(sessionId int32, playerId definition.PlayerId) {
	this.mux.Lock()
	defer this.mux.Unlock()

	this.items[sessionId] = playerId
}

func (this *SessionToPlayer) Remove(sessionId int32) {
	this.mux.Lock()
	defer this.mux.Unlock()

	delete(this.items, sessionId)
}
