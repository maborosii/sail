package sender

type Message interface {
	String()
}

type Pusher interface {
	Push(m *Message) error
}
