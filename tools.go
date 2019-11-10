package CocosSDK

import (
	. "cocos-go-sdk/common"
	"cocos-go-sdk/rpc"
	. "cocos-go-sdk/type"
	"cocos-go-sdk/wallet"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/tidwall/gjson"
	//"time"
	//"fmt"
)

type UTXO struct {
	Value   int64  `json:"value"`
	Address string `json:"address"`
	Sn      int64  `json:"sn"`
}

type Tx struct {
	TxHash      string      `json:"tx_hash,omitempty"`
	Inputs      []UTXO      `json:"inputs"`
	Outputs     []UTXO      `json:"outputs"`
	TxAt        string      `json:"tx_at"`
	BlockNumber int64       `json:"block_no,omitempty"`
	ConfirmedAt string      `json:"confirmed_at,omitempty"`
	Extra       interface{} `json:"extra"`
}

func DeserializeTransactions(tx_raw_hex string) (sign_tx *wallet.Signed_Transaction, err error) {

	var byte_s []byte
	//去除chainId
	tx_raw_hex = tx_raw_hex[64:]
	byte_s, err = hex.DecodeString(tx_raw_hex)
	if err != nil {
		return
	}
	sign_tx = new(wallet.Signed_Transaction)
	sign_tx.RefBlockNum = uint64(UintVar(byte_s[:2]))
	sign_tx.RefBlockPrefix = uint64(UintVar(byte_s[2:6]))
	sign_tx.Signatures = []string{}
	sign_tx.Operations = []Operation{}
	sign_tx.ExtensionsData = []interface{}{}
	time_bytes := byte_s[6:10]
	log.Println("time_bytes", time_bytes)
	uinx_time := UintVar(time_bytes)
	log.Println("uinx_time", uinx_time)

	sign_tx.Expiration = time.Unix(int64(uinx_time), 0).In(UTCZone).Format(TIME_FORMAT)
	byte_s = byte_s[10:]
	op_len_bytes := []byte{byte_s[0]}
	for i := 0; byte_s[i] > 0x80; i++ {
		op_len_bytes = append(op_len_bytes, byte_s[i+1])
	}
	op_len := Intvar(op_len_bytes)
	byte_s = byte_s[len(op_len_bytes):]
	for i := 0; i < int(op_len); i++ {
		if byte_s[0] != OP_TRANSFER {
			sign_tx = nil
			err = errors.New("op code id  is not OP_TRANSFER!!!")
			return
		}
		byte_s = byte_s[1:]
		from_bytes := byte_s[0:8]
		byte_s = byte_s[8:]
		to_bytes := byte_s[0:8]
		byte_s = byte_s[8:]
		amount := UintVar(byte_s[0:8])
		amount_asset_id_bytes := byte_s[8:16]
		byte_s = byte_s[16:]
		from_id := UintVar(from_bytes)
		to_id := UintVar(to_bytes)
		amount_asset_id := UintVar(amount_asset_id_bytes)
		tx := Transaction{
			From:           ObjectId(fmt.Sprintf("1.2.%d", from_id)),
			To:             ObjectId(fmt.Sprintf("1.2.%d", to_id)),
			ExtensionsData: []interface{}{},
			AmountData:     Amount{Amount: uint64(amount), AssetID: ObjectId(fmt.Sprintf("1.3.%d", amount_asset_id))},
		}
		if byte_s[0] != 0 {
			log.Println("test_momo", byte_s)
			//移除公钥信息
			from_bytes := make([]byte, 33)
			copy(from_bytes, byte_s[1:34])
			//from_bytes := byte_s[1:34]
			byte_s = byte_s[34:]
			to_bytes := make([]byte, 33)
			copy(to_bytes, byte_s[:33])
			byte_s = byte_s[33:]
			log.Println("to_bytes", to_bytes)
			//移除nonce信息
			nonce_bytes := byte_s[:8]
			byte_s = byte_s[8:]
			m := &Memo{
				From:  wallet.PublicKey(from_bytes).ToBase58String(),
				To:    wallet.PublicKey(to_bytes).ToBase58String(),
				Nonce: uint64(UintVar(nonce_bytes))}
			msg_len_bytes := []byte{byte_s[0]}

			for n := 0; byte_s[n] > 0x80; n++ {
				msg_len_bytes = append(msg_len_bytes, byte_s[n+1])
			}
			byte_s = byte_s[len(msg_len_bytes):]
			//移除msg信息
			msg_len := Intvar(msg_len_bytes)
			m.Message = hex.EncodeToString(byte_s[:msg_len])
			byte_s = byte_s[msg_len:]
			tx.MemoData = m
		} else {
			tx.MemoData = nil
		}
		sign_tx.Operations = append(sign_tx.Operations, Operation{OP_TRANSFER, tx})
	}
	return
}

