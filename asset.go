package CocosSDK

import (
	"cocos-go-sdk/rpc"
	. "cocos-go-sdk/type"
	"math"
)

/*吃单 NH 资产买入*/
func FillNhAsset(order_id string) error {
	if Wallet.Default.Info == nil {
		Wallet.Default.Info = rpc.GetAccountInfoByName(Wallet.Default.Name)
	}
	order_info := rpc.GetNhAssetOrderInfo(order_id)
	tx := &FillNhOrder{
		PriceAmount:      String(order_info.Price.Amount.String()),
		PriceAssetID:     ObjectId(order_info.Price.AssetID),
		NhAsset:          ObjectId(order_info.NhAssetID),
		Seller:           ObjectId(order_info.Seller),
		PriceAssetSymbol: String(order_info.AssetQualifier),
		Order:            ObjectId(order_id),
		FeePayingAccount: ObjectId(Wallet.Default.Info.ID),
		Extensions:       []interface{}{},
	}
	return Wallet.SignAndSendTX(OP_FILL_NHORDER, tx)
}

/*取消 NH 资产卖出单*/
func CancelNhAssetOrder(order_id string) error {
	if Wallet.Default.Info == nil {
		Wallet.Default.Info = rpc.GetAccountInfoByName(Wallet.Default.Name)
	}
	tx := &CancelOrder{
		Order:            ObjectId(order_id),
		FeePayingAccount: ObjectId(Wallet.Default.Info.ID),
		Extensions:       []interface{}{},
	}
	return Wallet.SignAndSendTX(OP_CANCEL_NH_ORDER, tx)
}

/*NH 资产卖出单*/
func SellNhAsset(otcaccount_name, asset_id, memo, pending_order_fee_asset, price_asset string, pending_order_fee_amount, price_amount uint64) error {
	if Wallet.Default.Info == nil {
		Wallet.Default.Info = rpc.GetAccountInfoByName(Wallet.Default.Name)
	}
	otcaccount_info := rpc.GetAccountInfoByName(otcaccount_name)
	pending_asset_info := rpc.GetTokenInfo(pending_order_fee_asset)
	price_asset_info := rpc.GetTokenInfo(price_asset)
	pending_precision := math.Pow10(pending_asset_info.Precision)
	price_precision := math.Pow10(price_asset_info.Precision)
	tx := &NhOrder{
		NhAsset: ObjectId(asset_id),

		PendingOrdersFee: Amount{Amount: uint64(float64(pending_order_fee_amount) * pending_precision), AssetID: ObjectId(pending_order_fee_asset)},
		Price:            Amount{Amount: uint64(float64(price_amount) * price_precision), AssetID: ObjectId(price_asset)},
		Seller:           ObjectId(Wallet.Default.Info.ID),
		Otcaccount:       ObjectId(otcaccount_info.ID),
		Expiration:       GetExpiration(),
		Memo:             String(memo),
	}
	return Wallet.SignAndSendTX(OP_SELL_NH_ASSET, tx)
}

/*NH 资产删除*/
func DeleteNhAsset(asset_id string) error {
	if Wallet.Default.Info == nil {
		Wallet.Default.Info = rpc.GetAccountInfoByName(Wallet.Default.Name)
	}
	tx := &DelNhAsset{
		NhAssetCreator: NhAssetCreator{
			FeePayingAccount: ObjectId(Wallet.Default.Info.ID)},
		NhAsset: ObjectId(asset_id),
	}
	return Wallet.SignAndSendTX(OP_DEL_NH_ASSET, tx)
}

