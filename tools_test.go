package CocosSDK

import (
	"encoding/json"
	//"encoding/json"
	//"encoding/json"
	//"encoding/json"
	//"encoding/json"
	//	"cocos-go-sdk/chain"

	//"cocos-go-sdk/rpc"
	//"cocos-go-sdk/type"
	//	"encoding/json"
	"testing"
)

//000000000001030ed1f4745aeb7194e1eea53bf6c4a217ba3b8f7d63ebad2e22543b99469bb03202b8764675d8a40f416ba1571bd67e2b6b61c0adf68446af8b3b878efeeccb6e0ee8cda626af16a6dc106094ec3577a9e705cdaf5ed006f9397b0000
//7d89b84f22af0b150780a2b121aa6c715b19261c8b7fe0fda3a564574ed7d3e9ed3b9ae2f85f1a1ac35d0100a2510000000000000097d505c7e205b0ad0100000000000001030ed1f4745aeb7194e1eea53bf6c4a217ba3b8f7d63ebad2e22543b99469bb03202b8764675d8a40f416ba1571bd67e2b6b61c0adf68446af8b3b878efeeccb6e0e9ce14683d52d68191031f641c2a14dfb96ed226bc5ba2783000000

/*
func TestInitSdk(t *testing.T) {
	InitSDK("47.93.62.96", 8049, false)
	t.Log(rpc.GetDynamicGlobalProperties())
}

func TestTxsForAddress(t *testing.T) {
	txs, err := TxsForAddress("gggg1")
	t.Log(err)
	byte_s, err := json.Marshal(txs)
	if err == nil {
		t.Log(string(byte_s))
	}
}

*/

/*


func TestGetTransaction(t *testing.T) {
	tx, err := GetTransaction("08b8b4890d8cf08135dee370dc2f146c84ba6868ab9fac6d9dc0c9e1ac4a3fa3")
	t.Log(err)
	byte_s, err := json.Marshal(tx)
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
*/

/*
func TestBuildTransaction(t *testing.T) {
	hex_str, err := BuildTransaction("gggg1", "ximenyan1111", 1.1, "C0C0S")
	t.Log(err)
	t.Log(hex_str)
}
func TestDeserializeTransactions(t *testing.T) {
	sign_tx, _ := DeserializeTransactions("c1ac4bb7bd7d94874a1cb98b39a8a582421d03d022dfa4be8c70567076e03ad0f83b7f4d06b01d3ec85d01001a0000000000000016000000000000008096980000000000040000000000000001030ed1f4745aeb7194e1eea53bf6c4a217ba3b8f7d63ebad2e22543b99469bb03203d53f078f6ea92d7d33a06bf0e23569e376baf516ed0f5efe9a1b714be5f031d16a23d583d67366d710661cd4569de5a081559c97e382a360700000")
	byte_s, err := json.Marshal(sign_tx)
	if err == nil {
		t.Log(string(byte_s))
	}
}*/

/*




func TestBuildTransaction(t *testing.T) {
	hex_str, err := BuildTransaction("gggg1", "ximenyan1111", 1, "COCOS")
	t.Log(err)
	t.Log(hex_str)
}
2019/11/10 22:31:49 t.Unix() 1573425109
2019/11/10 22:31:49 uint64(t.Unix() 1573425109
2019/11/10 22:31:49 expiration_data [213 143 200 93]

*/

func TestDeserialize(t *testing.T) {
	var hex_str string = "c1ac4bb7bd7d94874a1cb98b39a8a582421d03d022dfa4be8c70567076e03ad0486711d2c551899dc85d010016000000000000001a00000000000000a08601000000000000000000000000000103d53f078f6ea92d7d33a06bf0e23569e376baf516ed0f5efe9a1b714be5f031d1030ed1f4745aeb7194e1eea53bf6c4a217ba3b8f7d63ebad2e22543b99469bb032b4d412ed0c8e38561077883f0dfb4c3f8e1068c92ef3e9653f0000"
	tx, err := Deserialize(hex_str)
	byte_s, err := json.Marshal(tx)
	if err == nil {
		t.Log(string(byte_s))
	}
}
func TestSignTransaction(t *testing.T) {
	tx, err := SignTransaction("c1ac4bb7bd7d94874a1cb98b39a8a582421d03d022dfa4be8c70567076e03ad0486711d2c551899dc85d010016000000000000001a00000000000000a08601000000000000000000000000000103d53f078f6ea92d7d33a06bf0e23569e376baf516ed0f5efe9a1b714be5f031d1030ed1f4745aeb7194e1eea53bf6c4a217ba3b8f7d63ebad2e22543b99469bb032b4d412ed0c8e38561077883f0dfb4c3f8e1068c92ef3e9653f0000",
		"202c76ab413de66315922a95c65b0dc77073bf1f9a7e809b0aa51db9f1592e359c2de34ed115c039d356ca573e0d4dc818a258acfc0af48c44c6e4c8d2c9d57508")

	byte_s, err := json.Marshal(tx)
	if err == nil {
		t.Log(string(byte_s))
	}
}
