package notification

type Notification struct {
	Name     string
	UserInfo interface{}
}

func NewNotification(name string, userInfo interface{}) *Notification {
	var notification = &Notification{}
	notification.Name = name
	notification.UserInfo = userInfo
	return notification
}
