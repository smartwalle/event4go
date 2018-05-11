package main

import (
	"fmt"
	"time"
	"github.com/smartwalle/event4go"
)

func main() {

	var be = time.Now()

	var a event4go.EventHandler
	a = func (noti *event4go.Event) {
		//fmt.Println("a", noti)
	}
	var b = func (noti *event4go.Event) {
		//fmt.Println("b", noti)
	}

	event4go.DefaultCenter().AddHandler("hi1", a)
	event4go.DefaultCenter().AddHandler("hi2", b)
	event4go.DefaultCenter().RemoveHandlerWithName("hi1")

	for i:=0; i<1000000; i++ {
		event4go.DefaultCenter().PostEvent("hi1", fmt.Sprintf("%d", i))
		event4go.DefaultCenter().PostEvent("hi2", fmt.Sprintf("%d", i))
	}



	fmt.Println("dd", time.Now().Sub(be).Seconds())

	//event4go.DefaultCenter().RemoveHandler("hi", a)
	//event4go.DefaultCenter().RemoveHandler("hi", b)
}