func Deserialize(tx_raw_hex string) (tx *Tx, err error) {
	var byte_s []byte
	//去除chainId
	tx_raw_hex = tx_raw_hex[64:]
	byte_s, err = hex.DecodeString(tx_raw_hex)
	if err != nil {
		return
	}
	//fmt.Println("byte_s", byte_s)
	time_bytes := byte_s[6:10]
	log.Println("time_bytes", time_bytes)
	uinx_time := UintVar(time_bytes)
	log.Println("uinx_time", uinx_time)
	tx_at := time.Unix(int64(uinx_time), 0).In(UTCZone).Format(TIME_FORMAT)
	byte_s = byte_s[10:]
	//fmt.Println("byte_s", byte_s)
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
			byte_s = byte_s[1:]
			//fee_amount := UintVar(ReverseBytes(byte_s[1:9]))
			//fee_asset_id_bytes := byte_s[9:17]
			//byte_s = byte_s[17:]
			from_bytes := byte_s[0:8]
			byte_s = byte_s[8:]
			to_bytes := byte_s[0:8]
			byte_s = byte_s[8:]
			amount := UintVar(byte_s[0:8])
			amount_asset_id_bytes := byte_s[8:16]
			log.Println("amount_asset_id_bytes", amount_asset_id_bytes)
			byte_s = byte_s[16:]
			log.Println(byte_s)
			if byte_s[0] != 0 {
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
				msg_len := Intvar(msg_len_bytes)
				byte_s = byte_s[msg_len:]
			}

			amount_asset_id := UintVar(amount_asset_id_bytes)
			from_id := UintVar(from_bytes)
			log.Println("from_bytes", from_bytes)
			to_id := UintVar(to_bytes)
			log.Println("to_bytes", to_bytes)
			from_info := rpc.GetAccountInfo(fmt.Sprintf("1.2.%d", from_id))
			to_info := rpc.GetAccountInfo(fmt.Sprintf("1.2.%d", to_id))
			in := UTXO{
				Value:   amount,
				Address: from_info.Name,
				Sn:      amount_asset_id,
			}
			out := UTXO{
				Value:   amount,
				Address: to_info.Name,
				Sn:      amount_asset_id,
			}
			inputs = append(inputs, in)
			outputs = append(outputs, out)

			fmt.Println("op tx_at", tx_at)
			fmt.Println("op code", op_code)
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

func Getrawmempool() (txs []Tx, err error) {
	txs = []Tx{}
	defer func() {
		if recover() != nil {
			txs = nil
			err = errors.New("Getrawmempool Is Error!")
		}
	}()
	dgp := rpc.GetDynamicGlobalProperties()
	for no := dgp.LastIrreversibleBlockNum + 1; no <= dgp.HeadBlockNumber; no++ {
		tx_s, _ := Getblocktxs(int64(no))
		txs = append(txs, tx_s...)
	}
	return
}

func Getblocktxs(count int64) (txs []Tx, err error) {
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

			tx_at := tx.Get("1.expiration").String()
			for index, operation := range tx_operations {
				tx_op_code := operation.Get("0").Int()
				tx_op_data := operation.Get("1")
				if tx_op_code != OP_TRANSFER {
					continue
				}
				fee_amount := tx.Get(fmt.Sprintf("1.operation_results.%d.1.fees.0.amount", index)).Int()
				fee_asset_id_str := tx.Get(fmt.Sprintf("1.operation_results.%d.1.fees.0.asset_id", index)).String()
				fee_asset_id, _ := strconv.ParseInt(
					strings.Split(
						fee_asset_id_str, `.`)[2],
					10, 64)
				out_amount := tx_op_data.Get("amount.amount").Int()
				out_asset_id, _ := strconv.ParseInt(
					strings.Split(
						tx_op_data.Get("amount.asset_id").String(), `.`)[2],
					10, 64)
				from_info := rpc.GetAccountInfo(tx_op_data.Get("from").String())
				to_info := rpc.GetAccountInfo(tx_op_data.Get("to").String())

				in := UTXO{
					Value:   out_amount,
					Address: from_info.Name,
					Sn:      out_asset_id,
				}
				if fee_asset_id != out_asset_id {
					fee_in := UTXO{
						Value:   fee_amount,
						Address: from_info.Name,
						Sn:      fee_asset_id,
					}
					inputs = append(inputs, fee_in)
				} else {
					in.Value += fee_amount
				}
				out := UTXO{
					Value:   out_amount,
					Address: to_info.Name,
					Sn:      out_asset_id,
				}
				inputs = append(inputs, in)
				outputs = append(outputs, out)
			}
			if len(inputs) > 0 && len(outputs) > 0 {
				tx := Tx{
					TxHash:      tx_hash,
					Inputs:      inputs,
					Outputs:     outputs,
					TxAt:        tx_at,
					ConfirmedAt: block.Timestamp,
					Extra:       []interface{}{},
				}
				txs = append(txs, tx)
			}
		}
	}
	return
}