/*NH 资产转账*/
func TransferNhAsset(to_name, asset_id string) error {
	if Wallet.Default.Info == nil {
		Wallet.Default.Info = rpc.GetAccountInfoByName(Wallet.Default.Name)
	}
	to_info := rpc.GetAccountInfoByName(to_name)
	tx := &TransferNh{
		To:      ObjectId(to_info.ID),
		From:    ObjectId(Wallet.Default.Info.ID),
		NhAsset: ObjectId(asset_id),
	}
	//rpc.GetRequireFeeData(51, tx)
	//st := wallet.CreateSignTransaction(51, Wallet.Default.GetActiveKey(), tx)
	return Wallet.SignAndSendTX(OP_TRANSFER_NH_ASSET, tx)
}

/*創建NH資產*/
func CreateNhAsset(asset_symbol, world_view, owner_name, base_describe string) error {
	if Wallet.Default.Info == nil {
		Wallet.Default.Info = rpc.GetAccountInfoByName(Wallet.Default.Name)
	}
	owner_info := rpc.GetAccountInfoByName(owner_name)
	nh_asset := &NhAsset{
		//Fee:              EmptyFee(),
		AssetID:          String(asset_symbol),
		BaseDescribe:     String(base_describe),
		Owner:            ObjectId(owner_info.ID),
		FeePayingAccount: ObjectId(Wallet.Default.Info.ID),
		WorldView:        String(world_view),
	}
	return Wallet.SignAndSendTX(OP_CREATE_NH_ASSET, nh_asset)
}

/*批准 关联世界观的提议*/
func ApprovalsProposal(proposal_id string) error {
	if Wallet.Default.Info == nil {
		Wallet.Default.Info = rpc.GetAccountInfoByName(Wallet.Default.Name)
	}
	approval := &Approvals{
		FeePayingAccount:        ObjectId(Wallet.Default.Info.ID),
		Proposal:                ObjectId(proposal_id),
		ActiveApprovalsToAdd:    []Object{ObjectId(Wallet.Default.Info.ID)},
		ActiveApprovalsToRemove: []Object{},
		OwnerApprovalsToAdd:     []Object{},
		OwnerApprovalsToRemove:  []Object{},
		KeyApprovalsToAdd:       []Object{},
		KeyApprovalsToRemove:    []Object{},
		Extensions:              []interface{}{},
	}
	return Wallet.SignAndSendTX(OP_APPROVAL, approval)
}

/*提议关联世界观*/
func RelateWorldView(world_view string) error {
	if Wallet.Default.Info == nil {
		Wallet.Default.Info = rpc.GetAccountInfoByName(Wallet.Default.Name)
	}
	world_view_info := rpc.GetWorldViewInfo(world_view)
	creator := rpc.GetWorldViewCreator(world_view_info.WorldViewCreator)
	op_data := &ProposedOps{
		RelatedAccount: ObjectId(Wallet.Default.Info.ID),
		WorldView:      String(world_view),
		ViewOwner:      ObjectId(creator.Creator),
	}
	ops := OPS{
		ID:  OP_RELATE_WORLDVIEW,
		Ops: *op_data,
	}
	op := &RelatedWorldView{
		FeePayingAccount: ObjectId(Wallet.Default.Info.ID),
		ExpirationTime:   GetExpiration(),
		ProposedOps:      []OPS{ops},
		Extensions:       []interface{}{},
	}
	//fees := rpc.GetRequireFeeData(21, op)
	return Wallet.SignAndSendTX(OP_PROPOSAL, op)
}

/*創建世界觀*/
func CreateWorldView(name string) error {
	if Wallet.Default.Info == nil {
		Wallet.Default.Info = rpc.GetAccountInfoByName(Wallet.Default.Name)
	}
	world_view := &WorldView{
		Fee:              EmptyFee(),
		FeePayingAccount: ObjectId(Wallet.Default.Info.ID),
		WorldView:        String(name),
	}
	return Wallet.SignAndSendTX(OP_CREATE_WORLDVIEW, world_view)
}

