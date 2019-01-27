package channels

type Channel interface {
	Open() Channel
	RunScript([]byte)
}
