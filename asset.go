package CocosSDK

import (
	"cocos-go-sdk/rpc"
	. "cocos-go-sdk/type"
	"cocos-go-sdk/wallet"
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
		Fee:              EmptyFee(),
		FeePayingAccount: ObjectId(Wallet.Default.Info.ID),
		Extensions:       []interface{}{},
	}
	rpc.GetRequireFeeData(54, tx)
	st := wallet.CreateSignTransaction(54, Wallet.Default.GetActiveKey(), tx)
	return rpc.BroadcastTransaction(st)
}

/*取消 NH 资产卖出单*/
func CancelNhAssetOrder(order_id string) error {
	if Wallet.Default.Info == nil {
		Wallet.Default.Info = rpc.GetAccountInfoByName(Wallet.Default.Name)
	}
	tx := &CancelOrder{
		Order:            ObjectId(order_id),
		Fee:              EmptyFee(),
		FeePayingAccount: ObjectId(Wallet.Default.Info.ID),
		Extensions:       []interface{}{},
	}
	rpc.GetRequireFeeData(53, tx)
	st := wallet.CreateSignTransaction(53, Wallet.Default.GetActiveKey(), tx)
	return rpc.BroadcastTransaction(st)
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
		Fee:              EmptyFee(),
		Expiration:       GetExpiration(),
		Memo:             String(memo),
	}
	rpc.GetRequireFeeData(52, tx)
	st := wallet.CreateSignTransaction(52, Wallet.Default.GetActiveKey(), tx)
	return rpc.BroadcastTransaction(st)
}

/*NH 资产删除*/
func DeleteNhAsset(asset_id string) error {
	if Wallet.Default.Info == nil {
		Wallet.Default.Info = rpc.GetAccountInfoByName(Wallet.Default.Name)
	}
	tx := &DelNhAsset{
		NhAssetCreator: NhAssetCreator{
			FeePayingAccount: ObjectId(Wallet.Default.Info.ID),
			Fee:              EmptyFee()},
		NhAsset: ObjectId(asset_id),
	}
	rpc.GetRequireFeeData(50, tx)
	st := wallet.CreateSignTransaction(50, Wallet.Default.GetActiveKey(), tx)
	return rpc.BroadcastTransaction(st)
}

/*NH 资产转账*/
func TransferNhAsset(to_name, asset_id string) error {
	if Wallet.Default.Info == nil {
		Wallet.Default.Info = rpc.GetAccountInfoByName(Wallet.Default.Name)
	}
	to_info := rpc.GetAccountInfoByName(to_name)
	tx := &TransferNh{
		Fee:     EmptyFee(),
		To:      ObjectId(to_info.ID),
		From:    ObjectId(Wallet.Default.Info.ID),
		NhAsset: ObjectId(asset_id),
	}
	rpc.GetRequireFeeData(51, tx)
	st := wallet.CreateSignTransaction(51, Wallet.Default.GetActiveKey(), tx)
	return rpc.BroadcastTransaction(st)
}

/*創建NH資產*/
func CreateNhAsset(asset_symbol, world_view, owner_name, base_describe string) error {
	if Wallet.Default.Info == nil {
		Wallet.Default.Info = rpc.GetAccountInfoByName(Wallet.Default.Name)
	}
	owner_info := rpc.GetAccountInfoByName(owner_name)
	nh_asset := &NhAsset{
		Fee:              EmptyFee(),
		AssetID:          String(asset_symbol),
		BaseDescribe:     String(base_describe),
		Owner:            ObjectId(owner_info.ID),
		FeePayingAccount: ObjectId(Wallet.Default.Info.ID),
		WorldView:        String(world_view),
	}
	rpc.GetRequireFeeData(49, nh_asset)
	st := wallet.CreateSignTransaction(49, Wallet.Default.GetActiveKey(), nh_asset)
	return rpc.BroadcastTransaction(st)
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
	rpc.GetRequireFeeData(47, world_view)
	st := wallet.CreateSignTransaction(47, Wallet.Default.GetActiveKey(), world_view)
	return rpc.BroadcastTransaction(st)
}

/*创建 token*/
func CreateAsset(symbol, asset, _asset string, max_supply, precision, amount, _amount uint64) error {

	base := Amount{Amount: amount, AssetID: ObjectId(asset)}
	quote := Amount{Amount: _amount, AssetID: ObjectId(_asset)}
	if Wallet.Default.Info == nil {
		Wallet.Default.Info = rpc.GetAccountInfoByName(Wallet.Default.Name)
	}
	cm_op := CommonOptions{
		MaxSupply:            max_supply * precision,
		MarketFeePercent:     0,
		MaxMarketFee:         0,
		Flags:                0,
		IssuerPermissions:    79,
		CoreExchangeRateData: CoreExchangeRate{Base: base, Quote: quote},
		Description:          String(`{"main":"` + symbol + `","short_name":"","market":""}`),
		Extensions:           []interface{}{},
	}
	AssetData := &CreateAssetData{
		Fee:               EmptyFee(),
		Extensions:        []interface{}{},
		Precision:         precision,
		Issuer:            ObjectId(Wallet.Default.Info.ID),
		Symbol:            String(symbol),
		CommonOptionsData: cm_op,
	}
	rpc.GetRequireFeeData(8, AssetData)
	st := wallet.CreateSignTransaction(8, Wallet.Default.GetActiveKey(), AssetData)
	return rpc.BroadcastTransaction(st)
}

/*发币*/
func IssueToken(symbol, issue_to_account string, amount float64) {
	if Wallet.Default.Info == nil {
		Wallet.Default.Info = rpc.GetAccountInfoByName(Wallet.Default.Name)
	}
	to_info := rpc.GetAccountInfoByName(issue_to_account)
	asset_info := rpc.GetTokenInfoBySymbol(symbol)
	precision := math.Pow10(asset_info.Precision)
	issue := &IssueAsset{
		Fee:            EmptyFee(),
		Extensions:     []interface{}{},
		Issuer:         ObjectId(Wallet.Default.Info.ID),
		IssueToAccount: ObjectId(to_info.ID),
		AssetToIssue:   Amount{Amount: uint64(amount * precision), AssetID: ObjectId(asset_info.ID)},
	}
	rpc.GetRequireFeeData(13, issue)
	st := wallet.CreateSignTransaction(13, Wallet.Default.GetActiveKey(), issue)
	rpc.BroadcastTransaction(st)
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
