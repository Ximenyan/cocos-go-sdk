package rpc

type Contract struct {
	ContractData           [][]interface{} `json:"contract_data"`
	CreationDate           string          `json:"creation_date"`
	LuaCodeBID             string          `json:"lua_code_b_id"`
	ID                     string          `json:"id"`
	CheckContractAuthority bool            `json:"check_contract_authority"`
	CurrentVersion         string          `json:"current_version"`
	ContractAuthority      string          `json:"contract_authority"`
	Name                   string          `json:"name"`
	ContractABI            [][]interface{} `json:"contract_ABI"`
	Owner                  string          `json:"owner"`
}

func GetContract(contract_name string) *Contract {
	contract := new(Contract)
	req := CreateRpcRequest(CALL,
		[]interface{}{0, `get_contract`,
			[]interface{}{contract_name}})
	if resp, err := Client.Send(req); err == nil {
		if err = resp.GetInterface(contract); err == nil {
			return contract
		}
		return nil
	}
	return nil
}
