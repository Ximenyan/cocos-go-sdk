package rpc

import (
	"encoding/json"
)

// 返回信息
type RpcResp struct {
	Result  interface{} `json:"result"`
	Id      string      `json:"id"`
	JsonRpc string      `json:"jsonrpc"`
}

func (r *RpcResp) GetInterface(obj interface{}) (err error) {
	var byte_s []byte
	if byte_s, err = json.Marshal(r.Result); err == nil {
		if err = json.Unmarshal(byte_s, obj); err == nil {
			//fmt.Println("GetTransactionById:", obj)
			return nil
		}

	}
	return err
}
