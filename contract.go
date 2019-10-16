package CocosSDK

import (
	"cocos-go-sdk/rpc"
	. "cocos-go-sdk/type"
	"cocos-go-sdk/wallet"
	"io/ioutil"
)

func ReviseContract(c_name, data string) error {
	if Wallet.Default.Info == nil {
		Wallet.Default.Info = rpc.GetAccountInfoByName(Wallet.Default.Name)
	}
	contract_info := rpc.GetContract(c_name)
	contract := &UpdateContractData{
		ContractID: ObjectId(contract_info.ID),
		Extensions: []interface{}{},
		Data:       String(data),
		Reviser:    ObjectId(Wallet.Default.Info.ID),
	}
	contract.FeeData = Amount{Amount: 0, AssetID: ObjectId("1.3.0")}
	rpc.GetRequireFeeData(59, contract)
	st := wallet.CreateSignTransaction(59, Wallet.Default.GetActiveKey(), contract)
	rpc.BroadcastTransaction(st)
	return nil
}
func ReviseContractByFile(c_name, path string) error {
	byte_s, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}
	data := string(byte_s)
	return ReviseContract(c_name, data)
}

/*创建合约*/
/*c_auth:公钥*/
func CreateContract(c_name, c_auth, data string) error {
	if Wallet.Default.Info == nil {
		Wallet.Default.Info = rpc.GetAccountInfoByName(Wallet.Default.Name)
	}
	contract := &CreateContractData{
		Fee:               EmptyFee(),
		ContractAuthority: c_auth,
		Extensions:        []interface{}{},
		Data:              String(data),
		Name:              String(c_name),
		Owner:             ObjectId(Wallet.Default.Info.ID),
	}
	rpc.GetRequireFeeData(OP_CREATE_CONTRACT, contract)
	st := wallet.CreateSignTransaction(OP_CREATE_CONTRACT, Wallet.Default.GetActiveKey(), contract)
	return rpc.BroadcastTransaction(st)
}
func CreateContractByFile(c_name, c_auth, path string) error {
	byte_s, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	data := string(byte_s)
	return CreateContract(c_name, c_auth, data)
}
func Invoke(contract_name, func_name string, args ...interface{}) {
	contract_info := rpc.GetContract(contract_name)
	value_list := CreateValueList(args)
	if Wallet.Default.Info == nil {
		Wallet.Default.Info = rpc.GetAccountInfoByName(Wallet.Default.Name)
	}
	call_data := &CallData{
		ValueList:    value_list,
		Extensions:   []interface{}{},
		Caller:       ObjectId(Wallet.Default.Info.ID),
		ContractID:   ObjectId(contract_info.ID),
		FunctionName: String(func_name),
	}
	call_data.FeeData = Amount{Amount: 0, AssetID: ObjectId("1.3.0")}
	//fmt.Println(json.Marshal(call_data))
	rpc.GetRequireFeeData(44, call_data)
	st := wallet.CreateSignTransaction(44, Wallet.Default.GetActiveKey(), call_data)
	rpc.BroadcastTransaction(st)
}