func BalanceForAddress(address string) *[]rpc.Balance {
	return GetAccountBalances(address)
}
func TxsForAddress(address string) (txs []Tx, err error) {
	acct_info := rpc.GetAccountInfoByName(address)
	defer func() {
		if cover := recover(); cover != nil {
			txs = nil
			log.Println(cover)
			err = errors.New("Get Txs For Address Is Error!")
		}
	}()
	txs = []Tx{}
	for _, tx_info := range GetAccountHistorys(acct_info.ID) {
		if byte_s, err := json.Marshal(tx_info); err == nil {
			tx := gjson.ParseBytes(byte_s)
			operation := tx.Get("op")
			tx_op_code := operation.Get("0").Int()
			if tx_op_code != OP_TRANSFER {
				continue
			}
			block_num := tx.Get("block_num").Int()
			trx_in_block := tx.Get("trx_in_block").Int()
			block := GetBlock(block_num)
			tx_info := block.Transactions[trx_in_block]
			if byte_s, err := json.Marshal(tx_info); err == nil {
				tx := gjson.ParseBytes(byte_s)
				tx_hash := tx.Get("0").String()
				tx_at := tx.Get("1.expiration").String()
				tx_operations := tx.Get("1.operations").Array()
				inputs := []UTXO{}
				outputs := []UTXO{}
				for index, operation := range tx_operations {
					tx_op_code := operation.Get("0").Int()
					tx_op_data := operation.Get("1")
					if tx_op_code != OP_TRANSFER {
						continue
					}
					fee_amount := tx.Get(fmt.Sprintf("1.operation_results.%d.1.fees.0.amount", index)).Int()
					fee_asset_id_str := tx.Get(fmt.Sprintf("1.operation_results.%d.1.fees.0.asset_id", index)).String()
					fee_asset_id, _ := strconv.ParseInt(
						strings.Split(
							fee_asset_id_str, `.`)[2],
						10, 64)
					//fee_amount := tx_op_data.Get("fee.amount").Int()
					out_amount := tx_op_data.Get("amount.amount").Int()
					out_asset_id, _ := strconv.ParseInt(
						strings.Split(
							tx_op_data.Get("amount.asset_id").String(), `.`)[2],
						10, 64)
					from_info := rpc.GetAccountInfo(tx_op_data.Get("from").String())
					to_info := rpc.GetAccountInfo(tx_op_data.Get("to").String())
					in := UTXO{
						Value:   out_amount,
						Address: from_info.Name,
						Sn:      out_asset_id,
					}
					if fee_asset_id != out_asset_id {
						fee_in := UTXO{
							Value:   fee_amount,
							Address: from_info.Name,
							Sn:      fee_asset_id,
						}
						inputs = append(inputs, fee_in)
					} else {
						in.Value += fee_amount
					}
					out := UTXO{
						Value:   out_amount,
						Address: to_info.Name,
						Sn:      out_asset_id,
					}
					inputs = append(inputs, in)
					outputs = append(outputs, out)
				}
				if len(inputs) > 0 && len(outputs) > 0 {
					tx := Tx{
						TxHash:      tx_hash,
						Inputs:      inputs,
						Outputs:     outputs,
						TxAt:        block.Timestamp,
						BlockNumber: block_num,
						ConfirmedAt: tx_at,
						Extra:       []interface{}{},
					}
					txs = append(txs, tx)
				}
			}
		}
	}
	return
}

