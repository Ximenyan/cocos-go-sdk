package wallet

import (
	"cocos-go-sdk/chain"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	"cocos-go-sdk/common"
	"cocos-go-sdk/crypto/secp256k1"
	"cocos-go-sdk/rpc"
	. "cocos-go-sdk/type"
)

func CreateTransaction(prk *PrivateKey, from_name, to_name, tk_symbol string, value float64, memo string) *Transaction {
	to_info := rpc.GetAccountInfoByName(to_name)
	to_puk := to_info.GetActivePuKey()
	from_info := rpc.GetAccountInfoByName(from_name)
	from_puk := from_info.GetActivePuKey()
	m_data := CreateMemo(prk, from_puk, to_puk, memo)
	tk_info := rpc.GetTokenInfoBySymbol(tk_symbol)

	precision := math.Pow10(tk_info.Precision)
	t := &Transaction{
		Fee:            EmptyFee(),
		AmountData:     Amount{Amount: uint64(float64(value) * precision), AssetID: ObjectId(tk_info.ID)},
		ExtensionsData: []interface{}{},
		From:           ObjectId(from_info.ID),
		To:             ObjectId(to_info.ID),
		MemoData:       m_data,
	}
	return t
}

type Operation []interface{}

func (o Operation) GetBytes() []byte {
	id := int64(o[0].(int))
	id_data := common.VarInt(id, 8)
	opData := o[1].(OpData)
	trans_data := opData.GetBytes()
	byte_s := append(id_data, trans_data...)
	return byte_s
}

func CreateMemo(prk *PrivateKey, from, to, msg string) Memo {
	m := Memo{
		From:    from,
		To:      to,
		Message: msg,
		Nonce:   GetNonce(),
	}
	puk := PukFromBase58String(to)
	x, y := puk.GetPoint()
	cure := secp256k1.S256()
	x, y = cure.ScalarMult(x, y, prk.PrivKey)
	sha := sha512.New()
	byte_s := x.Bytes()
	sha.Write(byte_s)
	resss := sha.Sum(nil)
	noce_s := strconv.FormatUint(m.Nonce, 10)
	seed := noce_s + hex.EncodeToString(resss)
	sha.Reset()
	sha.Write([]byte(seed))
	seed_digest := sha.Sum(nil)
	s256 := sha256.New()
	s256.Write([]byte(msg))
	checksum := s256.Sum(nil)
	byte_s_msg := append(checksum[0:4], []byte(msg)...)
	num := 16 - len(byte_s_msg)%16
	for i := 0; i < num && num != 16; i++ {
		byte_s_msg = append(byte_s_msg, byte(num))
	}
	block, _ := aes.NewCipher(seed_digest[0:32])
	bm := cipher.NewCBCEncrypter(block, seed_digest[32:48])
	bm.CryptBlocks(byte_s_msg, byte_s_msg)
	m.Message = hex.EncodeToString(byte_s_msg)
	return m
}

type Signed_Transaction struct {
	RefBlockNum    uint64      `json:"ref_block_num"`
	RefBlockPrefix uint64      `json:"ref_block_prefix"`
	Expiration     string      `json:"expiration"`
	Operations     []Operation `json:"operations"`
	ExtensionsData Extensions  `json:"extensions"`
	Signatures     []string    `json:"signatures"`
}

func (o Signed_Transaction) GetBytes() []byte {
	block_num_data := common.VarUint(o.RefBlockNum, 16)
	//fmt.Println("block_num_data", block_num_data)
	block_prefix_data := common.VarUint(o.RefBlockPrefix, 32)
	//fmt.Println("block_prefix_data", block_prefix_data)
	t, _ := time.Parse(TIME_FORMAT, o.Expiration)
	expiration_data := common.VarUint(uint64(t.Unix()), 32)
	//fmt.Println("expiration_data", expiration_data)
	operations_data := common.Varint(uint64(len(o.Operations)))
	//fmt.Println("operations_data", operations_data)
	for _, op := range o.Operations {
		operations_data = append(operations_data, op.GetBytes()...)
	}
	extensions_data := o.ExtensionsData.GetBytes()
	byte_s := append(block_num_data,
		append(block_prefix_data,
			append(expiration_data,
				append(operations_data, extensions_data...)...)...)...)
	return byte_s
}

func CreateSignTransaction(opID int, prk *PrivateKey, t OpData) (st *Signed_Transaction, err error) {
	if prk == nil {
		return nil, errors.New("private key is nil!!")
	}
	op := Operation{opID, t}
	dgp := rpc.GetDynamicGlobalProperties()
	st = &Signed_Transaction{
		RefBlockNum:    dgp.Get_ref_block_num(),
		RefBlockPrefix: dgp.Get_ref_block_prefix(),
		Expiration:     time.Unix(time.Now().Unix(), 0).Format(TIME_FORMAT),
		Operations:     []Operation{op},
		ExtensionsData: []interface{}{},
		Signatures:     []string{},
	}
	byte_s := st.GetBytes()
	var cid []byte
	if cid, err = hex.DecodeString(chain.CocosBCXChain.Properties.ChainID); err != nil {
		return nil, err
	}
	byte_s = append(cid, byte_s...)
	fmt.Println(hex.EncodeToString(byte_s))
	msg := sha256digest(byte_s)
	st.Signatures = append(st.Signatures, prk.Sign(msg))
	return st, nil
}

func GetNonce() uint64 {
	rand.Seed(time.Now().Unix())
	return rand.Uint64()
}
