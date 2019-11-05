package CocosSDK

import (
	. "cocos-go-sdk/type"
)

type UTXO struct {
	Value   BigInt `json:"value"`
	Address string `json:"address"`
	Sn      int    `json:"sn"`
}

type Transaction struct {
	Inputs  []UTXO      `json:"inputs"`
	Outputs []UTXO      `json:"outputs"`
	Extra   interface{} `json:"extra"`
}

func Deserialize() {
	//GetTransactionById()
}

func Unsigned_tx_hash() {

}

func Getblockcount() {

}

func Getrawmempool() {

}

func Getblocktxs() {

}

func BalanceForAddress() {

}
func TxsForAddress() {

}

func GetTransaction() {

}

func BuildTransaction() {

}

func SignTransaction() {

}