func GetTransaction(tx_hash string) (tx *Tx, err error) {
	tx_info := GetTransactionById(tx_hash)
	block_info := GetTransactionInBlock(tx_hash)
	block := GetBlock(block_info.BlockNum)
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
		for index, operation := range tx_operations {
			tx_op_code := operation.Get("0")
			tx_op_data := operation.Get("1")
			if tx_op_code.Int() != OP_TRANSFER {
				continue
			}
			fee_amount := tx_data.Get(fmt.Sprintf("operation_results.%d.1.fees.0.amount", index)).Int()
			fee_asset_id_str := tx_data.Get(fmt.Sprintf("operation_results.%d.1.fees.0.asset_id", index)).String()
			fee_asset_id, _ := strconv.ParseInt(
				strings.Split(
					fee_asset_id_str, `.`)[2],
				10, 64)
			out_amount := tx_op_data.Get("amount.amount").Int()
			out_asset_id, _ := strconv.ParseInt(
				strings.Split(
					tx_op_data.Get("amount.asset_id").String(), `.`)[2],
				10, 64)
			from_info := rpc.GetAccountInfo(tx_op_data.Get("from").String())
			to_info := rpc.GetAccountInfo(tx_op_data.Get("to").String())
			in := UTXO{
				Value:   out_amount,
				Address: from_info.Name,
				Sn:      out_asset_id,
			}
			if fee_asset_id != out_asset_id {
				fee_in := UTXO{
					Value:   fee_amount,
					Address: from_info.Name,
					Sn:      fee_asset_id,
				}
				inputs = append(inputs, fee_in)
			} else {
				in.Value += fee_amount
			}
			out := UTXO{
				Value:   out_amount,
				Address: to_info.Name,
				Sn:      out_asset_id,
			}
			inputs = append(inputs, in)
			outputs = append(outputs, out)
		}
		if len(inputs) > 0 && len(outputs) > 0 {
			tx = &Tx{
				TxHash:      tx_hash,
				Inputs:      inputs,
				Outputs:     outputs,
				BlockNumber: block_info.BlockNum,
				ConfirmedAt: tx_at,
				TxAt:        block.Timestamp,
				Extra:       []interface{}{},
			}
		}
	}
	return
}

func BuildTransaction(from, to string, amount float64, symbol ...string) (tx_raw_hex string, err error) {
	asset_id := "1.3.0"
	var tk_info *rpc.TokenInfo
	from_info := rpc.GetAccountInfoByName(from)
	to_info := rpc.GetAccountInfoByName(to)
	if len(symbol) > 0 {
		tk_info = rpc.GetTokenInfoBySymbol(symbol[0])
	} else {
		tk_info = rpc.GetTokenInfo(asset_id)
	}
	precision := math.Pow10(tk_info.Precision)
	t := &Transaction{
		AmountData:     Amount{Amount: uint64(float64(amount) * precision), AssetID: ObjectId(tk_info.ID)},
		ExtensionsData: []interface{}{},
		From:           ObjectId(from_info.ID),
		To:             ObjectId(to_info.ID),
		MemoData:       nil,
	}
	op := Operation{OP_TRANSFER, t}
	dgp := rpc.GetDynamicGlobalProperties()
	st := &wallet.Signed_Transaction{
		RefBlockNum:    dgp.Get_ref_block_num(),
		RefBlockPrefix: dgp.Get_ref_block_prefix(),
		Expiration:     time.Unix(time.Now().Unix(), 0).Format(TIME_FORMAT),
		Operations:     []Operation{op},
		ExtensionsData: []interface{}{},
		Signatures:     []string{},
	}
	byte_s := st.GetBytes()
	var cid []byte
	if cid, err = hex.DecodeString(Chain.Properties.ChainID); err != nil {
		return
	}
	byte_s = append(cid, byte_s...)
	log.Println(byte_s)
	tx_raw_hex = hex.EncodeToString(byte_s)
	return
}

func SignTransaction(tx_raw_hex, signature string) (tx *Tx, e error) {
	sign_tx, err := DeserializeTransactions(tx_raw_hex)
	if err != nil {
		e = err
		return
	}

	if byte_s, err := json.Marshal(sign_tx); err == nil {
		tx_json := gjson.ParseBytes(byte_s)
		acct_id := tx_json.Get("operations.0.1.from").String()
		acct_info := rpc.GetAccountInfo(acct_id)
		if wallet.VerifySignature(tx_raw_hex, signature, acct_info.GetActivePuKey()) {
			sign_tx.Signatures = append(sign_tx.Signatures, signature)
			ch := make(chan int)
			rpc.BroadcastTransactionWithHandler(sign_tx, func(r *RpcResp) error {
				hash := r.Result.(string)
				tx, e = GetTransaction(hash)
				ch <- 1
				return nil
			})
			<-ch
		} else {
			e = err
		}
	} else {
		e = err
	}
	return
}
