package CocosSDK

import (
	//"cocos-go-sdk/type"
	//	"encoding/json"
	"testing"
)

//000000000001030ed1f4745aeb7194e1eea53bf6c4a217ba3b8f7d63ebad2e22543b99469bb03202b8764675d8a40f416ba1571bd67e2b6b61c0adf68446af8b3b878efeeccb6e0ee8cda626af16a6dc106094ec3577a9e705cdaf5ed006f9397b0000
//7d89b84f22af0b150780a2b121aa6c715b19261c8b7fe0fda3a564574ed7d3e9ed3b9ae2f85f1a1ac35d0100a2510000000000000097d505c7e205b0ad0100000000000001030ed1f4745aeb7194e1eea53bf6c4a217ba3b8f7d63ebad2e22543b99469bb03202b8764675d8a40f416ba1571bd67e2b6b61c0adf68446af8b3b878efeeccb6e0e9ce14683d52d68191031f641c2a14dfb96ed226bc5ba2783000000

func TestInitSdk(t *testing.T) {
	InitSDK("47.93.62.96", 8049, false)
	t.Log(Chain.Properties.ImmutableParameters)
}

func TestTxsForAddress(t *testing.T) {
	TxsForAddress("1.2.92823")
}

/*
func TestGetTransaction(t *testing.T) {
	tx, err := GetTransaction("c66b0e334409a4a1a8e2d45fa6864ddccc5ab5c2e20a72daeb5f633432fa0844")
	t.Log(err)
	byte_s, err := json.Marshal(tx)
	if err == nil {
		t.Log(string(byte_s))
	}
}

func TestGetBlockCount(t *testing.T) {
	t.Log(Getblockcount())
}
func TestPuk2Addr(t *testing.T) {
	t.Log(PublicToAddress("COCOS6wm6Cqmz82xdxsaXMAiffTRaLDNAS4UAEmyGfTxWq5PSCT2ekw"))
}
func TestAddr2Puk(t *testing.T) {
	t.Log(AddressToPublic("ximenyan1111"))
}
func TestGetblocktxs(t *testing.T) {
	txs, err := Getblocktxs(6254215)
	byte_s, err := json.Marshal(txs)
	if err == nil {
		t.Log(string(byte_s))
	}
}

func TestDeserialize(t *testing.T) {
	var hex_str string = "7d89b84f22af0b150780a2b121aa6c715b19261c8b7fe0fda3a564574ed7d3e9293a25b55f5b8c16c35d0100a2510000000000000097d505c7e205b0ad0100000000000001030ed1f4745aeb7194e1eea53bf6c4a217ba3b8f7d63ebad2e22543b99469bb03202b8764675d8a40f416ba1571bd67e2b6b61c0adf68446af8b3b878efeeccb6e0ee8cda626af16a6dc106094ec3577a9e705cdaf5ed006f9397b0000"
	t.Log(len(types.PukBytesFromBase58String("COCOS6wm6Cqmz82xdxsaXMAiffTRaLDNAS4UAEmyGfTxWq5PSCT2ekw")))
	tx, err := Deserialize(hex_str)
	byte_s, err := json.Marshal(tx)
	if err == nil {
		t.Log(string(byte_s))
	}
}
*/
