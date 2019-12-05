package CocosSDK

import (
	//"encoding/json"
	//"CocosSDK/wallet"
	//"CocosSDK/type"
	//"CocosSDK/rpc"
	"testing"
)

func TestInitSdk2(t *testing.T) {
	InitSDK("test.cocosbcx.net", false, 80)
}

func TestWallet(t *testing.T) {
	//导入账户
	Wallet.ImportAccount("ggggxxx", "12345678")
	//设置默认账户
	Wallet.SetDefaultAccount("ggggxxx", "12345678")
}

func TestTransfer(t *testing.T) {
	t.Log(Wallet.Default.GetActiveKey().GetPublicKey().ToBase58String())
	t.Log(Wallet.Default.GetOwnerKey().GetPublicKey().ToBase58String())
	t.Log(Wallet.Transfer("test1", "COCOS", 1))
	t.Log(Wallet.Transfer("test1", "COCOS", 1, "xxxx"))
	t.Log(Wallet.TransferEncodeMemo("test1", "COCOS", 1, "xxxx"))
	//查询投票信息
	//vote_info := GetVotingInfo()
	//t.Log(vote_info)
	//Wallet.Transfer("ximenyan1111", "COCOS", 1, "xixixi")

	//Wallet.UpgradeAccount("ximenyan1111")
	//Wallet.RegisterNhAssetCreator("ximenyan1111")
	//CreateToken("C5C5S", 1000000, 3)
	//IssueToken("C4C4S", "gggg2", 1000)
	//ReserveToken("C4C4S", 1)
	//Wallet.RegisterNhAssetCreator("gggg2")
	//UpdateToken("C5C5S", 10000000, 7)
	//CreateWorldView("BCX")
	//CreateNhAsset("COCOS", "BCX", "gggg1", `{"name":"乾坤大挪移"}`)
	//SellNhAsset("ximenyan1111", "4.2.1", "便宜货...", COCOS_ID, COCOS_ID, 5, 100)
	//CancelNhAssetOrder("4.3.0")
	//FillNhAsset("4.3.1")
	//DeleteNhAsset("4.2.0")
	//ReviseContractByFile("contract.test12343", "./test.lua")
	//TransferNhAsset("gggg2", "4.2.2")
	//t.Log(wallet.CreateKey().GetPublicKey().ToBase58String())
	/*
		//质押gas
		hash, err := Pledgegas("gggg2", 100)
		t.Log(hash, err)
		//赎回
		hash, err = Pledgegas("gggg2", 0)
		t.Log(hash, err)
		//投票
		hash, err = Vote("1.5.6", 100)
		t.Log(hash, err)
		//投票赎回
		hash, err = Vote("1.5.6", 0)
		t.Log(hash, err)
		//查询可领取的冻结资产
		GetVestingBalances()
		//领取冻结资产
		hash, err = WithdrawVestingBalance("1.13.30")
		t.Log(hash, err)*/
}
