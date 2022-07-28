package sender

import cm "sail/common"

type Pusher interface {
	Push(m cm.OutMessage) error
}

func PushMessage(p Pusher, m cm.OutMessage) error {
	p.Push(m)
	return nil
}
