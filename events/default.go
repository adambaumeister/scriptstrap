package events

type Event interface {
	GetTag() string
	GetHost() string
	GetState() string
}
