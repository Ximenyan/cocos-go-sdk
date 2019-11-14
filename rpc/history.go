package rpc

import (
	. "CocosSDK/type"
	"log"
)

type History []struct {
	ID         string        `json:"id"`
	Op         []interface{} `json:"op"`
	Result     []interface{} `json:"result"`
	BlockNum   int           `json:"block_num"`
	TrxInBlock int           `json:"trx_in_block"`
	OpInTrx    int           `json:"op_in_trx"`
	VirtualOp  int           `json:"virtual_op"`
}

func GetAccountHistory(acct_id string) History {
	last := "1.11.0"
	start := "1.11.0"
	get_end := false
	Historys := History{}
	for !get_end {
		historys := History{}
		req := CreateRpcRequest(CALL,
			[]interface{}{HISTORY_API_ID, `get_account_history`,
				[]interface{}{acct_id, last, 100, start}})
		if resp, err := Client.Send(req); err == nil {
			log.Println(resp.Result)
			if err = resp.GetInterface(&historys); err == nil {
				for i := 0; i < len(historys); i++ {
					log.Println(i, historys[i].ID)
				}
				Historys = append(Historys, historys...)
				start = historys[len(historys)-1].ID
				if len(historys) < 50 {
					get_end = true
				}
			}
		}
	}
	return Historys
}

func GetFillOrderHistory(asset_id, _asset_id string, limit uint64) []interface{} {
	req := CreateRpcRequest(CALL,
		[]interface{}{HISTORY_API_ID, `get_fill_order_history`,
			[]interface{}{asset_id, _asset_id, limit}})
	historys := &[]interface{}{}
	if resp, err := Client.Send(req); err == nil {
		log.Println(resp.Result)
		if err = resp.GetInterface(historys); err == nil {
			return *historys
		}
	}
	return nil
}

func GetMarketHistory(asset_id, _asset_id, start, end string, limit uint64) []interface{} {
	req := CreateRpcRequest(CALL,
		[]interface{}{HISTORY_API_ID, `get_market_history`,
			[]interface{}{asset_id, _asset_id, limit, start, end}})
	historys := &[]interface{}{}
	if resp, err := Client.Send(req); err == nil {
		log.Println(resp.Result)
		if err = resp.GetInterface(historys); err == nil {
			return *historys
		}
	}
	return nil
}
