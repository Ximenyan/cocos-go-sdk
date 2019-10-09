package chain

import (
	"cocos-go-sdk/common"
	"cocos-go-sdk/rpc"
	"encoding/hex"
	"log"
	"math/big"
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
		log.Fatalln("Get Chain Properties Error!!!")
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

/*{"id":"1.16.53","creation_date":"2019-07-23T02:11:00","owner":"1.2.35597","name":"contract.taiken","current_version":"f9589cf3f46b505006dd61a3f95c785a338fde8 f679d75cc9b4822afff2a41c7","contract_authority":"COCOS8ST86XoFJw4WxXEsk4psSFdhRFx7anzD6aMmfPYd2nsf24FYur","check_contract_authority":false,"contract_data":[[{"key":[2,{"v":"airDropPool"}]},[0,{"v":281109}]],[{"key":[2,{"v":"drawReward"}]},[4,{"v":[[{"key":[0,{"v":1}]},[4,{"v":[[{"key":[2,{"v":"ist"}]},[0,{"v":1}]],[{"key":[2,{"v":"it"}]},[0,{"v":99}]],[{"key":[2,{"v":"q"}]},[0,{"v":5}]],[{"key":[2,{"v":"w"}]},[0,{"v":355}]]]}]],[{"key":[0,{"v":2}]},[4,{"v":[[{"k ey":[2,{"v":"ist"}]},[0,{"v":16}]],[{"key":[2,{"v":"it"}]},[0,{"v":1}]],[{"key":[2,{"v":"q"}]},[0,{"v":10}]],[{"key":[2,{"v":"w"}]},[0,{"v":50}]]]}]],[{"key":[0,{"v":3}]},[4,{"v":[[{"key":[2,{"v":"ist"}]},[0,{"v":1}]],[{"key":[2,{"v":"it"}]},[0,{"v":99}]],[{"key":[2,{"v":"q"}]},[0,{"v":15}]],[{"key":[2,{"v":"w"}]},[0,{"v":55}]]]}]],[{"key":[0,{"v":4}]},[4,{"v":[[{"key":[2,{"v":"ist"}]},[0,{"v":15}]],[{"key":[2,{"v":"it"}]},[0,{"v":1}]],[{" key":[2,{"v":"q"}]},[0,{"v":2}]],[{"key":[2,{"v":"w"}]},[0,{"v":70}]]]}]],[{"key":[0,{"v":5}]},[4,{"v":[[{"key":[2,{"v":"ist"}]},[0,{"v":1}]],[{"key ":[2,{"v":"it"}]},[0,{"v":99}]],[{"key":[2,{"v":"q"}]},[0,{"v":10}]],[{"key":[2,{"v":"w"}]},[0,{"v":300}]]]}]],[{"key":[0,{"v":6}]},[4,{"v":[[{"key":[2,{"v":"ist"}]},[0,{"v":17}]],[{"key":[2,{"v":"it"}]},[0,{"v":1}]],[{"key":[2,{"v":"q"}]},[0,{"v":1}]],[{"key":[2,{"v":"w"}]},[0,{"v":60}]]]}]],[{"key":[0,{"v":7}]},[4,{"v":[[{"key":[2,{"v":"ist"}]},[0,{"v":1}]],[{"key":[2,{"v":"it"}]},[0,{"v":99}]],[{"key":[2,{"v":"q"}]},[0,{"v":20}]],[{"key ":[2,{"v":"w"}]},[0,{"v":40}]]]}]],[{"key":[0,{"v":8}]},[4,{"v":[[{"key":[2,{"v":"ist"}]},[0,{"v":18}]],[{"key":[2,{"v":"it"}]},[0,{"v":1}]],[{"key":[2,{"v":"q"}]},[0,{"v":2}]],[{"key":[2,{"v":"w"}]},[0,{"v":50}]]]}]],[{"key":[0,{"v":9}]},[4,{"v":[[{"key":[2,{"v":"ist"}]},[0,{"v":1}]],[{"key":[2,{"v":"it"}]},[0,{"v":99}]],[{"key":[2,{"v":"q"}]},[0,{"v":50}]],[{"key":[2,{"v":"w"}]},[0,{"v":15}]]]}]],[{"key":[0,{"v":10}]},[4,{"v":[[{"key":[2,{"v":"ist"}]},[0,{"v":1}]],[{"key":[2,{"v":"it"}]},[0,{"v":99}]],[{"key":[2,{"v":"q"}]},[0,{"v":100}]],[{"key":[2,{"v":"w"}]},[0,{"v":5}]]]}]]]}]],[{"key":[2,{"v":"drawRewardVersion"}]},[0,{"v":12}]],[{"key":[2,{"v":"gifts"}]},[4,{"v":[[{"key":[0,{"v":1}]},[4,{"v":[[{"key":[2,{"v":"cd"}]},[0,{"v":0}]],[{"key":[2,{"v":"q"}]},[0,{"v":15}]],[{"key":[2,{"v":"seed"}]},[0,{"v":2}]]]}]],[{"key":[0,{"v":2}]},[4,{"v":[[{"key":[2,{"v":"cd"}]},[0,{"v":1200}]],[{"key":[2,{"v":"q"}]},[0,{"v":20}]],[{"key":[2,{"v":"seed"}]},[0,{"v":1}]]]}]],[{"key":[0,{"v":3}]},[4,{"v":[[{"key":[2,{"v":"cd"}]},[0,{"v":1800}]],[{"key":[2,{"v":"q"}]},[0,{"v":10}]],[{"key":[2,{"v":"seed"}]},[0,{"v":2}]]]}]],[{"key":[0,{"v":4}]},[4,{"v":[[{"key":[2,{"v":"cd"}]},[0,{" v":2400}]],[{"key":[2,{"v":"q"}]},[0,{"v":5}]],[{"key":[2,{"v":"seed"}]},[0,{"v":2}]]]}]],[{"key":[0,{"v":5}]},[4,{"v":[[{"key":[2,{"v":"cd"}]},[0,{"v":3000}]],[{"key":[2,{"v":"q"}]},[0,{"v":5}]],[{"key":[2,{"v":"seed"}]},[0,{"v":1}]]]}]],[{"key":[0,{"v":6}]},[4,{"v":[[{"key":[2,{"v":"cd"}]},[0,{"v":3300}]],[{"key":[2,{"v":"q"}]},[0,{"v":5}]],[{"key":[2,{"v":"seed"}]},[0,{"v":1}]]]}]]]}]],[{"key":[2,{"v":"goods"}]},[4,{"v":[[{"key":[0,{"v":10}]},[4,{"v":[[{"key":[2,{"v":"ist"}]},[0,{"v":3}]],[{"key":[2,{"v":"it"}]},[0,{"v":1}]],[{"key":[2,{"v":"name"}]},[2,{"v":"861183c7104e2a"}]],[{" key":[2,{"v":"price"}]},[0,{"v":100}]],[{"key":[2,{"v":"q"}]},[0,{"v":10}]]]}]],[{"key":[0,{"v":11}]},[4,{"v":[[{"key":[2,{"v":"ist"}]},[0,{"v":4}]],[{"key":[2,{"v":"it"}]},[0,{"v":1}]],[{"key":[2,{"v":"name"}]},[2,{"v":"644794b16811104e2a"}]],[{"key":[2,{"v":"price"}]},[0,{"v":120}]],[{"key":[2,{"v":"q"}]},[0,{"v":10}]]]}]]]}]],[{"key":[2,{"v":"luckyDrawPool"}]},[0,{"v":93965}]],[{"key":[2,{"v":"luckyDrawTicket"}]},[4,{"v":[[{"key":[2,{"v":" 1.2.100480"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.100481"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.100483"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.11123"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.13932"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.15707"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.32823"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.33145"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.35582"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.35583"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.48160"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.54972"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.58750"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.60346"}]},[0,{"v":1}]],[{" key":[2,{"v":"1.2.60347"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.61290"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.63972"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.6 7316"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.68203"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.71432"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.7305"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.7320"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.7324"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.7325"}]},[0,{"v":1}]],[{"key":[2,{"v":"1. 2.7332"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.7363"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.73886"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.7464"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.84229"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.87262"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.90764"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.90839"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.92646"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.92823"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.92874"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.93121"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.93122"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.93124"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.93126"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.93127"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.93141"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.93155"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.93160"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.93162"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.93166"}]},[0,{"v":1}]],[{"k ey":[2,{"v":"1.2.93213"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.93235"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.93239"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.93 265"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.93272"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.93278"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.93284"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.93288"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.93348"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.93465"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.93496"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.93636"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.93640"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.93643"}]},[0,{" v":1}]],[{"key":[2,{"v":"1.2.93644"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.93652"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.93930"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.93934"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.93939"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.93941"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.93949"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.93952"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.94457"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.94468"}]},[0,{"v":1}]],[{"ke y":[2,{"v":"1.2.94474"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.94480"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.94481"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.944 84"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.94485"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.94486"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.94487"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.94488"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.94489"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.94490"}]},[0,{"v":1}]],[{"key":[2,{"v":" 1.2.94491"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.94493"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.94498"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.94499"}]},[0,{"v ":1}]],[{"key":[2,{"v":"1.2.94500"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.94503"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.94504"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.94506"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.94507"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.94508"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.94513"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.94514"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.94518"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.94527"}]},[0,{"v":1}]],[{"key ":[2,{"v":"1.2.94528"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.94539"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.94542"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.9462 1"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.94878"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.95026"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.95079"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.95230"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.95283"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.95503"}]},[0,{"v":1}]],[{"key":[2,{"v":"1 .2.95504"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.95506"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.95507"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.95509"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.95511"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.96424"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.96429"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.96499"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.96541"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.96601"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.96735"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.97344"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.97345"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.98324"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.98767"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.99384"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.99385"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.99393 "}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.99403"}]},[0,{"v":1}]],[{"key":[2,{"v":"1.2.99409"}]},[0,{"v":1}]]]}]],[{"key":[2,{"v":"seeds"}]},[4,{"v":[[{"ke y":[0,{"v":1}]},[4,{"v":[[{"key":[2,{"v":"name"}]},[2,{"v":"767d83dc"}]],[{"key":[2,{"v":"period"}]},[0,{"v":300}]],[{"key":[2,{"v":"power"}]},[0,{"v":1}]]]}]],[{"key":[0,{"v":2}]},[4,{"v":[[{"key":[2,{"v":"name"}]},[2,{"v":"73897c73"}]],[{"key":[2,{"v":"period"}]},[0,{"v":600}]],[{"key":[2,{"v":"power"}]},[0,{"v":3}]]]}]],[{"key":[0,{"v":3}]},[4,{"v":[[{"key":[2,{"v":"name"}]},[2,{"v":"861183c7"}]],[{"key":[2,{"v":"period"}]},[0,{"v":900}]],[{"key":[2,{"v":"power"}]},[0,{"v":4}]]]}]],[{"key":[0,{"v":4}]},[4,{"v":[[{"key":[2,{"v":"name"}]},[2,{"v":"644794b16811"}]],[{"key":[2,{"v ":"period"}]},[0,{"v":900}]],[{"key":[2,{"v":"power"}]},[0,{"v":5}]]]}]],[{"key":[0,{"v":5}]},[4,{"v":[[{"key":[2,{"v":"name"}]},[2,{"v":"80e1841d535c"}]],[{"key":[2,{"v":"period"}]},[0,{"v":1200}]],[{"key":[2,{"v":"power"}]},[0,{"v":6}]]]}]],[{"key":[0,{"v":6}]},[4,{"v":[[{"key":[2,{"v":"name"}]},[2,{"v":"83045b50"}]],[{"key":[2,{"v":"period"}]},[0,{"v":1200}]],[{"key":[2,{"v":"power"}]},[0,{"v":7}]]]}]],[{"key":[0,{"v":7}]},[4,{"v":[[{"k ey":[2,{"v":"name"}]},[2,{"v":"6708997c"}]],[{"key":[2,{"v":"period"}]},[0,{"v":1200}]],[{"key":[2,{"v":"power"}]},[0,{"v":8}]]]}]],[{"key":[0,{"v":8}]},[4,{"v":[[{"key":[2,{"v":"name"}]},[2,{"v":"95505934"}]],[{"key":[2,{"v":"period"}]},[0,{"v":1500}]],[{"key":[2,{"v":"power"}]},[0,{"v":9}]]]}]],[{"key":[0,{"v":9}]},[4,{"v":[[{"key":[2,{"v":"name"}]},[2,{"v":"535774dc"}]],[{"key":[2,{"v":"period"}]},[0,{"v":1500}]],[{"key":[2,{"v":"power "}]},[0,{"v":10}]]]}]],[{"key":[0,{"v":10}]},[4,{"v":[[{"key":[2,{"v":"name"}]},[2,{"v":"57238bde6811"}]],[{"key":[2,{"v":"period"}]},[0,{"v":1500}]],[{"key":[2,{"v":"power"}]},[0,{"v":11}]]]}]],[{"key":[0,{"v":11}]},[4,{"v":[[{"key":[2,{"v":"name"}]},[2,{"v":"94fe4e4b82b1"}]],[{"key":[2,{"v ":"period"}]},[0,{"v":1800}]],[{"key":[2,{"v":"power"}]},[0,{"v":18}]]]}]],[{"key":[0,{"v":12}]},[4,{"v":[[{"key":[2,{"v":"name"}]},[2,{"v":"8c4682 bd"}]],[{"key":[2,{"v":"period"}]},[0,{"v":1500}]],[{"key":[2,{"v":"power"}]},[0,{"v":12}]]]}]],[{"key":[0,{"v":13}]},[4,{"v":[[{"key":[2,{"v":"name"}]},[2,{"v":"9a6c94c385af"}]],[{"key":[2,{"v":"period"}]},[0,{"v":1800}]],[{"key":[2,{"v":"power"}]},[0,{"v":13}]]]}]],[{"key":[0,{"v":14}]},[4,{"v":[[{"key":[2,{"v":"name"}]},[2,{"v":"7ea25b9d77f3"}]],[{"key":[2,{"v":"period"}]},[0,{"v":1800}]],[{"key":[2,{"v":"power"}]},[0,{"v":14}]]]}]],[{"k ey":[0,{"v":15}]},[4,{"v":[[{"key":[2,{"v":"name"}]},[2,{"v":"6a316843"}]],[{"key":[2,{"v":"period"}]},[0,{"v":1800}]],[{"key":[2,{"v":"power"}]},[0,{"v":15}]]]}]],[{"key":[0,{"v":16}]},[4,{"v":[[{"key":[2,{"v":"name"}]},[2,{"v":"6cf080af"}]],[{"key":[2,{"v":"period"}]},[0,{"v":1800}]],[{"key":[2,{"v":"power"}]},[0,{"v":16}]]]}]],[{"key":[0,{"v":17}]},[4,{"v":[[{"key":[2,{"v":"name"}]},[2,{"v":"661f8fb084dd"}]],[{"key":[2,{"v":"period"}]},[0,{"v":1800}]],[{"key":[2,{"v":"power"}]},[0,{"v":17}]]]}]],[{"key":[0,{"v":18}]},[4,{"v":[[{"key":[2,{"v":"name"}]},[2,{"v":"COCOS"}]],[{"key":[2,{"v":"period"}]},[0,{"v":1800}]],[{"key":[2,{"v":"power"}]},[0,{"v":20}]]]}]]]}]],[{"key":[2,{"v":"shopVersion"}]},[0,{"v":24}]]],"contract_ABI":[[{"key":[2,{"v":"addOrUpdateSeed"}]},[5,{"is_var_arg":false,"arglist":["json"]}]],[{"key":[2,{"v":"beginAirDropDispatch"}]},[5,{"is_var_arg":false,"arglist":[]}]],[{"key":[2,{"v":"buy"}]},[5,{"is_var_arg":false,"arglist":["ver","goodsId"]}]],[{"key":[2,{"v":"closeAirDrop"}]},[5,{"is_var_arg":false,"arglist":[]}]],[{"key ":[2,{"v":"deletePrivate"}]},[5,{"is_var_arg":false,"arglist":["what"]}]],[{"key":[2,{"v":"deletePublic"}]},[5,{"is_var_arg":false,"arglist":["what"]}]],[{"k ey":[2,{"v":"draw"}]},[5,{"is_var_arg":false,"arglist":["ver"]}]],[{"key":[2,{"v":"finishAirDropDispatch"}]},[5,{"is_var_arg":false,"arglist":[]}]],[{"key":[2,{"v":"getGift"}]},[5,{"is_var_arg":false,"arglist":[]}]],[{"key":[2,{"v":"init"}]},[5,{"is_var_arg":false,"arglist":[]}]],[{"key":[2,{"v":"openAirDrop"}]},[5,{"is_var_arg":false,"arglist":["startTs","endTs","symbol","precision","amount","sponsor"]}]],[{"key":[2,{"v":"pick"}]},[5,{"is_var_arg":false,"arglist":["slotIdx"]}]],[{"key":[2,{"v":"plant"}]},[5,{"is_var_arg":false,"arglist":["seed","slotIdx"]}]],[{"key":[2,{"v":"requireAirDrop"}]},[5,{"is_var_arg":false,"arglist":["tks"]}]],[{"key":[2,{"v":"setDrawReward"}]},[5,{"is_var_arg":false,"arglist":["json"]}]],[{"key":[2,{"v":"setGift"}]},[5,{"is_var_arg":false,"ar glist":["json"]}]],[{"key":[2,{"v":"setGoods"}]},[5,{"is_var_arg":false,"arglist":["json"]}]]],"lua_code_b_id":"2.2.53"}*/
type DynamicGlobalProperties struct {
	AccountsRegistered_thisInterval int    `json:"accounts_registered_this_interval"`
	CurrentAslot                    int    `json:"current_aslot"`
	CurrentTransactionCount         int    `json:"current_transaction_count"`
	CurrentWitness                  string `json:"current_witness"`
	DynamicFlags                    int    `json:"dynamic_flags"`
	HeadBlockID                     string `json:"head_block_id"`
	HeadBlockNumber                 int    `json:"head_block_number"`
	ID                              string `json:"id"`
	LastBudgetTime                  string `json:"last_budget_time"`
	LastIrreversib_leBlockNum       int    `json:"last_irreversible_block_num"`
	NextMaintenanceTime             string `json:"next_maintenance_time"`
	RecentSlotsFilled               string `json:"recent_slots_filled"`
	RecentlyMissedCount             int    `json:"recently_missed_count"`
	Time                            string `json:"time"`
	WitnessBudget                   int    `json:"witness_budget"`
}

func GetDynamicGlobalProperties() *DynamicGlobalProperties {
	dgp := &DynamicGlobalProperties{}
	req := rpc.CreateRpcRequest(rpc.CALL,
		[]interface{}{0, `get_dynamic_global_properties`,
			[]interface{}{}})
	if resp, err := rpc.Client.Send(req); err == nil {
		if err = resp.GetInterface(dgp); err == nil {
			return dgp
		}
	}
	return nil
}

func (p *DynamicGlobalProperties) Get_ref_block_num() uint64 {
	return uint64(p.HeadBlockNumber & 0xffff)
}

func (p *DynamicGlobalProperties) Get_ref_block_prefix() uint64 {
	byte_s, _ := hex.DecodeString(p.HeadBlockID)
	//log.Println("Get_ref_block_prefix::", p.HeadBlockNumber)
	//log.Println("Get_ref_block_prefix::", p.HeadBlockID)
	//log.Println("Get_ref_block_prefix:::", common.ReverseBytes(byte_s[4:8]))
	ref_block_prefix := new(big.Int).SetBytes(common.ReverseBytes(byte_s[4:8])).Uint64()
	//log.Println(ref_block_prefix)
	return ref_block_prefix
}
