package manual

import (
	"github.com/adamb/scriptdeliver/events/base"
	"strings"
)

type ManualEvent struct {
	base.EventBase
}

func (m *ManualEvent) In(data string) error {
	splt := strings.Split(data, ",")
	m.Host = splt[0]
	m.Tag = splt[1]

	return nil
}
