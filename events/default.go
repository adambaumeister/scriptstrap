package events

import "github.com/adamb/scriptdeliver/events/manual"

type Event interface {
	GetTag() string
	GetHost() string
}

func RouteEvent(data interface{}) Event {
	switch v := data.(type) {
	case string:
		e := manual.ManualEvent{}
		e.In(v)
		return e
	default:
		return nil
	}
}
