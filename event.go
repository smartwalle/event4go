package event4go

type Event struct {
	Name     string
	UserInfo interface{}
}

func newEvent(name string, userInfo interface{}) *Event {
	var e = &Event{}
	e.Name = name
	e.UserInfo = userInfo
	return e
}
