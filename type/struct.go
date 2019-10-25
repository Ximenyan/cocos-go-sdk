package types

import (
	"cocos-go-sdk/common"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"cocos-go-sdk/crypto/base58-go"
)

func PukBytesFromBase58String(base58Str string) []byte {
	byte_s, _ := base58.BitcoinEncoding.Decode([]byte(base58Str)[5:])
	big_i, _ := new(big.Int).SetString(string(byte_s), 10)
	data := big_i.Bytes()
	puk := data[0 : len(data)-4]
	return puk
}

type BigInt struct {
	big.Int
}

func (i *BigInt) UnmarshalJSON(data []byte) error {
	if data[0] == '"' {
		data = data[1 : len(data)-1]
	}
	big := new(big.Int)
	if err := json.Unmarshal(data, big); err != nil {
		return err
	}
	i.SetBytes(big.Bytes())
	return nil
}

type Object interface {
	GetBytes() []byte
}
type ObjectId string

func (o ObjectId) GetBytes() []byte {
	num := strings.Split(string(o), `.`)[2]
	i, _ := strconv.ParseUint(num, 10, 64)
	return common.Varint(i)
}

type Optional ObjectId

func (o Optional) GetBytes() []byte {
	if ObjectId(o) == EMPTY_ID {
		return []byte{0x0}
	} else {
		return append([]byte{0x1}, ObjectId(o).GetBytes()...)
	}
}

type Expiration string

