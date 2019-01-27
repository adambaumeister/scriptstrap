package manual

import "strings"

type ManualEvent struct {
	Host string
	Tag  string
}

func (m *ManualEvent) In(data string) error {
	var s string
	splt := strings.Split(s, ",")
	m.Host = splt[0]
	m.Tag = splt[1]

}
