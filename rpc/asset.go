package rpc

import (
	"encoding/json"
	"math/big"
)

type NhAssetOrderInfo struct {
	ID             string `json:"id"`
	Seller         string `json:"seller"`
	Otcaccount     string `json:"otcaccount"`
	NhAssetID      string `json:"nh_asset_id"`
	AssetQualifier string `json:"asset_qualifier"`
	WorldView      string `json:"world_view"`
	BaseDescribe   string `json:"base_describe"`
	NhHash         string `json:"nh_hash"`
	Price          struct {
		Amount  *big.Int `json:"amount"`
		AssetID string   `json:"asset_id"`
	} `json:"price"`
	Memo       string `json:"memo"`
	Expiration string `json:"expiration"`
}
type NHAssetInfo struct {
	ID                   string        `json:"id"`
	NhHash               string        `json:"nh_hash"`
	NhAssetCreator       string        `json:"nh_asset_creator"`
	NhAssetOwner         string        `json:"nh_asset_owner"`
	NhAssetActive        string        `json:"nh_asset_active"`
	Dealership           string        `json:"dealership"`
	AssetQualifier       string        `json:"asset_qualifier"`
	WorldView            string        `json:"world_view"`
	BaseDescribe         string        `json:"base_describe"`
	Parent               []interface{} `json:"parent"`
	Child                []interface{} `json:"child"`
	DescribeWithContract []interface{} `json:"describe_with_contract"`
	CreateTime           string        `json:"create_time"`
	LimitList            []interface{} `json:"limit_list"`
	LimitType            string        `json:"limit_type"`
}
type AssetsList struct {
	Assets    []NHAssetInfo
	WorldView []string
	Page      int
	Limit     int
}

type OrdersList struct {
	Orders []NhAssetOrderInfo
	Page   int
	Limit  int
}

func GetNhAssetOrderInfo(id string) *NhAssetOrderInfo {
	req := CreateRpcRequest(CALL,
		[]interface{}{0, `get_objects`,
			[]interface{}{[]string{id}}})
	if resp, err := Client.Send(req); err == nil {
		orders := &[]*NhAssetOrderInfo{}
		if byte_s, err := json.Marshal(resp.Result); err == nil {
			if err = json.Unmarshal(byte_s, orders); err == nil {
				return (*orders)[0]
			}
		}
	}
	return nil
}

/*根据标准资产assst_name 查看NH市场订单*/
func GetNhAssetOrderList(assst_name, world_view string, page, page_size int) *OrdersList {
	parms :=
		[]interface{}{assst_name, world_view, "", page_size, page, true}
	req := CreateRpcRequest(CALL,
		[]interface{}{0, `list_nh_asset_order`, parms})
	if resp, err := Client.Send(req); err == nil {
		orders := &[]NhAssetOrderInfo{}
		list := &OrdersList{}
		if res_arr, s := resp.Result.([]interface{}); s {
			if byte_s, err := json.Marshal(res_arr[0]); err == nil {
				if err = json.Unmarshal(byte_s, orders); err == nil {
					list.Orders = *orders
					list.Page = page
					list.Limit = int(res_arr[1].(float64))
					return list
				}
			}
		}
	}
	return nil
}
func GetAccountNhAssetOrderList(owner_name string, page, page_size int) *OrdersList {
	acc_info := GetAccountInfoByName(owner_name)
	parms :=
		[]interface{}{acc_info.ID, page_size, page}
	req := CreateRpcRequest(CALL,
		[]interface{}{0, `list_account_nh_asset_order`, parms})
	if resp, err := Client.Send(req); err == nil {
		orders := &[]NhAssetOrderInfo{}
		list := &OrdersList{}
		if res_arr, s := resp.Result.([]interface{}); s {
			if byte_s, err := json.Marshal(res_arr[0]); err == nil {
				if err = json.Unmarshal(byte_s, orders); err == nil {
					list.Orders = *orders
					list.Page = page
					list.Limit = int(res_arr[1].(float64))
					return list
				}
			}
		}
	}
	return nil
}

func GetNhAssetList(acc_name string, page, page_size, _type int, world_view []string) *AssetsList {
	acc_info := GetAccountInfoByName(acc_name)
	parms :=
		[]interface{}{acc_info.ID, world_view, page_size, page, _type}
	req := CreateRpcRequest(CALL,
		[]interface{}{0, `list_account_nh_asset`, parms})
	if resp, err := Client.Send(req); err == nil {
		assets := &[]NHAssetInfo{}
		list := &AssetsList{}
		if res_arr, s := resp.Result.([]interface{}); s {
			if byte_s, err := json.Marshal(res_arr[0]); err == nil {
				if err = json.Unmarshal(byte_s, assets); err == nil {
					list.Assets = *assets
					list.WorldView = world_view
					list.Page = page
					list.Limit = int(res_arr[1].(float64))
					return list
				}
			}
		}
	}
	return nil
}
