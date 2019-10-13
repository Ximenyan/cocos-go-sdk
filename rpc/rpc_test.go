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

func TestTXInfo(t *testing.T) {
	GetTransactionById("34d90151d88b1dd1b8edb1652f383c28dc6fc5cbb834f813ae51fdac0d543214")
}

func TestGetContract(t *testing.T) {
	c := GetContract("contract.taiken")
	t.Log(c)
}
