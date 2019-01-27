package events

type Event interface {
	In(interface{})
}
