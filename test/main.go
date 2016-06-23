package main

import (
	"github.com/smartwalle/notification"
	"fmt"
	"time"
)

func main() {

	var be = time.Now()

	var a notification.NotificationHandler
	a = func (noti *notification.Notification) {
		//fmt.Println("a", noti)
	}
	var b = func (noti *notification.Notification) {
		//fmt.Println("b", noti)
	}

	notification.DefaultCenter().AddObserver("hi1", a)
	notification.DefaultCenter().AddObserver("hi2", b)
	notification.DefaultCenter().RemoveObserverWithName("hi1")

	for i:=0; i<1000000; i++ {
		notification.DefaultCenter().PostNotification("hi1", fmt.Sprintf("%d", i))
		notification.DefaultCenter().PostNotification("hi2", fmt.Sprintf("%d", i))
	}



	fmt.Println("dd", time.Now().Sub(be).Seconds())

	//notification.DefaultCenter().RemoveObserver("hi", a)
	//notification.DefaultCenter().RemoveObserver("hi", b)
}