package CocosSDK

import (
	"CocosSDK/rpc"
	"encoding/json"
	"testing"
)

func TestInitSdk(t *testing.T) {
	InitSDK("123.56.98.47", 80, false)
	t.Log(rpc.GetDynamicGlobalProperties())
}

func TestTxsForAddress(t *testing.T) {
	txs, err := TxsForAddress("gggg1", 10, "263c47271171a5f99c839475f232a742074f848ddaba558e1b151106bf8dfbd1")
	byte_s, err := json.Marshal(txs)
	if err == nil {
		t.Log(string(byte_s))
	}
}

func TestPuk2Addr(t *testing.T) {
	t.Log(PublicToAddress("COCOS6wm6Cqmz82xdxsaXMAiffTRaLDNAS4UAEmyGfTxWq5PSCT2ekw"))
}

func TestAddr2Puk(t *testing.T) {
	t.Log(AddressToPublic("ximenyan1111"))
}

func TestGetBlockCount(t *testing.T) {
	t.Log(Getblockcount())
}

func TestGetrawmempool(t *testing.T) {
	t.Log(Getrawmempool())
}

func TestGetblocktxs(t *testing.T) {
	txs, err := Getblocktxs(77559)
	byte_s, err := json.Marshal(txs)
	if err == nil {
		t.Log(string(byte_s))
	}
}

func TestBalanceForAddress(t *testing.T) {
	balances := BalanceForAddress("ximenyan1111")
	byte_s, err := json.Marshal(balances)
	if err == nil {
		t.Log(string(byte_s))
	}
}

func TestDeserializeTransactions(t *testing.T) {
	sign_tx, _ := DeserializeTransactions("c1ac4bb7bd7d94874a1cb98b39a8a582421d03d022dfa4be8c70567076e03ad0f83b7f4d06b01d3ec85d01001a0000000000000016000000000000008096980000000000040000000000000001030ed1f4745aeb7194e1eea53bf6c4a217ba3b8f7d63ebad2e22543b99469bb03203d53f078f6ea92d7d33a06bf0e23569e376baf516ed0f5efe9a1b714be5f031d16a23d583d67366d710661cd4569de5a081559c97e382a360700000")
	byte_s, err := json.Marshal(sign_tx)
	if err == nil {
		t.Log(string(byte_s))
	}
}

func TestSignTransaction(t *testing.T) {
	tx, err := SignTransaction("c1ac4bb7bd7d94874a1cb98b39a8a582421d03d022dfa4be8c70567076e03ad0486711d2c551899dc85d010016000000000000001a00000000000000a08601000000000000000000000000000103d53f078f6ea92d7d33a06bf0e23569e376baf516ed0f5efe9a1b714be5f031d1030ed1f4745aeb7194e1eea53bf6c4a217ba3b8f7d63ebad2e22543b99469bb032b4d412ed0c8e38561077883f0dfb4c3f8e1068c92ef3e9653f0000",
		[]string{"202c76ab413de66315922a95c65b0dc77073bf1f9a7e809b0aa51db9f1592e359c2de34ed115c039d356ca573e0d4dc818a258acfc0af48c44c6e4c8d2c9d57508"})

	byte_s, err := json.Marshal(tx)
	if err == nil {
		t.Log(string(byte_s))
	}
}

func TestGetTransaction(t *testing.T) {
	tx, err := GetTransaction("0b40101202a469ae6700c5eca2512cb2ecb8fcd309949410240af465dc67d143")
	t.Log(err)
	byte_s, err := json.Marshal(tx)
	if err == nil {
		t.Log(string(byte_s))
	}
}

func TestBuildTransaction(t *testing.T) {
	hex_str, err := BuildTransaction("gggg2", "ximenyan1111", 1, "COCOS")
	t.Log(err)
	t.Log(hex_str)
}

func TestDeserialize(t *testing.T) {
	var hex_str string = "c1ac4bb7bd7d94874a1cb98b39a8a582421d03d022dfa4be8c70567076e03ad0d0b340b1ff472a46c95d010016000000000000001a00000000000000a08601000000000000000000000000000103d53f078f6ea92d7d33a06bf0e23569e376baf516ed0f5efe9a1b714be5f031d1030ed1f4745aeb7194e1eea53bf6c4a217ba3b8f7d63ebad2e22543b99469bb0326882e68e442706441093f14bfaecd9e70710b9053d2d27324c0000"
	tx, err := Deserialize(hex_str)
	byte_s, err := json.Marshal(tx)
	if err == nil {
		t.Log(string(byte_s))
	}
}

func TestUnsignedTxHash(t *testing.T) {
	hash := UnsignedTxHash("c1ac4bb7bd7d94874a1cb98b39a8a582421d03d022dfa4be8c70567076e03ad008bcc75d6e5f01f4cf5d010016000000000000001a00000000000000a08601000000000000000000000000000103d53f078f6ea92d7d33a06bf0e23569e376baf516ed0f5efe9a1b714be5f031d1030ed1f4745aeb7194e1eea53bf6c4a217ba3b8f7d63ebad2e22543b99469bb0326e614f3b308beaa01081355d8f9325a76997c83b2dfb17652e0000")
	if hash == "263c47271171a5f99c839475f232a742074f848ddaba558e1b151106bf8dfbd1" {
		t.Log("Test Unsigned Tx Hash success!!")
	} else {
		t.Error("Test Unsigned Tx Hash Error!")
	}
}
