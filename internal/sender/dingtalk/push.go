package dingtalk

type PushBody interface {
	String()
}

type Pusher interface {
	Push(p PushBody) error
}
