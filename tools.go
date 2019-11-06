package CocosSDK

import (
	. "cocos-go-sdk/type"
	"encoding/hex"
	"fmt"
	//"time"
	//"fmt"
)

type UTXO struct {
	Value   BigInt `json:"value"`
	Address string `json:"address"`
	Sn      int    `json:"sn"`
}

type Tx struct {
	Inputs  []UTXO      `json:"inputs"`
	Outputs []UTXO      `json:"outputs"`
	Extra   interface{} `json:"extra"`
}

func Deserialize(tx_raw_hex string) (tx *Tx, err error) {
	var byte_s []byte
	tx_raw_hex = tx_raw_hex[64:]
	byte_s, err = hex.DecodeString(tx_raw_hex)
	if err != nil {
		return
	}
	time_bytes := byte_s[6:10]
	fmt.Println(byte_s)
	fmt.Println(time_bytes)
	fmt.Println(byte_s[10:])
	return
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
