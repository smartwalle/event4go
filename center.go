package event4go

import (
	"errors"
	"sync"
	"time"
	"unsafe"
)

var (
	ErrTimeout         = errors.New("event send timeout")
	ErrNotFoundHandler = errors.New("not found handler")
)

var instance *Center
var once sync.Once

func Default() *Center {
	once.Do(func() {
		instance = NewCenter()
	})
	return instance
}

type Center struct {
	mu         *sync.Mutex
	hChainList map[string]*handlerChain
}

func NewCenter() *Center {
	var center = &Center{}
	center.mu = &sync.Mutex{}
	center.hChainList = make(map[string]*handlerChain)
	return center
}

func (this *Center) Handle(name string, handler EventHandler) {
	if len(name) == 0 {
		return
	}

	if handler == nil {
		return
	}

	this.mu.Lock()
	defer this.mu.Unlock()

	var hChain, ok = this.hChainList[name]
	if ok == false {
		hChain = newHandlerChain()
		this.hChainList[name] = hChain
	}

	var handlerValue = *(*int)(unsafe.Pointer(&handler))
	for _, ob := range hChain.handlerList {
		if *(*int)(unsafe.Pointer(&ob)) == handlerValue {
			return
		}
	}

	hChain.handlerList = append(hChain.handlerList, handler)
}

func (this *Center) RemoveHandler(name string, handler EventHandler) {
	if len(name) == 0 {
		return
	}

	if handler == nil {
		return
	}

	this.mu.Lock()
	defer this.mu.Unlock()

	var hChain, ok = this.hChainList[name]
	if ok == false {
		return
	}

	var index = -1
	var observerValue = *(*int)(unsafe.Pointer(&handler))
	for i, ob := range hChain.handlerList {
		if *(*int)(unsafe.Pointer(&ob)) == observerValue {
			index = i
		}
	}

	if index >= 0 {
		hChain.handlerList = append(hChain.handlerList[:index], hChain.handlerList[index+1:]...)
	}

	if len(hChain.handlerList) == 0 {
		close(hChain.event)
		delete(this.hChainList, name)
	}
}

func (this *Center) RemoveHandlerWithName(name string) {
	if len(name) == 0 {
		return
	}

	this.mu.Lock()
	defer this.mu.Unlock()

	var hChain, ok = this.hChainList[name]
	if ok == false {
		return
	}

	close(hChain.event)
	hChain.handlerList = nil
	delete(this.hChainList, name)
}

func (this *Center) RemoveAllHandler() {
	this.mu.Lock()
	defer this.mu.Unlock()

	for key, hChain := range this.hChainList {
		close(hChain.event)
		hChain.handlerList = nil
		delete(this.hChainList, key)
	}
}

func (this *Center) Post(name string, userInfo interface{}) error {
	if len(name) == 0 {
		return ErrNotFoundHandler
	}
	var event = newEvent(name, userInfo)

	this.mu.Lock()
	var hChain = this.hChainList[name]
	this.mu.Unlock()

	if hChain != nil {
		select {
		case hChain.event <- event:
			return nil
		case <-time.After(time.Second * 3):
			return ErrTimeout
		}
	}
	return ErrNotFoundHandler
}
