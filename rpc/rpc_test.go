package rpc

import (
	"encoding/json"
	"testing"
)

const TEST_NET = "47.93.62.96"
const LOCAL = "192.168.0.166"

var _ = InitClient(TEST_NET, 8049, false)

/*
func TestGetAccountsInfo(t *testing.T) {
	res := GetIdsByPubkeys([]string{"COCOS6zfzShioGyBcvcyFB4Xfzcdo8T7vbXKzRqRG4mJH7aws9BJ88e"})
	res2 := GetAccountsInfo(res[0])
	t.Log((*res2)[0].Name)
}

func TestLookByName(t *testing.T) {
	res := GetAccountInfoByName("sept925")
	byte_s, _ := json.Marshal(res)
	t.Log(string(byte_s))
}*/

func TestGetBalances(t *testing.T) {
	res := GetAccountBalances("1.2.94622")
	byte_s, _ := json.Marshal(res)
	t.Log(string(byte_s))
}

func TestGetTokenInfo(t *testing.T) {
	res := GetTokenInfoBySymbol("WTH")
	byte_s, _ := json.Marshal(res)
	t.Log(string(byte_s))
}

func TestGetOrderInfo(t *testing.T) {
	res := GetNhAssetOrderInfo("4.3.807")
	byte_s, _ := json.Marshal(res)
	t.Log(string(byte_s))
}

func TestListNhAsset(t *testing.T) {
	res := GetNhAssetList("ximenyan1111", 1, 10, 3, []string{"block_chain"})
	byte_s, _ := json.Marshal(res)
	t.Log(string(byte_s))
}
func TestListNhAssetOrder(t *testing.T) {
	res := GetAccountNhAssetOrderList("ximenyan1111", 1, 10)
	byte_s, _ := json.Marshal(res)
	t.Log(string(byte_s))
}
func TestTXInfo(t *testing.T) {
	res := GetTransactionById("49a78da275347277ee86cbbb08020a3ce12c82a1f7b2640b14c4f0d27ac64300")
	byte_s, _ := json.Marshal(res)
	t.Log(string(byte_s))
}

func TestGetContract(t *testing.T) {
	c := GetContract("contract.taiken")
	t.Log(c)
}
