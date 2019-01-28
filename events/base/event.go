package base

type EventBase struct {
	Host string
	Tag  string
}

func (e EventBase) GetHost() string {
	return e.Host
}

func (e EventBase) GetTag() string {
	return e.Tag
}
