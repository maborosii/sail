package sender

import (
	"sail/global"
)

// var a Pusher
// var _ = a.(*dt.DingTalkPusher)
var PusherList = NewPusherList()

func init() {
	PusherList.RegisterPusher(global.PusherOfDingtalk)
}