func (o Expiration) GetBytes() []byte {
	t, _ := time.Parse(`2006-01-02T15:04:05`, string(o))
	expiration_data := common.VarUint(uint64(t.Unix()), 32)
	return expiration_data
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
	from := PukBytesFromBase58String(o.From)
	to := PukBytesFromBase58String(o.To)
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

type Extensions []interface{}

func (o Extensions) GetBytes() []byte {
	byte_s := []byte{0}
	return byte_s
}

type OpData interface {
	GetBytes() []byte
	SetFee(amount uint64)
}

/*手续费*/
type Fee struct {
	FeeData Amount `json:"fee"`
}

func (o *Fee) SetFee(amount uint64) {
	o.FeeData.Amount = amount
	return
}

type NhAssetCreator struct {
	Fee
	FeePayingAccount ObjectId `json:"fee_paying_account"`
}

func (o NhAssetCreator) GetBytes() []byte {
	fee_data := o.FeeData.GetBytes()
	fpa_data := o.FeePayingAccount.GetBytes()
	byte_s := append(fee_data, fpa_data...)
	return byte_s
}

type WorldView struct {
	Fee
	FeePayingAccount ObjectId `json:"fee_paying_account"`
	WorldView        String   `json:"world_view"`
}

func (o WorldView) GetBytes() []byte {
	fee_data := o.FeeData.GetBytes()
	fpa_data := o.FeePayingAccount.GetBytes()
	wv_data := o.WorldView.GetBytes()
	byte_s := append(fee_data,
		append(fpa_data, wv_data...)...)
	return byte_s
}

type ProposedOps struct {
	Fee
	RelatedAccount ObjectId `json:"related_account"`
	WorldView      String   `json:"world_view"`
	ViewOwner      ObjectId `json:"view_owner"`
}

func (o ProposedOps) GetBytes() []byte {
	fee_data := o.FeeData.GetBytes()
	ra_data := o.RelatedAccount.GetBytes()
	wv_data := o.WorldView.GetBytes()
	owner_data := o.ViewOwner.GetBytes()
	byte_s := append(fee_data,
		append(ra_data,
			append(wv_data, owner_data...)...)...)
	return byte_s
}

type OPS struct {
	ID  uint64
	Ops Object
}

func (o OPS) GetBytes() []byte {
	byte_s := append(
		common.Varint(o.ID), o.Ops.GetBytes()...)
	return byte_s
}

func (o OPS) MarshalJSON() ([]byte, error) {
	byte_s, _ := json.Marshal(o.Ops)
	return []byte(fmt.Sprintf(`{"op":[%d,%s]}`, o.ID, string(byte_s))), nil
}

type RelatedWorldView struct {
	Fee
	FeePayingAccount    ObjectId    `json:"fee_paying_account"`
	ExpirationTime      Expiration  `json:"expiration_time"`
	ProposedOps         []OPS       `json:"proposed_ops"`
	ReviewPeriodSeconds interface{} `json:"review_period_seconds,omitempty"`
	Extensions          Extensions  `json:"extensions"`
}

func (o RelatedWorldView) GetBytes() []byte {
	fee_data := o.FeeData.GetBytes()
	fpa_data := o.FeePayingAccount.GetBytes()
	exp_data := o.ExpirationTime.GetBytes()
	p_ops_data := common.Varint(uint64(len(o.ProposedOps)))
	for i := 0; i < len(o.ProposedOps); i++ {
		p_ops_data = append(p_ops_data, o.ProposedOps[i].GetBytes()...)
	}
	rps_data := []byte{0x0}
	ext_data := o.Extensions.GetBytes()
	byte_s := append(fee_data,
		append(fpa_data,
			append(exp_data,
				append(p_ops_data,
					append(rps_data, ext_data...)...)...)...)...)
	return byte_s
}

type Array []Object

func (o Array) GetBytes() []byte {
	byte_s := common.Varint(uint64(len(o)))
	for i := 0; i < len(o); i++ {
		byte_s = append(byte_s, o[i].GetBytes()...)
	}
	return byte_s
}

type Approvals struct {
	Fee
	FeePayingAccount        ObjectId   `json:"fee_paying_account"`
	Proposal                ObjectId   `json:"proposal"`
	ActiveApprovalsToAdd    Array      `json:"active_approvals_to_add"`
	ActiveApprovalsToRemove Array      `json:"active_approvals_to_remove"`
	OwnerApprovalsToAdd     Array      `json:"owner_approvals_to_add"`
	OwnerApprovalsToRemove  Array      `json:"owner_approvals_to_remove"`
	KeyApprovalsToAdd       Array      `json:"key_approvals_to_add"`
	KeyApprovalsToRemove    Array      `json:"key_approvals_to_remove"`
	Extensions              Extensions `json:"extensions"`
}

func (o Approvals) GetBytes() []byte {
	fee_data := o.FeeData.GetBytes()
	fpa_data := o.FeePayingAccount.GetBytes()
	proposal_data := o.Proposal.GetBytes()
	active_add_data := o.ActiveApprovalsToAdd.GetBytes()
	active_remove_data := o.ActiveApprovalsToRemove.GetBytes()
	owner_add_data := o.OwnerApprovalsToAdd.GetBytes()
	owner_remove_data := o.OwnerApprovalsToRemove.GetBytes()
	key_add_data := o.KeyApprovalsToAdd.GetBytes()
	key_remove_data := o.KeyApprovalsToRemove.GetBytes()
	extensions_data := o.Extensions.GetBytes()
	byte_s := append(fee_data,
		append(fpa_data,
			append(proposal_data,
				append(active_add_data,
					append(active_remove_data,
						append(owner_add_data,
							append(owner_remove_data,
								append(key_add_data,
									append(key_remove_data, extensions_data...)...)...)...)...)...)...)...)...)
	return byte_s
}

type NhAsset struct {
	Fee
	Owner            ObjectId `json:"owner"`
	BaseDescribe     String   `json:"base_describe"`
	AssetID          String   `json:"asset_id"`
	FeePayingAccount ObjectId `json:"fee_paying_account"`
	WorldView        String   `json:"world_view"`
}

func (o NhAsset) GetBytes() []byte {
	fee_data := o.FeeData.GetBytes()
	fpa_data := o.FeePayingAccount.GetBytes()
	owner_data := o.Owner.GetBytes()
	asset_data := o.AssetID.GetBytes()
	wv_data := o.WorldView.GetBytes()
	des_data := o.BaseDescribe.GetBytes()
	byte_s := append(fee_data,
		append(fpa_data,
			append(owner_data,
				append(asset_data,
					append(wv_data, des_data...)...)...)...)...)
	return byte_s
}

type TransferNh struct {
	Fee
	From    ObjectId `json:"from"`
	To      ObjectId `json:"to"`
	NhAsset ObjectId `json:"nh_asset"`
}

func (o TransferNh) GetBytes() []byte {
	fee_data := o.FeeData.GetBytes()
	from_data := o.From.GetBytes()
	to_data := o.To.GetBytes()
	nh_asset_data := o.NhAsset.GetBytes()
	byte_s := append(fee_data,
		append(from_data,
			append(to_data, nh_asset_data...)...)...)
	return byte_s
}

type DelNhAsset struct {
	NhAssetCreator
	NhAsset ObjectId `json:"nh_asset"`
}

func (o DelNhAsset) GetBytes() []byte {
	fee_data := o.FeeData.GetBytes()
	fpa_data := o.FeePayingAccount.GetBytes()
	nh_asset_data := o.NhAsset.GetBytes()
	byte_s := append(fee_data, append(fpa_data, nh_asset_data...)...)
	return byte_s
}

type NhOrder struct {
	Fee
	Seller           ObjectId   `json:"seller"`
	Otcaccount       ObjectId   `json:"otcaccount"`
	PendingOrdersFee Amount     `json:"pending_orders_fee"`
	NhAsset          ObjectId   `json:"nh_asset"`
	Memo             String     `json:"memo"`
	Price            Amount     `json:"price"`
	Expiration       Expiration `json:"expiration"`
}

func (o NhOrder) GetBytes() []byte {
	fee_data := o.FeeData.GetBytes()
	seller_data := o.Seller.GetBytes()
	otc_acc_data := o.Otcaccount.GetBytes()
	pof_data := o.PendingOrdersFee.GetBytes()
	nh_asset_data := o.NhAsset.GetBytes()
	memo_data := o.Memo.GetBytes()
	price_data := o.Price.GetBytes()
	expiration_data := o.Expiration.GetBytes()
	byte_s := append(fee_data,
		append(seller_data,
			append(otc_acc_data,
				append(pof_data,
					append(nh_asset_data,
						append(memo_data,
							append(price_data, expiration_data...)...)...)...)...)...)...)
	return byte_s
}

type FillNhOrder struct {
	Fee
	Order            ObjectId   `json:"order"`
	FeePayingAccount ObjectId   `json:"fee_paying_account"`
	Seller           ObjectId   `json:"seller"`
	NhAsset          ObjectId   `json:"nh_asset"`
	PriceAmount      String     `json:"price_amount"`
	PriceAssetID     ObjectId   `json:"price_asset_id"`
	PriceAssetSymbol String     `json:"price_asset_symbol"`
	Extensions       Extensions `json:"extensions"`
}

func (o FillNhOrder) GetBytes() []byte {
	fee_data := o.FeeData.GetBytes()
	order_data := o.Order.GetBytes()
	fpa_data := o.FeePayingAccount.GetBytes()
	seller_data := o.Seller.GetBytes()
	nh_asset_data := o.NhAsset.GetBytes()
	price_amount_data := o.PriceAmount.GetBytes()
	price_id_data := o.PriceAssetID.GetBytes()
	price_symbol_data := o.PriceAssetSymbol.GetBytes()
	extensions_data := o.Extensions.GetBytes()
	byte_s := append(fee_data,
		append(order_data,
			append(fpa_data,
				append(seller_data,
					append(nh_asset_data,
						append(price_amount_data,
							append(price_id_data,
								append(price_symbol_data, extensions_data...)...)...)...)...)...)...)...)
	//fmt.Println("CancelOrder::", len(byte_s))
	return byte_s
}

type CancelOrder struct {
	Fee
	Order            ObjectId   `json:"order"`
	FeePayingAccount ObjectId   `json:"fee_paying_account"`
	Extensions       Extensions `json:"extensions"`
}

func (o CancelOrder) GetBytes() []byte {
	fee_data := o.FeeData.GetBytes()
	order_data := o.Order.GetBytes()
	fpa_data := o.FeePayingAccount.GetBytes()
	extensions_data := o.Extensions.GetBytes()
	byte_s := append(fee_data,
		append(order_data,
			append(fpa_data, extensions_data...)...)...)
	return byte_s
}

type UpgradeAccount struct {
	Fee
	AccountToUpgrade        ObjectId   `json:"account_to_upgrade"`
	UpgradeToLifetimeMember bool       `json:"upgrade_to_lifetime_member"`
	ExtensionsData          Extensions `json:"extensions"`
}

func CreateUpgradeAccount(name string, account_id string) *UpgradeAccount {
	u := &UpgradeAccount{
		ExtensionsData:          []interface{}{},
		AccountToUpgrade:        ObjectId(account_id),
		UpgradeToLifetimeMember: true,
	}
	u.FeeData = Amount{Amount: 0, AssetID: COCOS_ID}
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
		ka_data = append(ka_data, append(PukBytesFromBase58String(o.KeyAuths[i][0]), common.VarInt(int64(tmp), 16)...)...)
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
	mk_data := PukBytesFromBase58String(o.MemoKey)
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

type String string

func (o String) GetBytes() []byte {
	byte_s := append([]byte{byte(len(o))}, []byte(o)...)
	return byte_s
}

type RegisterData struct {
	Fee
	Referrer        ObjectId   `json:"referrer"`
	ExtensionsData  Extensions `json:"extensions"`
	Active          KeyInfo    `json:"active"`
	OptionsData     Options    `json:"options"`
	Owner           KeyInfo    `json:"owner"`
	ReferrerPercent int64      `json:"referrer_percent"`
	Name            String     `json:"name"`
	Registrar       ObjectId   `json:"registrar"`
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
		Name:            String(name),
		ExtensionsData:  []interface{}{},
		Active:          active_key,
		Owner:           owner_key,
		ReferrerPercent: 5000,
		OptionsData:     opData,
	}
	r.FeeData = Amount{Amount: 0, AssetID: "1.3.0"}
	return r
}

type CoreExchangeRate struct {
	Base  Amount `json:"base"`
	Quote Amount `json:"quote"`
}

func (o CoreExchangeRate) GetBytes() []byte {
	byte_s := append(o.Base.GetBytes(), o.Quote.GetBytes()...)
	return byte_s
}

type CommonOptions struct {
	MaxSupply            uint64           `json:"max_supply"`
	MarketFeePercent     uint64           `json:"market_fee_percent"`
	MaxMarketFee         uint64           `json:"max_market_fee"`
	IssuerPermissions    uint64           `json:"issuer_permissions"`
	Flags                uint64           `json:"flags"`
	CoreExchangeRateData CoreExchangeRate `json:"core_exchange_rate"`
	Description          String           `json:"description"`
	Extensions           Extensions       `json:"extensions"`
}

func (o CommonOptions) GetBytes() []byte {
	MaxSupply_data := common.VarUint(o.MaxSupply, 64)
	MarketFeePercent_data := common.VarUint(o.MarketFeePercent, 16)
	MaxMarketFee_data := common.VarUint(o.MaxMarketFee, 64)
	IssuerPermissions_data := common.VarUint(o.IssuerPermissions, 16)
	Flags_data := common.VarUint(o.Flags, 16)
	CoreExchangeRate_data := o.CoreExchangeRateData.GetBytes()
	des_data := o.Description.GetBytes()
	extensions_data := o.Extensions.GetBytes()
	byte_s := append(MaxSupply_data,
		append(MarketFeePercent_data,
			append(MaxMarketFee_data,
				append(IssuerPermissions_data,
					append(Flags_data,
						append(CoreExchangeRate_data,
							append(des_data, extensions_data...)...)...)...)...)...)...)
	return byte_s
}

/*创建代币的数据结构*/
type CreateAssetData struct {
	Fee
	Issuer            ObjectId      `json:"issuer"`
	Symbol            String        `json:"symbol"`
	Precision         uint64        `json:"precision"`
	CommonOptionsData CommonOptions `json:"common_options"`
	Extensions        Extensions    `json:"extensions"`
}

func (o CreateAssetData) GetBytes() []byte {
	fee_data := o.FeeData.GetBytes()
	issuer_data := o.Issuer.GetBytes()
	symbol_data := o.Symbol.GetBytes()
	precision_data := common.VarUint(o.Precision, 8)
	cod_data := o.CommonOptionsData.GetBytes()
	bo_data := common.VarUint(0, 8)
	extensions_data := o.Extensions.GetBytes()
	byte_s := append(fee_data,
		append(issuer_data,
			append(symbol_data,
				append(precision_data,
					append(cod_data,
						append(bo_data, extensions_data...)...)...)...)...)...)
	return byte_s
}

/*创建代币的数据结构*/
type UpdateAssetData struct {
	Fee
	AssetToUpdate  ObjectId      `json:"asset_to_update"`
	Issuer         ObjectId      `json:"issuer"`
	NewIssuer      Optional      `json:"new_issuer,omitempty"`
	NewOptionsData CommonOptions `json:"new_options"`
	Extensions     Extensions    `json:"extensions"`
}

func (o UpdateAssetData) GetBytes() []byte {
	fee_data := o.FeeData.GetBytes()
	issuer_data := o.Issuer.GetBytes()
	new_issuer_data := o.NewIssuer.GetBytes()
	asset_data := o.AssetToUpdate.GetBytes()
	cod_data := o.NewOptionsData.GetBytes()
	extensions_data := o.Extensions.GetBytes()
	byte_s := append(fee_data,
		append(issuer_data,
			append(asset_data,
				append(new_issuer_data,
					append(cod_data, extensions_data...)...)...)...)...)
	return byte_s
}

/*创建發行代币的数据结构*/
type IssueAsset struct {
	Fee
	Issuer         ObjectId   `json:"issuer"`
	AssetToIssue   Amount     `json:"asset_to_issue"`
	IssueToAccount ObjectId   `json:"issue_to_account"`
	Extensions     Extensions `json:"extensions"`
}

func (o IssueAsset) GetBytes() []byte {
	byte_s := append(o.FeeData.GetBytes(),
		append(o.Issuer.GetBytes(),
			append(o.AssetToIssue.GetBytes(),
				append(o.IssueToAccount.GetBytes(),
					append([]byte{0x0},
						o.Extensions.GetBytes()...)...)...)...)...)
	return byte_s
}

type ReserveTokenData struct {
	Extensions      Extensions `json:"extensions"`
	Payer           ObjectId   `json:"payer"`
	AmountToReserve Amount     `json:"amount_to_reserve"`
	Fee
}

func (o ReserveTokenData) GetBytes() []byte {
	byte_s := append(o.FeeData.GetBytes(),
		append(o.Payer.GetBytes(),
			append(o.AmountToReserve.GetBytes(),
				o.Extensions.GetBytes()...)...)...)
	return byte_s
}

type ClaimTokenFees struct {
	Fee
	Issuer        ObjectId   `json:"issuer"`
	AmountToClaim Amount     `json:"amount_to_claim"`
	Extensions    Extensions `json:"extensions"`
}

func (o ClaimTokenFees) GetBytes() []byte {
	byte_s := append(o.FeeData.GetBytes(),
		append(o.Issuer.GetBytes(),
			append(o.AmountToClaim.GetBytes(),
				o.Extensions.GetBytes()...)...)...)
	return byte_s
}

type TokenFeePoolData struct {
	AssetID ObjectId `json:"asset_id"`
	Fee
	FromAccount ObjectId   `json:"from_account"`
	Extensions  Extensions `json:"extensions"`
	Amount      uint64     `json:"amount"`
}

func (o TokenFeePoolData) GetBytes() []byte {
	byte_s := append(o.FeeData.GetBytes(),
		append(o.FromAccount.GetBytes(),
			append(o.AssetID.GetBytes(),
				append(common.VarUint(o.Amount, 64),
					o.Extensions.GetBytes()...)...)...)...)
	return byte_s
}

type Transaction struct {
	Fee
	From           ObjectId   `json:"from"`
	To             ObjectId   `json:"to"`
	AmountData     Amount     `json:"amount"`
	MemoData       Memo       `json:"memo"`
	ExtensionsData Extensions `json:"extensions"`
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

type CallData struct {
	Fee
	Caller       ObjectId   `json:"caller"`
	ContractID   ObjectId   `json:"contract_id"`
	FunctionName String     `json:"function_name"`
	ValueList    ValueList  `json:"value_list"`
	Extensions   Extensions `json:"extensions"`
}

func (o CallData) GetBytes() []byte {
	fee_data := o.FeeData.GetBytes()
	value_list_data := o.ValueList.GetBytes()
	caller_data := o.Caller.GetBytes()
	func_name_data := o.FunctionName.GetBytes()
	contract_id_data := o.ContractID.GetBytes()
	extensions_data := o.Extensions.GetBytes()
	byte_s := append(fee_data,
		append(caller_data,
			append(contract_id_data,
				append(func_name_data,
					append(value_list_data, extensions_data...)...)...)...)...)
	return byte_s
}

type CreateContractData struct {
	Fee
	Extensions        Extensions `json:"extensions"`
	Owner             ObjectId   `json:"owner"`
	Name              String     `json:"name"`
	ContractAuthority string     `json:"contract_authority"`
	Data              String     `json:"data"`
}

func (o CreateContractData) GetBytes() []byte {
	fee_data := o.FeeData.GetBytes()
	c_auth_data := PukBytesFromBase58String(o.ContractAuthority)
	owner_data := o.Owner.GetBytes()
	data_data := o.Data.GetBytes()
	name_data := o.Name.GetBytes()
	extensions_data := o.Extensions.GetBytes()
	byte_s := append(fee_data,
		append(owner_data,
			append(name_data,
				append(data_data,
					append(c_auth_data, extensions_data...)...)...)...)...)
	return byte_s
}

type UpdateContractData struct {
	Fee
	Reviser    ObjectId   `json:"reviser"`
	ContractID ObjectId   `json:"contract_id"`
	Extensions Extensions `json:"extensions"`
	Data       String     `json:"data"`
}

func (o UpdateContractData) GetBytes() []byte {
	fee_data := o.FeeData.GetBytes()
	contract_id_data := o.ContractID.GetBytes()
	reviser_data := o.Reviser.GetBytes()
	data_data := o.Data.GetBytes()
	extensions_data := o.Extensions.GetBytes()
	byte_s := append(fee_data,
		append(reviser_data,
			append(contract_id_data,
				append(data_data, extensions_data...)...)...)...)
	return byte_s
}

type Int uint64

func (o Int) GetBytes() []byte {
	return common.Varint(uint64(o))
}

type Policy struct {
	ID             uint64
	StartClaim     Expiration
	VestingSeconds uint64
}

func (o Policy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`[%d,{"start_claim":"%s","vesting_seconds":%d}]`, o.ID, o.StartClaim, o.VestingSeconds)), nil
}

