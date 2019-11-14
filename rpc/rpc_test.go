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
}
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
	res := GetNhAssetOrderInfo("4.3.937")
	byte_s, _ := json.Marshal(res)
	t.Log(string(byte_s))
	t.Log(res.Price.Amount)
}
*/
/*
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

func TestQueryToken(t *testing.T) {
	req := CreateRpcRequest(CALL,
		[]interface{}{1, `database`,
			[]interface{}{}})
	if _, err := Client.Send(req); err != nil {
		return
	}
	res := QueryTokenList()
	byte_s, _ := json.Marshal(res)
	t.Log(string(byte_s))
}

func TestLookWorldView(t *testing.T) {
	res := GetWorldViewInfo("block_chain")
	byte_s, _ := json.Marshal(res)
	t.Log(string(byte_s))
}

func TestQueryBlockHeader(t *testing.T) {
	req := CreateRpcRequest(CALL,
		[]interface{}{1, `database`,
			[]interface{}{}})
	if _, err := Client.Send(req); err != nil {
		return
	}
	res := GetBlockHeader(5510688)
	byte_s, _ := json.Marshal(res)
	t.Log(string(byte_s))
}

func TestQueryVstBalance(t *testing.T) {
	req := CreateRpcRequest(CALL,
		[]interface{}{1, `database`,
			[]interface{}{}})
	if _, err := Client.Send(req); err != nil {
		return
	}
	res := GetVestingBalancesByName("ximenyan1111")
	byte_s, _ := json.Marshal(res)
	t.Log(string(byte_s))
}

func TestFillOrderHistory(t *testing.T) {
	req := CreateRpcRequest(CALL,
		[]interface{}{1, `history`,
			[]interface{}{}})
	if _, err := Client.Send(req); err != nil {
		return
	}
	res := GetFillOrderHistory("1.3.0", "1.3.1", 10)
	byte_s, _ := json.Marshal(res)
	t.Log(string(byte_s))
}

func TestMarketHistory(t *testing.T) {
	req := CreateRpcRequest(CALL,
		[]interface{}{1, `history`,
			[]interface{}{}})
	if _, err := Client.Send(req); err != nil {
		return
	}
	res := GetMarketHistory("1.3.0", "1.3.1", "2019-10-10T06:56:11", "2019-10-12T06:56:11", 86400)
	byte_s, _ := json.Marshal(res)
	t.Log(string(byte_s))
}
*/
