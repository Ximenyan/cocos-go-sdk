package chain // import "CocosSDK/chain"

import (
	"CocosSDK/rpc"
	"log"
)

type ChainProperties struct {
	BaseContract        string `json:"base_contract"`
	ChainID             string `json:"chain_id"`
	ID                  string `json:"id"`
	ImmutableParameters struct {
		MinCommitteeMemberCount int `json:"min_committee_member_count"`
		MinWitnessCount         int `json:"min_witness_count"`
		NumSpecialAccounts      int `json:"num_special_accounts"`
		NumSpecialAssets        int `json:"num_special_assets"`
	} `json:"immutable_parameters"`
}

type Chain struct {
	Properties *ChainProperties
}

var CocosBCXChain *Chain

func GetChainProperties() *ChainProperties {
	req := rpc.CreateRpcRequest(rpc.CALL,
		[]interface{}{0, `get_chain_properties`,
			[]interface{}{}})
	if resp, err := rpc.Client.Send(req); err == nil {
		var c ChainProperties
		if err = resp.GetInterface(&c); err == nil {
			return &c
		}
	}
	return nil
}
func InitChain() {
	Login("", "")
	Database()
	History()
	Network_broadcast()
	Propertie := GetChainProperties()
	if Propertie == nil {
		log.Panic("Get Chain Properties Error!!!")
	}
	CocosBCXChain = &Chain{
		Properties: Propertie,
	}
}
func Login(user, pwd string) bool {
	req := rpc.CreateRpcRequest(rpc.CALL,
		[]interface{}{1, `login`,
			[]interface{}{user, pwd}})
	if resp, err := rpc.Client.Send(req); err == nil {
		return resp.Result.(bool)
	}
	return false
}

func Database() error {
	req := rpc.CreateRpcRequest(rpc.CALL,
		[]interface{}{1, `database`,
			[]interface{}{}})
	if _, err := rpc.Client.Send(req); err != nil {
		return err
	}
	return nil
}

func History() error {
	req := rpc.CreateRpcRequest(rpc.CALL,
		[]interface{}{1, `history`,
			[]interface{}{}})
	if _, err := rpc.Client.Send(req); err != nil {
		return err
	}
	return nil
}

func Network_broadcast() error {
	req := rpc.CreateRpcRequest(rpc.CALL,
		[]interface{}{1, `network_broadcast`,
			[]interface{}{}})
	if _, err := rpc.Client.Send(req); err != nil {
		return err
	}
	return nil
}

/*获取节点ID*/
func GetChainID() string {
	req := rpc.CreateRpcRequest(rpc.CALL,
		[]interface{}{0, `get_chain_id`,
			[]interface{}{}})
	if resp, err := rpc.Client.Send(req); err == nil {
		return resp.Result.(string)
	}
	return rpc.EMPTY
}
