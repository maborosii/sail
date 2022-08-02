package sender

import (
	"fmt"
	"sail/global"
)

// var a Pusher
// var _ = a.(*dt.DingTalkPusher)
var PusherList = NewPusherList()

func init() {
	fmt.Println("init sender...")
	PusherList.RegisterPusher(global.PusherOfDingtalk)
}
