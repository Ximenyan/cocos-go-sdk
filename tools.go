package CocosSDK

import (
	. "cocos-go-sdk/common"
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
	//time_bytes := byte_s[6:10]
	byte_s = byte_s[10:]
	op_len_bytes := []byte{byte_s[0]}
	for i := 0; byte_s[i] > 0x80; i++ {
		op_len_bytes = append(op_len_bytes, byte_s[i+1])
	}
	op_len := Intvar(op_len_bytes)
	op_code := 0
	byte_s = byte_s[len(op_len_bytes):]
	for i := 0; i < int(op_len); i++ {
		if byte_s[0] == byte(op_code) {
			fee_amount := UintVar(ReverseBytes(byte_s[1:9]))
			fee_asset_id_bytes := []byte{byte_s[9]}
			for n := 9; byte_s[n] > 0x80; n++ {
				fee_asset_id_bytes = append(fee_asset_id_bytes, byte_s[n+1])
			}
			byte_s = byte_s[9+len(fee_asset_id_bytes):]

			from_bytes := []byte{byte_s[0]}
			for n := 0; byte_s[n] > 0x80; n++ {
				from_bytes = append(from_bytes, byte_s[n+1])
			}
			byte_s = byte_s[len(from_bytes):]

			to_bytes := []byte{byte_s[0]}
			for n := 0; byte_s[n] > 0x80; n++ {
				to_bytes = append(to_bytes, byte_s[n+1])
			}
			byte_s = byte_s[len(to_bytes):]
			amount := UintVar(ReverseBytes(byte_s[0:8]))
			amount_asset_id_bytes := []byte{byte_s[8]}
			for n := 8; byte_s[n] > 0x80; n++ {
				amount_asset_id_bytes = append(amount_asset_id_bytes, byte_s[n+1])
			}
			byte_s = byte_s[8+len(amount_asset_id_bytes):]
			//移除公钥信息
			byte_s = byte_s[66:]
			//移除nonce信息
			byte_s = byte_s[8:]
			msg_len_bytes := []byte{byte_s[0]}
			for n := 0; byte_s[n] > 0x80; n++ {
				msg_len_bytes = append(msg_len_bytes, msg_len_bytes[n+1])
			}
			byte_s = byte_s[len(to_bytes):]
			//移除msg信息

			fee_asset_id := Intvar(fee_asset_id_bytes)
			amount_asset_id := Intvar(amount_asset_id_bytes)
			from_id := Intvar(from_bytes)
			to_id := Intvar(to_bytes)

			fmt.Println("op code", op_code)
			fmt.Println("Fee amount", fee_amount)
			fmt.Println("fee_asset_id", "1.3.", fee_asset_id)
			fmt.Println("from id", "1.2.", from_id)
			fmt.Println("to id", "1.2.", to_id)
			fmt.Println(" amount", amount)
			fmt.Println("amount_asset_id", "1.3.", amount_asset_id)
		}
	}
	fmt.Println(byte_s)
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
