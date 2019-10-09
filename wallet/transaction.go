package wallet

import (
	"cocos-go-sdk/chain"
	"cocos-go-sdk/common"
	"cocos-go-sdk/crypto/secp256k1"
	"cocos-go-sdk/rpc"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type ObjectId string
type Object interface {
	GetBytes() []byte
}

func (o ObjectId) GetBytes() []byte {
	num := strings.Split(string(o), `.`)[2]
	i, _ := strconv.ParseUint(num, 10, 64)
	return common.Varint(i)
}

type Amount struct {
	Amount  uint64   `json:"amount"`
	AssetID ObjectId `json:"asset_id"`
}

func (a Amount) GetBytes() []byte {
	byte_s := append(common.VarUint(a.Amount, 64), a.AssetID.GetBytes()...)
	return byte_s
}

type Memo struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Nonce   uint64 `json:"nonce"`
	Message string `json:"message"`
}

func (o Memo) GetBytes() []byte {
	from := PukFromBase58String(o.From)
	to := PukFromBase58String(o.To)
	nonce := common.VarUint(o.Nonce, 64)
	msg, _ := hex.DecodeString(o.Message)
	msg = append([]byte{byte(len(msg))}, msg...)
	byte_s := append([]byte{0x01},
		append(from,
			append(to,
				append(nonce, msg...)...)...)...,
	)
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

type Extensions []interface{}

func (o Extensions) GetBytes() []byte {
	byte_s := []byte{0}
	return byte_s
}

type OpData interface {
	GetBytes() []byte
}

type UpgradeAccount struct {
	FeeData                 Amount     `json:"fee"`
	AccountToUpgrade        ObjectId   `json:"account_to_upgrade"`
	UpgradeToLifetimeMember bool       `json:"upgrade_to_lifetime_member"`
	ExtensionsData          Extensions `json:"extensions"`
}

func CreateUpgradeAccount(name string) *UpgradeAccount {
	info := rpc.GetAccountInfoByName(name)
	u := &UpgradeAccount{
		FeeData:                 Amount{Amount: 10000000, AssetID: "1.3.0"},
		ExtensionsData:          []interface{}{},
		AccountToUpgrade:        ObjectId(info.ID),
		UpgradeToLifetimeMember: true,
	}
	return u
}

func (o UpgradeAccount) GetBytes() []byte {
	fee_data := o.FeeData.GetBytes()
	atu_data := o.AccountToUpgrade.GetBytes()
	utlm_data := []byte{0}
	if o.UpgradeToLifetimeMember {
		utlm_data = []byte{1}
	}
	extensions_data := o.ExtensionsData.GetBytes()
	byte_s := append(fee_data,
		append(atu_data,
			append(utlm_data, extensions_data...)...)...)
	return byte_s
}

type KeyInfo struct {
	WeightThreshold int64         `json:"weight_threshold"`
	AccountAuths    []interface{} `json:"account_auths"`
	KeyAuths        [][]string    `json:"key_auths"`
	ExtensionsData  Extensions    `json:"extensions"`
}

func (o KeyInfo) GetBytes() []byte {
	wt_data := common.VarInt(o.WeightThreshold, 32)
	ka_data := []byte{byte(len(o.KeyAuths))}
	for i := 0; i < len(o.KeyAuths); i++ {
		tmp, _ := strconv.Atoi(o.KeyAuths[i][1])
		ka_data = append(ka_data, append(PukFromBase58String(o.KeyAuths[i][0]), common.VarInt(int64(tmp), 16)...)...)
	}
	aa_data := []byte{byte(len(o.AccountAuths))}
	extensions_data := o.ExtensionsData.GetBytes()
	byte_s := append(wt_data,
		append(aa_data,
			append(ka_data, extensions_data...)...)...)
	return byte_s
}

type Options struct {
	ExtensionsData Extensions    `json:"extensions"`
	NumCommittee   int64         `json:"num_committee"`
	MemoKey        string        `json:"memo_key"`
	Votes          []interface{} `json:"votes"`
	NumWitness     int64         `json:"num_witness"`
	VotingAccount  ObjectId      `json:"voting_account"`
}

func (o Options) GetBytes() []byte {
	nc_data := common.VarInt(o.NumCommittee, 16)
	nw_data := common.VarInt(o.NumWitness, 16)
	mk_data := PukFromBase58String(o.MemoKey)
	votes_data := []byte{0x00}
	va_data := o.VotingAccount.GetBytes()
	extensions_data := o.ExtensionsData.GetBytes()
	byte_s := append(mk_data,
		append(va_data,
			append(nw_data,
				append(nc_data,
					append(votes_data, extensions_data...)...)...)...)...)
	return byte_s
}

type name_string string

func (o name_string) GetBytes() []byte {
	byte_s := append([]byte{byte(len(o))}, []byte(o)...)
	return byte_s
}

type RegisterData struct {
	Referrer        ObjectId    `json:"referrer"`
	ExtensionsData  Extensions  `json:"extensions"`
	Active          KeyInfo     `json:"active"`
	OptionsData     Options     `json:"options"`
	FeeData         Amount      `json:"fee"`
	Owner           KeyInfo     `json:"owner"`
	ReferrerPercent int64       `json:"referrer_percent"`
	Name            name_string `json:"name"`
	Registrar       ObjectId    `json:"registrar"`
}

func (o RegisterData) GetBytes() []byte {
	fee_data := o.FeeData.GetBytes()
	referrer_data := o.Referrer.GetBytes()
	registrar_data := o.Registrar.GetBytes()
	referrer_percent_data := common.VarInt(o.ReferrerPercent, 16)
	name_data := o.Name.GetBytes()
	owner_data := o.Owner.GetBytes()
	active_data := o.Active.GetBytes()
	op_data := o.OptionsData.GetBytes()
	extensions_data := o.ExtensionsData.GetBytes()
	byte_s := append(fee_data,
		append(registrar_data,
			append(referrer_data,
				append(referrer_percent_data,
					append(name_data,
						append(owner_data,
							append(active_data,
								append(op_data, extensions_data...)...)...)...)...)...)...)...)
	return byte_s
}
func CreateRegisterData(active_PubKey, owner_PubKey, name, referrer, registrar string) *RegisterData {

	active_key := KeyInfo{
		ExtensionsData:  []interface{}{},
		AccountAuths:    []interface{}{},
		WeightThreshold: 1,
		KeyAuths:        [][]string{[]string{active_PubKey, "1"}},
	}
	owner_key := KeyInfo{
		ExtensionsData:  []interface{}{},
		AccountAuths:    []interface{}{},
		WeightThreshold: 1,
		KeyAuths:        [][]string{[]string{owner_PubKey, "1"}},
	}
	opData := Options{
		ExtensionsData: []interface{}{},
		NumWitness:     0,
		NumCommittee:   0,
		MemoKey:        active_PubKey,
		Votes:          []interface{}{},
		VotingAccount:  ObjectId(registrar),
	}
	r := &RegisterData{
		Referrer:        ObjectId(referrer),
		Registrar:       ObjectId(registrar),
		Name:            name_string(name),
		ExtensionsData:  []interface{}{},
		Active:          active_key,
		Owner:           owner_key,
		FeeData:         Amount{Amount: 5148, AssetID: "1.3.0"},
		ReferrerPercent: 5000,
		OptionsData:     opData,
	}
	return r
}

type Transaction struct {
	FeeData        Amount     `json:"fee"`
	From           ObjectId   `json:"from"`
	To             ObjectId   `json:"to"`
	AmountData     Amount     `json:"amount"`
	MemoData       Memo       `json:"memo"`
	ExtensionsData Extensions `json:"extensions"`
}

func CreateTransaction(prk *PrivateKey, from_name, to_name, tk_symbol string, value uint64) *Transaction {
	to_info := rpc.GetAccountInfoByName(to_name)
	to_puk := to_info.GetActivePuKey()
	from_info := rpc.GetAccountInfoByName(from_name)
	from_puk := from_info.GetActivePuKey()
	m_data := CreateMemo(prk, from_puk, to_puk, from_name)
	t := &Transaction{
		FeeData:        Amount{Amount: 20898, AssetID: "1.3.0"},
		AmountData:     Amount{Amount: value, AssetID: "1.3.28"},
		ExtensionsData: []interface{}{},
		From:           ObjectId(from_info.ID),
		To:             ObjectId(to_info.ID),
		MemoData:       m_data,
	}
	return t
}
func (o Transaction) GetBytes() []byte {
	fee_data := o.FeeData.GetBytes()
	from_data := o.From.GetBytes()
	to_data := o.To.GetBytes()
	amount_data := o.AmountData.GetBytes()
	memo_data := o.MemoData.GetBytes()
	extensions_data := o.ExtensionsData.GetBytes()
	byte_s := append(fee_data,
		append(from_data,
			append(to_data,
				append(amount_data,
					append(memo_data, extensions_data...)...)...)...)...)
	return byte_s
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
	block_prefix_data := common.VarUint(o.RefBlockPrefix, 32)
	t, _ := time.Parse(`2006-01-02T15:04:05`, o.Expiration)
	expiration_data := common.VarUint(uint64(t.Unix()), 32)
	operations_data := common.Varint(uint64(len(o.Operations)))
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

func CreateSignTransaction(opID int, prk *PrivateKey, t OpData) *Signed_Transaction {
	op := Operation{opID, t}
	dgp := chain.GetDynamicGlobalProperties()
	//time.Sleep(time.Second * 5)
	s := &Signed_Transaction{
		RefBlockNum:    dgp.Get_ref_block_num(),
		RefBlockPrefix: dgp.Get_ref_block_prefix(),
		Expiration:     time.Unix(time.Now().Unix(), 0).Format(`2006-01-02T15:04:05`),
		Operations:     []Operation{op},
		ExtensionsData: []interface{}{},
		Signatures:     []string{},
	}
	byte_s := s.GetBytes()
	cid, _ := hex.DecodeString(chain.CocosBCXChain.Properties.ChainID)
	byte_s = append(cid, byte_s...)
	msg := sha256digest(byte_s)
	//hex.DecodeString(chain.Chain.Properties.ChainID)
	s.Signatures = append(s.Signatures, prk.Sign(msg))
	return s
}

//func Create

func GetRequireFeeData(t *Transaction) *Amount {
	fee := &[]*Amount{}
	req := rpc.CreateRpcRequest(rpc.CALL,
		[]interface{}{0, `get_required_fees`,
			[]interface{}{[]interface{}{[]interface{}{0, t}}, "1.3.0"}})
	if resp, err := rpc.Client.Send(req); err == nil {
		if err = resp.GetInterface(fee); err == nil {
			return (*fee)[0]
		}
		return nil
	}

	return nil
}

func GetNonce() uint64 {
	rand.Seed(time.Now().Unix())
	return rand.Uint64()
}
