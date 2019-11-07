package CocosSDK

import (
	. "cocos-go-sdk/common"
	"cocos-go-sdk/rpc"
	. "cocos-go-sdk/type"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/tidwall/gjson"
	//"time"
	//"fmt"
)

type UTXO struct {
	Value   int64  `json:"value"`
	Address string `json:"address"`
	Sn      int    `json:"sn"`
}

type Tx struct {
	TxHash  string      `json:"tx_hash"`
	Inputs  []UTXO      `json:"inputs"`
	Outputs []UTXO      `json:"outputs"`
	TxAt    string      `json:"tx_at"`
	Extra   interface{} `json:"extra"`
}

func Deserialize(tx_raw_hex string) (tx *Tx, err error) {
	var byte_s []byte
	//去除chainId
	tx_raw_hex = tx_raw_hex[64:]
	byte_s, err = hex.DecodeString(tx_raw_hex)
	if err != nil {
		return
	}
	time_bytes := byte_s[6:10]
	uinx_time := UintVar(ReverseBytes(time_bytes))
	tx_at := time.Unix(int64(uinx_time), 0).Format(TIME_FORMAT)
	byte_s = byte_s[10:]
	op_len_bytes := []byte{byte_s[0]}
	for i := 0; byte_s[i] > 0x80; i++ {
		op_len_bytes = append(op_len_bytes, byte_s[i+1])
	}
	op_len := Intvar(op_len_bytes)
	op_code := 0
	byte_s = byte_s[len(op_len_bytes):]
	inputs := []UTXO{}
	outputs := []UTXO{}
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
			byte_s = byte_s[67:]
			//移除nonce信息
			byte_s = byte_s[8:]
			msg_len_bytes := []byte{byte_s[0]}
			for n := 0; byte_s[n] > 0x80; n++ {
				msg_len_bytes = append(msg_len_bytes, byte_s[n+1])
			}
			byte_s = byte_s[len(msg_len_bytes):]
			//移除msg信息
			//fmt.Println(msg_len_bytes)
			msg_len := Intvar(msg_len_bytes)
			byte_s = byte_s[msg_len:]
			fee_asset_id := Intvar(fee_asset_id_bytes)
			amount_asset_id := Intvar(amount_asset_id_bytes)
			from_id := Intvar(from_bytes)
			to_id := Intvar(to_bytes)
			from_info := rpc.GetAccountInfo(fmt.Sprintf("1.2.%d", from_id))
			to_info := rpc.GetAccountInfo(fmt.Sprintf("1.2.%d", to_id))
			in := UTXO{
				Value:   fee_amount + amount,
				Address: from_info.Name,
				Sn:      0,
			}
			out := UTXO{
				Value:   amount,
				Address: to_info.Name,
				Sn:      0,
			}
			inputs = append(inputs, in)
			outputs = append(outputs, out)

			fmt.Println("op tx_at", tx_at)
			fmt.Println("op code", op_code)
			fmt.Println("Fee amount", fee_amount)
			fmt.Println("fee_asset_id", "1.3.", fee_asset_id)
			fmt.Println("from id", from_info.Name)
			fmt.Println("to id", to_info.Name)
			fmt.Println(" amount", amount)
			fmt.Println("amount_asset_id", "1.3.", amount_asset_id)
		}
	}
	tx = &Tx{
		Inputs:  inputs,
		Outputs: outputs,
		TxAt:    tx_at,
		Extra:   []interface{}{},
	}
	return
}

func Unsigned_tx_hash() {

}

func PublicToAddress(puk string) (address string, err error) {
	acct := rpc.GetAccountInfoByPublicKey(puk)
	if acct != nil {
		address = acct.Name
	} else {
		err = errors.New("not found the public key in database.")
	}
	return
}

func AddressToPublic(address string) (puk string, err error) {
	acct := rpc.GetAccountInfoByName(address)
	if acct != nil {
		puk = acct.GetActivePuKey()
	} else {
		err = errors.New("not found the name in database.")
	}
	return
}

