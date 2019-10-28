package rpc

import (
	. "Go-SDK/type"
	"log"
)

func GetAccountHistory(acct_id string) []interface{} {
	req := CreateRpcRequest(CALL,
		[]interface{}{HISTORY_API_ID, `get_account_history`,
			[]interface{}{acct_id}})
	historys := &[]interface{}{}
	if resp, err := Client.Send(req); err == nil {
		log.Println(resp.Result)
		if err = resp.GetInterface(historys); err == nil {
			return *historys
		}
	}
	return nil
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
