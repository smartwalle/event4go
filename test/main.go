package main

import (
	"github.com/smartwalle/notification"
	"fmt"
	"time"
)

func main() {
	var a notification.NotificationHandler
	a = func (noti *notification.Notification) {
		fmt.Println("a", noti)
		notification.DefaultCenter().RemoveObserver("hi", a)
	}
	var b = func (noti *notification.Notification) {
		fmt.Println("b", noti)
	}

	notification.DefaultCenter().AddObserver("hi", a)
	notification.DefaultCenter().AddObserver("hi", b)

	notification.DefaultCenter().PostNotification("hi", "1")
	notification.DefaultCenter().PostNotification("hi", "2")

	time.Sleep(time.Second * 2)

	notification.DefaultCenter().RemoveObserver("hi", a)
	notification.DefaultCenter().RemoveObserver("hi", b)
}