/*更新 token*/
func UpdateToken(symbol string, max_supply, precision uint64, new_issuer ...string) error {

	//base := Amount{Amount: amount, AssetID: ObjectId(asset)}
	//quote := Amount{Amount: _amount, AssetID: ObjectId(_asset)}
	update_asset_info := rpc.GetTokenInfoBySymbol(symbol)

	if Wallet.Default.Info == nil {
		Wallet.Default.Info = rpc.GetAccountInfoByName(Wallet.Default.Name)
	}
	precision = uint64(math.Pow10(int(precision)))
	cm_op := CommonOptions{
		MaxSupply:         max_supply * precision,
		MarketFeePercent:  0,
		MaxMarketFee:      0,
		Flags:             0,
		IssuerPermissions: 79,
		//CoreExchangeRateData: CoreExchangeRate{Base: base, Quote: quote},
		Description: String(`{"main":"` + symbol + `","short_name":"","market":""}`),
		Extensions:  []interface{}{},
	}
	var newIssuer ObjectId
	if len(new_issuer) >= 1 {
		new_issuer_info := rpc.GetAccountInfoByName(new_issuer[0])
		newIssuer = ObjectId(new_issuer_info.ID)
	} else {
		newIssuer = EMPTY_ID
	}
	AssetData := &UpdateAssetData{
		Fee:            EmptyFee(),
		Extensions:     []interface{}{},
		NewIssuer:      Optional(newIssuer),
		Issuer:         ObjectId(Wallet.Default.Info.ID),
		AssetToUpdate:  ObjectId(update_asset_info.ID),
		NewOptionsData: cm_op,
	}
	return Wallet.SignAndSendTX(OP_UPDATE_TOKEN, AssetData)
}

/*销毁 token*/
func ReserveToken(symbol string, amount float64) error {
	asset_info := rpc.GetTokenInfoBySymbol(symbol)
	precision := math.Pow10(asset_info.Precision)

	if Wallet.Default.Info == nil {
		Wallet.Default.Info = rpc.GetAccountInfoByName(Wallet.Default.Name)
	}
	AssetData := &ReserveTokenData{
		Extensions:      []interface{}{},
		Payer:           ObjectId(Wallet.Default.Info.ID),
		AmountToReserve: Amount{Amount: uint64(float64(amount) * precision), AssetID: ObjectId(asset_info.ID)},
		Fee:             EmptyFee(),
	}
	return Wallet.SignAndSendTX(OP_RESERVE_TOKEN, AssetData)
}

/*创建 token*/
func CreateToken(symbol string, max_supply, precision uint64) error {

	//base := Amount{Amount: amount, AssetID: ObjectId(asset)}
	//quote := Amount{Amount: _amount, AssetID: ObjectId(_asset)}
	if Wallet.Default.Info == nil {
		Wallet.Default.Info = rpc.GetAccountInfoByName(Wallet.Default.Name)
	}
	new_precision := uint64(math.Pow10(int(precision)))

	cm_op := CommonOptions{
		MaxSupply:         max_supply * new_precision,
		MarketFeePercent:  0,
		MaxMarketFee:      0,
		Flags:             0,
		IssuerPermissions: 79,
		//CoreExchangeRateData: CoreExchangeRate{Base: base, Quote: quote},
		Description: String(`{"main":"` + symbol + `","short_name":"","market":""}`),
		Extensions:  []interface{}{},
	}
	AssetData := &CreateAssetData{
		Extensions:        []interface{}{},
		Precision:         precision,
		Issuer:            ObjectId(Wallet.Default.Info.ID),
		Symbol:            String(symbol),
		CommonOptionsData: cm_op,
	}
	return Wallet.SignAndSendTX(OP_CREATE_ASSET_TOKEN, AssetData)
}

/*发行人 可以领取累计的手续费*/
func ClaimFees(symbol string, value float64) error {
	asset_info := rpc.GetTokenInfoBySymbol(symbol)
	precision := math.Pow10(asset_info.Precision)
	ctf := &ClaimTokenFees{
		Extensions:    []interface{}{},
		Issuer:        ObjectId(Wallet.Default.Info.ID),
		AmountToClaim: Amount{Amount: uint64(float64(value) * precision), AssetID: ObjectId(asset_info.ID)},
		Fee:           EmptyFee(),
	}
	return Wallet.SignAndSendTX(OP_CLAIM_FEES, ctf)
}

