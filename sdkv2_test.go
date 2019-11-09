package CocosSDK

import (
	"testing"
)

func TestInitSdk2(t *testing.T) {
	InitSDK("123.56.98.47", 80, false)
	//t.Log(rpc.GetDynamicGlobalProperties())
}

func TestWallet(t *testing.T) {
	Wallet.ImportAccount("ximenyan1111", "xmlcpp123")
	Wallet.ImportAccount("gggg1", "12345678")
	Wallet.SetDefaultAccount("gggg1", "12345678")
}

func TestTransfer(t *testing.T) {
	Wallet.Transfer("ximenyan1111", "COCOS", "hi ---", 1.1)
}
