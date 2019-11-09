package rpc

import (
	. "cocos-go-sdk/type"
	"encoding/json"
	"log"
)

type AccountInfo struct {
	Active struct {
		AccountAuths    []interface{}   `json:"account_auths"`
		AddressAuths    []interface{}   `json:"address_auths"`
		KeyAuths        [][]interface{} `json:"key_auths"`
		WeightThreshold int             `json:"weight_threshold"`
	} `json:"active"`
	ContractAssetLocked struct {
		LockDetails []interface{} `json:"lock_details"`
		LockedTotal []interface{} `json:"locked_total"`
	} `json:"contract_asset_locked"`
	ID                            string `json:"id"`
	LifetimeReferrer              string `json:"lifetime_referrer"`
	LifetimeReferrerFeePercentage int    `json:"lifetime_referrer_fee_percentage"`
	MembershipExpirationDate      string `json:"membership_expiration_date"`
	Name                          string `json:"name"`
	NetworkFeePercentage          int    `json:"network_fee_percentage"`
	Options                       struct {
		Extensions    []interface{} `json:"extensions"`
		MemoKey       string        `json:"memo_key"`
		NumCommittee  int           `json:"num_committee"`
		NumWitness    int           `json:"num_witness"`
		Votes         []interface{} `json:"votes"`
		VotingAccount string        `json:"voting_account"`
	} `json:"options"`
	Owner struct {
		AccountAuths    []interface{}   `json:"account_auths"`
		AddressAuths    []interface{}   `json:"address_auths"`
		KeyAuths        [][]interface{} `json:"key_auths"`
		WeightThreshold int             `json:"weight_threshold"`
	} `json:"owner"`
	Referrer                  string `json:"referrer"`
	ReferrerRewardsPercentage int    `json:"referrer_rewards_percentage"`
	Registrar                 string `json:"registrar"`
	Statistics                string `json:"statistics"`
}
type Balance struct {
	Amount  interface{} `json:"amount"`
	AssetID string      `json:"asset_id"`
}

func (info AccountInfo) GetActivePuKey() string {
	if key, success := info.Active.KeyAuths[0][0].(string); success {
		return key
	}
	return EMPTY
}

func (info AccountInfo) GetMomoPuKey() string {
	if key, success := info.Active.KeyAuths[0][0].(string); success {
		return key
	}
	return EMPTY
}
func (info AccountInfo) GetOwnerPuKey() string {
	if key, success := info.Owner.KeyAuths[0][0].(string); success {
		return key
	}
	return EMPTY
}

type AccountsInfo []*AccountInfo

func GetIdsByPubkeys(pubKeys []string) [][]string {
	ids := &[][]string{}
	req := CreateRpcRequest(CALL,
		[]interface{}{0, `get_key_references`,
			[]interface{}{pubKeys}})
	if resp, err := Client.Send(req); err != nil {
		log.Println(err)
		return *ids
	} else {
		if byte_s, err := json.Marshal(resp.Result); err == nil {
			if err = json.Unmarshal(byte_s, ids); err == nil {
				return *ids
			}
		}
	}
	return *ids
}

func GetAccountsInfo(objIds []string) *AccountsInfo {
	accounts := &AccountsInfo{}
	req := CreateRpcRequest(CALL,
		[]interface{}{0, `get_objects`,
			[]interface{}{objIds}})
	if resp, err := Client.Send(req); err == nil {
		if err = resp.GetInterface(accounts); err == nil {
			return accounts
		}
	}
	return nil
}

func GetAccountInfo(objId string) *AccountInfo {
	accounts := &AccountsInfo{}
	params := []interface{}{objId}
	req := CreateRpcRequest(CALL,
		[]interface{}{0, `get_objects`,
			[]interface{}{params}})
	if resp, err := Client.Send(req); err == nil {
		if err = resp.GetInterface(accounts); err == nil {
			return (*accounts)[0]
		}
	}
	return nil
}
func GetAccountInfoByPublicKey(publicKey string) *AccountInfo {
	ids := GetIdsByPubkeys([]string{publicKey})[0]
	accsInfo := GetAccountsInfo(ids)
	if len(*accsInfo) > 0 {
		return (*accsInfo)[0]
	} else {
		return nil
	}
}

func GetAccountBalances(id string) *[]Balance {
	req := CreateRpcRequest(CALL,
		[]interface{}{0, `get_account_balances`,
			[]interface{}{id, []interface{}{}}})
	if resp, err := Client.Send(req); err == nil {
		balances := &[]Balance{}
		if byte_s, err := json.Marshal(resp.Result); err == nil {
			if err = json.Unmarshal(byte_s, balances); err == nil {
				return balances
			}
		}
	}
	return nil
}

func GetTokenInfoBySymbol(symbol string) *TokenInfo {
	req := CreateRpcRequest(CALL,
		[]interface{}{0, `lookup_asset_symbols`,
			[]interface{}{[]interface{}{symbol}}})
	if resp, err := Client.Send(req); err == nil {
		tokens := &[]*TokenInfo{}
		if byte_s, err := json.Marshal(resp.Result); err == nil {
			if err = json.Unmarshal(byte_s, tokens); err == nil {
				return (*tokens)[0]
			}
			log.Println(err)
		}
	}
	return nil
}
func GetTokenInfosBySymbol(symbols []string) *TokenInfo {
	req := CreateRpcRequest(CALL,
		[]interface{}{0, `lookup_asset_symbols`,
			[]interface{}{symbols}})
	if resp, err := Client.Send(req); err == nil {
		tokens := &[]*TokenInfo{}
		if byte_s, err := json.Marshal(resp.Result); err == nil {
			if err = json.Unmarshal(byte_s, tokens); err == nil {
				return (*tokens)[0]
			}
		}
	}
	return nil
}
func GetTokenInfo(id string) *TokenInfo {
	req := CreateRpcRequest(CALL,
		[]interface{}{0, `get_objects`,
			[]interface{}{[]interface{}{id}}})
	if resp, err := Client.Send(req); err == nil {
		tokens := &[]*TokenInfo{}
		if byte_s, err := json.Marshal(resp.Result); err == nil {
			if err = json.Unmarshal(byte_s, tokens); err == nil {
				return (*tokens)[0]
			}
		}
	}
	return nil
}

func GetAccountInfoByName(name string) *AccountInfo {
	req := CreateRpcRequest(CALL,
		[]interface{}{0, `lookup_account_names`,
			[]interface{}{[]string{name}}})
	if resp, err := Client.Send(req); err == nil {
		accounts := &AccountsInfo{}
		if byte_s, err := json.Marshal(resp.Result); err == nil {
			if err = json.Unmarshal(byte_s, accounts); err == nil {
				return (*accounts)[0]
			}
		}
	}
	return nil
}

func BroadcastTransactionWithHandler(tx interface{}, handler ...func(r *RpcResp) error) error {
	req := CreateRpcRequest(CALL,
		[]interface{}{BROADCAST_API_ID, `broadcast_transaction`,
			[]interface{}{tx}})
	return Client.SendWithHandler(req, handler[0])
}

func BroadcastTransaction(tx interface{}) error {
	req := CreateRpcRequest(CALL,
		[]interface{}{BROADCAST_API_ID, `broadcast_transaction`,
			[]interface{}{tx}})
	if resp, err := Client.Send(req); err == nil {
		txId := ""
		if err = resp.GetInterface(&txId); err == nil {
			log.Println("BroadcastTransaction TXID::", txId)
			return nil
		}
		return err
	} else {
		return err
	}
}
