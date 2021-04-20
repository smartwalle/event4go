package main

import (
	"fmt"
	"github.com/smartwalle/event4go"
	"time"
)

func main() {

	var be = time.Now()

	var a event4go.EventHandler
	a = func(noti *event4go.Event) {
		fmt.Println("a", noti)
		time.Sleep(time.Second * 6)
	}
	//var b = func(noti *event4go.Event) {
	//	fmt.Println("b", noti)
	//}

	event4go.Default().Handle("hi1", a)
	//event4go.Default().Handle("hi2", b)
	//event4go.Default().RemoveHandlerWithName("hi1")

	for i := 0; i < 1000000; i++ {
		event4go.Default().Post("hi1", fmt.Sprintf("%d", i))
		event4go.Default().Post("hi2", fmt.Sprintf("%d", i))
	}

	fmt.Println("dd", time.Now().Sub(be).Seconds())

	//event4go.Default().RemoveHandler("hi", a)
	//event4go.Default().RemoveHandler("hi", b)
}