func Getblockcount() int {
	return rpc.GetDynamicGlobalProperties().HeadBlockNumber
}

func Getrawmempool() {

}

func Getblocktxs(count int) (txs []Tx, err error) {
	block := GetBlock(count)
	defer func() {
		if recover() != nil {
			txs = nil
			err = errors.New("Getblocktxs Is Error!")
		}
	}()
	txs = []Tx{}
	for _, tx_info := range block.Transactions {
		if byte_s, err := json.Marshal(tx_info); err == nil {
			tx := gjson.ParseBytes(byte_s)
			tx_hash := tx.Get("0").String()
			tx_operations := tx.Get("1.operations").Array()
			inputs := []UTXO{}
			outputs := []UTXO{}
			for _, operation := range tx_operations {
				tx_op_code := operation.Get("0").Int()
				tx_op_data := operation.Get("1")
				if tx_op_code != OP_TRANSFER {
					continue
				}
				fee_amount := tx_op_data.Get("fee.amount").Int()
				out_amount := tx_op_data.Get("amount.amount").Int()
				from_info := rpc.GetAccountInfo(tx_op_data.Get("from").String())
				to_info := rpc.GetAccountInfo(tx_op_data.Get("to").String())
				in := UTXO{
					Value:   fee_amount + out_amount,
					Address: from_info.Name,
					Sn:      0,
				}
				out := UTXO{
					Value:   out_amount,
					Address: to_info.Name,
					Sn:      0,
				}
				inputs = append(inputs, in)
				outputs = append(outputs, out)
			}
			if len(inputs) > 0 && len(outputs) > 0 {
				tx := Tx{
					TxHash:  tx_hash,
					Inputs:  inputs,
					Outputs: outputs,
					TxAt:    block.Timestamp,
					Extra:   []interface{}{},
				}
				txs = append(txs, tx)
			}
		}
	}
	return
}

func BalanceForAddress() {

}
func TxsForAddress(address string) {
	GetAccountHistorys(address)
}

func GetTransaction(tx_hash string) (tx *Tx, err error) {
	tx_info := GetTransactionById(tx_hash)
	if tx_info == nil {
		err = errors.New("transaction not found!!!!")
		return
	}
	defer func() {
		if recover() != nil {
			tx = nil
			err = errors.New("Getblocktxs Is Error!")
		}
	}()
	if byte_s, err := json.Marshal(tx_info); err == nil {
		tx_data := gjson.ParseBytes(byte_s)
		tx_at := tx_data.Get("expiration").String()
		tx_operations := tx_data.Get("operations").Array()
		inputs := []UTXO{}
		outputs := []UTXO{}
		for _, operation := range tx_operations {
			tx_op_code := operation.Get("0")
			tx_op_data := operation.Get("1")
			if tx_op_code.Int() != OP_TRANSFER {
				continue
			}
			fee_amount := tx_op_data.Get("fee.amount").Int()
			out_amount := tx_op_data.Get("amount.amount").Int()
			from_info := rpc.GetAccountInfo(tx_op_data.Get("from").String())
			to_info := rpc.GetAccountInfo(tx_op_data.Get("to").String())
			in := UTXO{
				Value:   fee_amount + out_amount,
				Address: from_info.Name,
				Sn:      0,
			}
			out := UTXO{
				Value:   out_amount,
				Address: to_info.Name,
				Sn:      0,
			}
			inputs = append(inputs, in)
			outputs = append(outputs, out)
		}
		if len(inputs) > 0 && len(outputs) > 0 {
			tx = &Tx{
				Inputs:  inputs,
				Outputs: outputs,
				TxAt:    tx_at,
				Extra:   []interface{}{},
			}
		}
	}
	return
}

func BuildTransaction() {

}

func SignTransaction() {

}