/*注资手续费池*/
func TokenFundFeePool(symbol string, amount float64) error {
	asset_info := rpc.GetTokenInfoBySymbol(symbol)
	precision := math.Pow10(asset_info.Precision)
	if Wallet.Default.Info == nil {
		Wallet.Default.Info = rpc.GetAccountInfoByName(Wallet.Default.Name)
	}
	feePool := &TokenFeePoolData{
		AssetID:     ObjectId(asset_info.ID),
		FromAccount: ObjectId(Wallet.Default.Info.ID),
		Amount:      uint64(float64(amount) * precision),
		Extensions:  []interface{}{},
	}
	return Wallet.SignAndSendTX(OP_FUND_FEEPOOL, feePool)
}

func Pledgegas(mortgager, beneficiary string, collateral float64) error {
	m_info := rpc.GetAccountInfoByName(mortgager)
	b_info := rpc.GetAccountInfoByName(beneficiary)
	tk_info := rpc.GetTokenInfo(COCOS_ID)
	precision := math.Pow10(tk_info.Precision)
	p := &PledgeGas{
		Mortgager:   ObjectId(m_info.ID),
		Beneficiary: ObjectId(b_info.ID),
		Collateral:  uint64(collateral * precision),
	}
	return Wallet.SignAndSendTX(OP_PLEDGE_GAS, p)
}

func CreateVestingBalance(symbol string, amount float64) error {
	if Wallet.Default.Info == nil {
		Wallet.Default.Info = rpc.GetAccountInfoByName(Wallet.Default.Name)
	}
	asset_info := rpc.GetTokenInfoBySymbol(symbol)
	precision := math.Pow10(asset_info.Precision)
	p := Policy{
		ID:             1,
		StartClaim:     GetExpiration(),
		VestingSeconds: 0,
	}
	v := &VestingBalanceCreate{
		Owner:   ObjectId(Wallet.Default.Info.ID),
		Amount:  Amount{Amount: uint64(amount * precision), AssetID: asset_info.ID},
		Policy:  p,
		Creator: ObjectId(Wallet.Default.Info.ID),
	}
	return Wallet.SignAndSendTX(OP_VESTING_CREATE, v)
}

func WithdrawVestingBalance(balance_id string) error {
	if Wallet.Default.Info == nil {
		Wallet.Default.Info = rpc.GetAccountInfoByName(Wallet.Default.Name)
	}
	balances := GetVestingBalances(Wallet.Default.Name)
	var balance_info rpc.VestingBalances
	for _, balance := range balances {
		if balance.ID == balance_id {
			balance_info = balance
			break
		}
	}
	v := &VestingBalanceWithdraw{
		VestingBalance: ObjectId(balance_id),
		Owner:          ObjectId(Wallet.Default.Info.ID),
		Amount:         Amount{AssetID: ObjectId(balance_info.Balance.AssetID), Amount: balance_info.GetBalanceAmount()},
	}
	return Wallet.SignAndSendTX(OP_VESTING_WITHDRAW, v)
}

/*发币*/
func IssueToken(symbol, issue_to_account string, amount float64) error {
	if Wallet.Default.Info == nil {
		Wallet.Default.Info = rpc.GetAccountInfoByName(Wallet.Default.Name)
	}
	to_info := rpc.GetAccountInfoByName(issue_to_account)
	asset_info := rpc.GetTokenInfoBySymbol(symbol)
	precision := math.Pow10(asset_info.Precision)
	issue := &IssueAsset{
		//Fee:            EmptyFee(),
		Extensions:     []interface{}{},
		Issuer:         ObjectId(Wallet.Default.Info.ID),
		IssueToAccount: ObjectId(to_info.ID),
		AssetToIssue:   Amount{Amount: uint64(amount * precision), AssetID: ObjectId(asset_info.ID)},
	}
	return Wallet.SignAndSendTX(OP_ISSUE_TOKEN, issue)
}

