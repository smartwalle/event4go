package event4go

type EventHandler func(notification *Event)

type handlerChain struct {
	event       chan *Event
	handlerList []EventHandler
}

func newHandlerChain() *handlerChain {
	var c = &handlerChain{}
	c.event = make(chan *Event, 32)
	go c.run()
	return c
}

func (this *handlerChain) run() {
	for {
		select {
		case e, ok := <-this.event:
			if ok == false {
				return
			}

			for _, handler := range this.handlerList {
				handler(e)
			}
		}
	}
}
