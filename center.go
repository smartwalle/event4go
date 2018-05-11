package event4go

import (
	"sync"
	"unsafe"
)

var instance *Center
var once sync.Once

func DefaultCenter() *Center {
	once.Do(func() {
		instance = NewCenter()
	})
	return instance
}

type Center struct {
	mutex            *sync.Mutex
	handlerChainList map[string]*handlerChain
}

func NewCenter() *Center {
	var center = &Center{}
	center.mutex = &sync.Mutex{}
	center.handlerChainList = make(map[string]*handlerChain)
	return center
}

func (this *Center) AddHandler(name string, handler EventHandler) {
	if len(name) == 0 {
		return
	}

	if handler == nil {
		return
	}

	this.mutex.Lock()
	defer this.mutex.Unlock()

	var handlerChain, ok = this.handlerChainList[name]
	if ok == false {
		handlerChain = newHandlerChain()
		this.handlerChainList[name] = handlerChain
	}

	var handlerValue = *(*int)(unsafe.Pointer(&handler))
	for _, ob := range handlerChain.handlerList {
		if *(*int)(unsafe.Pointer(&ob)) == handlerValue {
			return
		}
	}

	handlerChain.handlerList = append(handlerChain.handlerList, handler)
}

func (this *Center) RemoveHandler(name string, handler EventHandler) {
	if len(name) == 0 {
		return
	}

	if handler == nil {
		return
	}

	this.mutex.Lock()
	defer this.mutex.Unlock()

	var handlerChain, ok = this.handlerChainList[name]
	if ok == false {
		return
	}

	var index = -1
	var observerValue = *(*int)(unsafe.Pointer(&handler))
	for i, ob := range handlerChain.handlerList {
		if *(*int)(unsafe.Pointer(&ob)) == observerValue {
			index = i
		}
	}

	if index >= 0 {
		handlerChain.handlerList = append(handlerChain.handlerList[:index], handlerChain.handlerList[index+1:]...)
	}

	if len(handlerChain.handlerList) == 0 {
		close(handlerChain.event)
		delete(this.handlerChainList, name)
	}
}

func (this *Center) RemoveHandlerWithName(name string) {
	if len(name) == 0 {
		return
	}

	this.mutex.Lock()
	defer this.mutex.Unlock()

	var handlerChain, ok = this.handlerChainList[name]
	if ok == false {
		return
	}

	close(handlerChain.event)
	handlerChain.handlerList = nil
	delete(this.handlerChainList, name)
}

func (this *Center) RemoveAllHandler() {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	for key, handlerChain := range this.handlerChainList {
		close(handlerChain.event)
		handlerChain.handlerList = nil
		delete(this.handlerChainList, key)
	}
}

func (this *Center) PostEvent(name string, userInfo interface{}) {
	if len(name) == 0 {
		return
	}
	var notification = NewEvent(name, userInfo)
	var handlerChain, ok = this.handlerChainList[name]

	if ok {
		handlerChain.event <- notification
	}
}
