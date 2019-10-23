package CocosSDK

import (
	"cocos-go-sdk/rpc"
	. "cocos-go-sdk/type"
	"io/ioutil"
)

/*更新合约*/
func ReviseContract(c_name, data string) error {
	if Wallet.Default.Info == nil {
		Wallet.Default.Info = rpc.GetAccountInfoByName(Wallet.Default.Name)
	}
	contract_info := rpc.GetContract(c_name)
	contract := &UpdateContractData{
		Fee:        EmptyFee(),
		ContractID: ObjectId(contract_info.ID),
		Extensions: []interface{}{},
		Data:       String(data),
		Reviser:    ObjectId(Wallet.Default.Info.ID),
	}
	return Wallet.SignAndSendTX(OP_REVISE_CONTRACT, contract)
}

/*更新合约*/
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
	return Wallet.SignAndSendTX(OP_CREATE_CONTRACT, contract)
}

/*创建合约*/
func CreateContractByFile(c_name, c_auth, path string) error {
	byte_s, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	data := string(byte_s)
	return CreateContract(c_name, c_auth, data)
}

/*调用合约*/
func InvokeContract(contract_name, func_name string, args ...interface{}) error {
	contract_info := GetContract(contract_name)
	value_list := CreateValueList(args)
	if Wallet.Default.Info == nil {
		Wallet.Default.Info = rpc.GetAccountInfoByName(Wallet.Default.Name)
	}
	call_data := &CallData{
		Fee:          EmptyFee(),
		ValueList:    value_list,
		Extensions:   []interface{}{},
		Caller:       ObjectId(Wallet.Default.Info.ID),
		ContractID:   ObjectId(contract_info.ID),
		FunctionName: String(func_name),
	}
	return Wallet.SignAndSendTX(OP_INVOKE_CONTRACT, call_data)
}

/*查询Contract*/
func GetContract(contract_name string) *rpc.Contract {
	return rpc.GetContract(contract_name)
}
