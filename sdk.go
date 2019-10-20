package CocosSDK

import (
	"cocos-go-sdk/chain"
	"cocos-go-sdk/rpc"
	"cocos-go-sdk/wallet"
	"log"
	"sync"
)

var once *sync.Once = &sync.Once{}
var Wallet *wallet.Wallet
var Chain *chain.Chain

/*
*初始化SDK
 */
func InitSDK(host string, port int, use_ssl bool) {
	once.Do(
		func() {
			defer func() {
				if err := recover(); err != nil {
					log.Panicln("SDK Init Error:", err)
				}
			}()
			if err := rpc.InitClient(host, port, use_ssl); err != nil {
				log.Panicln("SDK Init Error:", err)
			}
			chain.InitChain()
			Chain = chain.CocosBCXChain
			Wallet = wallet.CreateWallet()
		})
}
