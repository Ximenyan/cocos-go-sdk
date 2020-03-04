package CocosSDK

import (
	"CocosSDK/rpc"
	. "CocosSDK/type"
	"io/ioutil"
)

/*更新合约*/
func ReviseContract(c_name, data string) (string, error) {
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
	return Wallet.SignAndSendTX(OP_REVISE_CONTRACT, contract)
}

/*更新合约*/
func ReviseContractByFile(c_name, path string) (tx_hash string, err error) {
	byte_s, err := ioutil.ReadFile(path)

	if err != nil {
		return tx_hash, err
	}
	data := string(byte_s)
	return ReviseContract(c_name, data)
}

/*创建合约*/
/*c_auth:公钥*/
func CreateContract(c_name, c_auth, data string) (string, error) {
	if Wallet.Default.Info == nil {
		Wallet.Default.Info = rpc.GetAccountInfoByName(Wallet.Default.Name)
	}
	contract := &CreateContractData{
		ContractAuthority: c_auth,
		Extensions:        []interface{}{},
		Data:              String(data),
		Name:              String(c_name),
		Owner:             ObjectId(Wallet.Default.Info.ID),
	}
	return Wallet.SignAndSendTX(OP_CREATE_CONTRACT, contract)
}

/*创建合约*/
func CreateContractByFile(c_name, c_auth, path string) (tx_hash string, err error) {
	byte_s, err := ioutil.ReadFile(path)
	if err != nil {
		return tx_hash, err
	}
	data := string(byte_s)
	return CreateContract(c_name, c_auth, data)
}

/*调用合约*/
func InvokeContract(contract_name, func_name string, args ...interface{}) (string, error) {
	contract_info := GetContract(contract_name)
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
	return Wallet.SignAndSendTX(OP_INVOKE_CONTRACT, call_data)
}

/*查询Contract*/
func GetContract(contract_name string) *rpc.Contract {
	return rpc.GetContract(contract_name)
}

/*查询AccountContract data by name*/
func GetAccountContractData(acct_name,contract_name string) *rpc.Contract {
	acct := rpc.GetAccountInfoByName(acct_name)
	contract := rpc.GetContract(contract_name)
	return rpc.GetAccountContractData(acct.ID,contract.ID)
}


/*查询AccountContract data by id*/
func GetAccountContractDataById(acct_id, contract_id string) *rpc.Contract {
	return rpc.GetAccountContractData(acct_id,contract_id)
}