func Vote(id string, value float64) error {
	tk_info := rpc.GetTokenInfo(COCOS_ID)
	precision := math.Pow10(tk_info.Precision)
	info := rpc.GetObject(id)

	if Wallet.Default.Info == nil {
		Wallet.Default.Info = rpc.GetAccountInfoByName(Wallet.Default.Name)
	}
	v := &VoteData{
		LockWithVote: OPArray{Int32(0), Amount{Amount: uint64(value * precision), AssetID: COCOS_ID}},
		Account:      ObjectId(Wallet.Default.Info.ID),
		NewOptions: NewOptions{
			MemoKey:    Wallet.Default.GetMemoKey().GetPublicKey().ToBase58String(),
			Votes:      Array{VoteId(info.Get("vote_id").String())},
			Extensions: Extensions{}},
		Extensions: Extensions{},
	}
	return Wallet.SignAndSendTX(OP_VOTE, v, Wallet.Default.GetActiveKey(), Wallet.Default.GetOwnerKey())
}

/*查询订单信息*/
func GetNhAssetOrderInfo(id string) *rpc.NhAssetOrderInfo {
	return rpc.GetNhAssetOrderInfo(id)
}

/*查询owner订单列表*/
func GetAccountNhAssetOrderList(owner_name string, page, page_size int) *rpc.OrdersList {
	return rpc.GetAccountNhAssetOrderList(owner_name, page, page_size)
}

/*查询订单列表*/
func GetNhAssetOrderList(asset_name, world_view string, page, page_size int) *rpc.OrdersList {
	return rpc.GetNhAssetOrderList(asset_name, world_view, page, page_size)
}

/*查询NH 资产列表*/
func GetNhAssetList(acc_name string, page, page_size, _type int, world_view ...string) *rpc.AssetsList {
	return rpc.GetNhAssetList(acc_name, page, page_size, _type, world_view)
}

/*查询账户Balance*/
func GetAccountBalances(acc_name string) *[]rpc.Balance {
	acc_info := rpc.GetAccountInfoByName(acc_name)
	return rpc.GetAccountBalances(acc_info.ID)
}

/*查询链上所有token信息*/
func GetAllTokenInfo() []*rpc.TokenInfo {
	return rpc.QueryTokenList()
}

/*查询收到的所有提议*/
func GetAllProposals(acct_id string) *[]rpc.Proposal {
	return rpc.GetProposals(acct_id)
}

/*查询 某条提议*/
func GetAllProposal(proposal_id string) *[]rpc.Proposal {
	return rpc.GetProposals(proposal_id)
}

/*通过Symbol查询token信息*/
func GetTokenInfoBySymbol(symbol string) *rpc.TokenInfo {
	return rpc.GetTokenInfoBySymbol(symbol)
}

/*通过id查询token信息*/
func GetTokenInfoById(id string) *rpc.TokenInfo {
	return rpc.GetTokenInfo(id)
}

/*查询账户待提取的奖励*/
func GetVestingBalances(acct_name string) []rpc.VestingBalances {
	return rpc.GetVestingBalancesByName(acct_name)
}

/*查询账户操作记录*/
func GetAccountHistorys(acct_name string) rpc.History {
	return rpc.GetAccountHistory(acct_name)
}

/*获取市场限价单交易历史*/
func GetFillOrderHistory(asset_id, _asset_id string, limit uint64) []interface{} {
	return rpc.GetFillOrderHistory(asset_id, _asset_id, limit)
}

/*查询某个时间段的交易市场行情。*/
func GetMarketHistory(asset_id, _asset_id, start, end string, limit uint64) []interface{} {
	return rpc.GetMarketHistory(asset_id, _asset_id, start, end, limit)
}
