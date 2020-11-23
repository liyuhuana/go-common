package network_tcp

import (
	"log"
	"sync"
	"sync/atomic"
	"time"

	"github.com/liyuhuana/go-common/recover"
)

var (
	once sync.Once
	inst *Monitor
)

type Monitor struct {
	stop   chan bool
	reads  int32
	writes int32
}

func MonitorInst() *Monitor {
	once.Do(func() {
		inst = &Monitor{
			stop:   make(chan bool),
			reads:  0,
			writes: 0,
		}
	})
	return inst
}

func (this *Monitor) Start() {
	go func() {
		defer recover.Recover()

		log.Println("Monitor start running...")

		interval := time.Second * 5
		timer := time.NewTimer(interval)

		for {
			select {
			case <-timer.C:
				timer.Reset(interval)
				log.Printf("Read:[%d/s] write:[%d/s]\n", atomic.LoadInt32(&this.reads)/5, atomic.LoadInt32(&this.writes)/5)
				this.resetMonitor()
			case <-this.stop:
				log.Println("Monitor stop.")
				return
			}
		}
	}()
}

func (this *Monitor) IncrRead(cnt int32) {
	atomic.AddInt32(&this.reads, cnt)
}

func (this *Monitor) IncrWrite(cnt int32) {
	atomic.AddInt32(&this.writes, cnt)
}

func (this *Monitor) resetMonitor() {
	atomic.StoreInt32(&this.reads, 0)
	atomic.StoreInt32(&this.writes, 0)
}

func (this *Monitor) Stop() {
	this.stop<-true
}
