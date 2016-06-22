package notification

import (
	"sync"
	"unsafe"
)

var instance *NotificationCenter
var once sync.Once

func DefaultCenter() *NotificationCenter {
	once.Do(func() {
		instance = NewNotificationCenter()
	})
	return instance
}

type NotificationCenter struct {
	mutex            *sync.Mutex
	handlerChainList map[string]*handlerChain
}

func NewNotificationCenter() *NotificationCenter {
	var center = &NotificationCenter{}
	center.mutex = &sync.Mutex{}
	center.handlerChainList = make(map[string]*handlerChain)
	return center
}

func (this *NotificationCenter) AddObserver(name string, handler NotificationHandler) {
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

func (this *NotificationCenter) RemoveObserver(name string, handler NotificationHandler) {
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
		close(handlerChain.notification)
		delete(this.handlerChainList, name)
	}
}

func (this *NotificationCenter) PostNotification(name string, userInfo interface{}) {
	if len(name) == 0 {
		return
	}
	var notification = NewNotification(name, userInfo)
	var handlerChain, ok = this.handlerChainList[name]

	if ok {
		handlerChain.notification <- notification
	}
}
