package CocosSDK

import (
	"cocos-go-sdk/rpc"
	. "cocos-go-sdk/type"
)

func Listener(msg string, hander func(r *Notice) error) {
	rpc.Client.Subscribe(msg, hander)
}
