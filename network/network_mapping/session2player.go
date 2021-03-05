package network_mapping

import (
	"github.com/liyuhuana/go-common/definition"
	"sync"
)

type SessionToUser struct {
	mux sync.RWMutex
	items map[int32]definition.UserId
}

var (
	once sync.Once
	inst *SessionToUser
)

func Inst() *SessionToUser {
	once.Do(func() {
		inst = &SessionToUser{
			mux:   sync.RWMutex{},
			items: make(map[int32]definition.UserId),
		}
	})
	return inst
}

func (this *SessionToUser) Get(sessionId int32) definition.UserId {
	this.mux.RLock()
	defer this.mux.RUnlock()

	if playerId, ok := this.items[sessionId]; ok {
		return playerId
	}
	return definition.EmptyUserId
}

func (this *SessionToUser) Add(sessionId int32, playerId definition.UserId) {
	this.mux.Lock()
	defer this.mux.Unlock()

	this.items[sessionId] = playerId
}

func (this *SessionToUser) Remove(sessionId int32) {
	this.mux.Lock()
	defer this.mux.Unlock()

	delete(this.items, sessionId)
}
