package types

import (
	"encoding/json"
	"errors"
)

type Notice struct {
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
}

// 返回信息

type RpcResp struct {
	Result interface{} `json:"result"`
	Id     string      `json:"id"`
	Error  *struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
	JsonRpc string `json:"jsonrpc"`
}

func (r *RpcResp) GetInterface(obj interface{}) (err error) {
	var byte_s []byte
	if r.Error != nil {
		return errors.New(r.Error.Message)
	}
	if r.Result == nil {
		err = errors.New("reslut is nil!!")
		return
	}
	if byte_s, err = json.Marshal(r.Result); err == nil {
		if err = json.Unmarshal(byte_s, obj); err == nil {
			return nil
		}
	}
	return err
}