func (o Policy) GetBytes() []byte {
	byte_s := append(common.Varint(o.ID),
		append(o.StartClaim.GetBytes(),
			common.VarUint(o.VestingSeconds, 32)...)...)
	return byte_s
}

type VestingBalanceCreate struct {
	Policy Policy   `json:"policy"`
	Owner  ObjectId `json:"owner"`
	Amount Amount   `json:"amount"`
	Fee
	Creator ObjectId `json:"creator"`
}

func (o VestingBalanceCreate) GetBytes() []byte {
	byte_s := append(o.FeeData.GetBytes(),
		append(o.Creator.GetBytes(),
			append(o.Owner.GetBytes(),
				append(o.Amount.GetBytes(),
					o.Policy.GetBytes()...)...)...)...)
	return byte_s
}

type VestingBalanceWithdraw struct {
	Fee
	Owner          ObjectId `json:"owner"`
	Amount         Amount   `json:"amount"`
	VestingBalance ObjectId `json:"vesting_balance"`
}

func (o VestingBalanceWithdraw) GetBytes() []byte {
	byte_s := append(o.FeeData.GetBytes(),
		append(o.VestingBalance.GetBytes(),
			append(o.Owner.GetBytes(),
				o.Amount.GetBytes()...)...)...)
	return byte_s
}
