package rpc

import (
	"encoding/json"
	"time"
)

const jsonRpcVersion = "2.0"

// 请求信息
type RpcRequest struct {
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Id      int64         `json:"id"`
	JsonRpc string        `json:"jsonrpc"`
}

func CreateRpcRequest(Method string, Params []interface{}) *RpcRequest {
	r := &RpcRequest{Method, Params, time.Now().UnixNano(), jsonRpcVersion}
	return r
}

func (r *RpcRequest) ToString() string {
	byte_s, _ := json.Marshal(r)
	return string(byte_s)
}
