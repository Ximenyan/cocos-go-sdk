package CocosSDK

import (
	//"encoding/json"
	//"CocosSDK/wallet"
	//. "CocosSDK/type"
	"testing"
)

func TestInitSdk2(t *testing.T) {
	InitSDK("123.56.98.47", 80, false)
}

func TestWallet(t *testing.T) {
	Wallet.ImportAccount("ximenyan1111", "xmlcpp123")
	Wallet.ImportAccount("gggg1", "12345678")
=======
}

func TestWallet(t *testing.T) {
	//导入账户
	Wallet.ImportAccount("gggg1", "12345678")
	//设置默认账户
>>>>>>> master
	Wallet.SetDefaultAccount("gggg1", "12345678")
}

func TestTransfer(t *testing.T) {
<<<<<<< HEAD
	Wallet.Transfer("ximenyan1111", "COCOS", "hi ---", 1.1)
=======

	//查询投票信息
	//vote_info := GetVotingInfo()
	//t.Log(vote_info)
	//Wallet.Transfer("ximenyan1111", "COCOS", 1, "xixixi")
	//Wallet.CreateAccount("ccccqwe123", "123345")
	//Wallet.UpgradeAccount("ximenyan1111")
	//Wallet.RegisterNhAssetCreator("ximenyan1111")
	//CreateToken("C5C5S", 1000000, 3)
	//IssueToken("C4C4S", "gggg2", 1000)
	//ReserveToken("C4C4S", 1)
	//Wallet.RegisterNhAssetCreator("gggg2")
	//UpdateToken("C5C5S", 10000000, 3)
	//CreateWorldView("BCX")
	//CreateNhAsset("COCOS", "BCX", "gggg1", `{"name":"乾坤大挪移"}`)
	//SellNhAsset("ximenyan1111", "4.2.1", "便宜货...", COCOS_ID, COCOS_ID, 5, 100)
	//CancelNhAssetOrder("4.3.0")
	//FillNhAsset("4.3.1")
	//DeleteNhAsset("4.2.0")
	//ReviseContractByFile("contract.test12343", "./test.lua")
	//TransferNhAsset("gggg2", "4.2.2")

	//质押gas
	Pledgegas("gggg2", 100)
	//赎回
	Pledgegas("gggg2", 0)
	//投票
	err := Vote("1.5.6", 10000)
	//投票赎回
	err = Vote("1.5.6", 0)
	//查询可领取的冻结资产
	GetVestingBalances()
	//领取冻结资产
	err = WithdrawVestingBalance("1.13.30")
	t.Log(err)
>>>>>>> master
}
