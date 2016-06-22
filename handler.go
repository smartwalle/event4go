package notification

type NotificationHandler func(notification *Notification)

type handlerChain struct {
	notification chan *Notification
	handlerList  []NotificationHandler
}

func newHandlerChain() *handlerChain {
	var c = &handlerChain{}
	c.notification = make(chan *Notification)
	go c.run()
	return c
}

func (this *handlerChain) run() {
	for {
		select {
		case noti, ok := <-this.notification:
			if ok == false {
				return
			}

			for _, handler := range this.handlerList {
				handler(noti)
			}
		}
	}
}
