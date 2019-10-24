package rpc

import (
	. "cocos-go-sdk/type"
	"encoding/json"
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
		Amount  BigInt `json:"amount"`
		AssetID string `json:"asset_id"`
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

type TokenInfo struct {
	ID        ObjectId `json:"id"`
	Symbol    string   `json:"symbol"`
	Precision int      `json:"precision"`
	Issuer    string   `json:"issuer"`
	Options   struct {
		MaxSupply         interface{} `json:"max_supply"`
		MarketFeePercent  interface{} `json:"market_fee_percent"`
		MaxMarketFee      interface{} `json:"max_market_fee"`
		IssuerPermissions interface{} `json:"issuer_permissions"`
		Flags             int         `json:"flags"`
		CoreExchangeRate  struct {
			Base struct {
				Amount  interface{} `json:"amount"`
				AssetID string      `json:"asset_id"`
			} `json:"base"`
			Quote struct {
				Amount  interface{} `json:"amount"`
				AssetID string      `json:"asset_id"`
			} `json:"quote"`
		} `json:"core_exchange_rate"`
		Description string        `json:"description"`
		Extensions  []interface{} `json:"extensions"`
	} `json:"options"`
	DynamicAssetDataID string `json:"dynamic_asset_data_id"`
}

func GetTokensInfo(ids []string) []*TokenInfo {
	req := CreateRpcRequest(CALL,
		[]interface{}{0, `get_objects`,
			[]interface{}{ids}})
	if resp, err := Client.Send(req); err == nil {
		tokens := &[]*TokenInfo{}
		if byte_s, err := json.Marshal(resp.Result); err == nil {
			if err = json.Unmarshal(byte_s, tokens); err == nil {
				return *tokens
			}
		}
	}
	return nil
}

func QueryTokenList() []*TokenInfo {
	parms :=
		[]interface{}{"A", 100}
	req := CreateRpcRequest(CALL,
		[]interface{}{DATABASE_API_ID, `list_assets`, parms})
	if resp, err := Client.Send(req); err == nil {
		tokens := &[]*TokenInfo{}
		if err = resp.GetInterface(tokens); err == nil {
			return *tokens
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

type WorldViewInfo struct {
	ID                ObjectId `json:"id"`
	WorldView         string   `json:"world_view"`
	WorldViewCreator  string   `json:"world_view_creator"`
	RelatedNhtCreator []string `json:"related_nht_creator"`
}
type WorldViewCreator struct {
	ID        string   `json:"id"`
	WorldView []string `json:"world_view"`
	Creator   string   `json:"creator"`
}
type Proposal struct {
	ID                  string `json:"id"`
	ExpirationTime      string `json:"expiration_time"`
	ProposedTransaction struct {
		RefBlockNum    int             `json:"ref_block_num"`
		RefBlockPrefix int             `json:"ref_block_prefix"`
		Expiration     string          `json:"expiration"`
		Operations     [][]interface{} `json:"operations"`
		Extensions     []string        `json:"extensions"`
	} `json:"proposed_transaction"`
	RequiredActiveApprovals  []string      `json:"required_active_approvals"`
	AvailableActiveApprovals []interface{} `json:"available_active_approvals"`
	RequiredOwnerApprovals   []interface{} `json:"required_owner_approvals"`
	AvailableOwnerApprovals  []interface{} `json:"available_owner_approvals"`
	AvailableKeyApprovals    []interface{} `json:"available_key_approvals"`
	TrxHash                  string        `json:"trx_hash"`
	PermittedClean           bool          `json:"permitted_clean"`
}

func GetWorldViewCreator(creator_id string) *WorldViewCreator {
	parms :=
		[]interface{}{[]string{creator_id}}
	req := CreateRpcRequest(CALL,
		[]interface{}{0, `get_objects`, parms})
	creators := &[]*WorldViewCreator{}
	if resp, err := Client.Send(req); err == nil {
		if err = resp.GetInterface(creators); err == nil {
			return (*creators)[0]
		}
	}
	return nil
}
func GetWorldViewInfo(world_view string) *WorldViewInfo {
	parms :=
		[]interface{}{[]string{world_view}}
	req := CreateRpcRequest(CALL,
		[]interface{}{0, `lookup_world_view`, parms})
	world_views := &[]*WorldViewInfo{}
	if resp, err := Client.Send(req); err == nil {
		if err = resp.GetInterface(world_views); err == nil {
			return (*world_views)[0]
		}
	}
	return nil
}

func GetProposals(acct_id string) *[]Proposal {
	parms := []interface{}{acct_id}
	req := CreateRpcRequest(CALL,
		[]interface{}{0, `get_proposed_transactions`, parms})
	proposals := &[]Proposal{}
	if resp, err := Client.Send(req); err == nil {
		if err = resp.GetInterface(proposals); err == nil {
			return proposals
		}
	}
	return nil
}

func GetProposal(proposal_id string) *Proposal {
	parms :=
		[]interface{}{[]string{proposal_id}}
	req := CreateRpcRequest(CALL,
		[]interface{}{0, `get_objects`, parms})
	proposals := &[]*Proposal{}
	if resp, err := Client.Send(req); err == nil {
		if err = resp.GetInterface(proposals); err == nil {
			return (*proposals)[0]
		}
	}
	return nil
}
