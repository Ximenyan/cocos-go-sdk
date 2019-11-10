package CocosSDK

import (
	//"cocos-go-sdk/wallet"
	"testing"
)

func TestInitSdk2(t *testing.T) {
	InitSDK("123.56.98.47", 80, false)
	//t.Log(rpc.GetDynamicGlobalProperties())
}

func TestWallet(t *testing.T) {
	//Wallet.ImportAccount("ximenyan1111", "xmlcpp123")
	//Wallet.ImportAccount("gggg1", "12345678")
}

func TestTransfer(t *testing.T) {
	//Wallet.SetDefaultAccount("gggg1", "12345678")

	//str, _ := wallet.DecodeMemo(Wallet.Default.GetMemoKey(), "COCOS6wm6Cqmz82xdxsaXMAiffTRaLDNAS4UAEmyGfTxWq5PSCT2ekw", "ba5adaf0feb5ce25254183a0cadb7fe9", 10078031519760515374)
	//t.Log(str)
	//Wallet.SetDefaultAccount("ximenyan1111", "xmlcpp123")
	//t.Log(Wallet.Default.GetActiveKey().ToHexString())
	//t.Log(GetVotingInfo())
	//Wallet.Transfer("ximenyan1111", "COCOS", "sss", 1)
	//Wallet.CreateAccount("ccccqwe123", "123345")
	//Wallet.UpgradeAccount("ximenyan1111")
	//Wallet.RegisterNhAssetCreator("ximenyan1111")
	//IssueToken("C1C1S", "ximenyan1111", 123.1)
	//CreateToken("C3C3S", 1000000, 3)
	//ReserveToken("C1C1S", 1)
}
