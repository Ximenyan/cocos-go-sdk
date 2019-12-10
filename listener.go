package CocosSDK

import (
	"CocosSDK/rpc"
	. "CocosSDK/type"
)

func Listener(msg string, hander func(r *Notice) error) {
	rpc.Client.Subscribe(msg, hander)
}
