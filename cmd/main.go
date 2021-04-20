package main

import (
	"fmt"
	"github.com/smartwalle/event4go"
	"time"
)

func main() {

	var be = time.Now()

	var a = func(event *event4go.Event) {
		fmt.Println("handler 1", event)
	}
	var b = func(event *event4go.Event) {
		fmt.Println("handler 2", event)
	}

	event4go.Default().Handle("event1", a)
	event4go.Default().Handle("event2", b)

	for i := 0; i < 1000; i++ {
		event4go.Default().Post("event1", fmt.Sprintf("%d", i))
		event4go.Default().Post("event2", fmt.Sprintf("%d", i))
	}

	fmt.Println("time", time.Now().Sub(be).Seconds())

	//event4go.Default().RemoveHandler("hi", a)
	//event4go.Default().RemoveHandler("hi", b)
}
