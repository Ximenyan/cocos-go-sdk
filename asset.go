package CocosSDK

import (
	"cocos-go-sdk/rpc"
	. "cocos-go-sdk/type"
	"cocos-go-sdk/wallet"
	"math"
)

/*创建 token*/
func CreateAsset(symbol, asset, _asset string, max_supply, precision, amount, _amount uint64) {
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
		Extensions:        []interface{}{},
		Precision:         precision,
		Issuer:            ObjectId(Wallet.Default.Info.ID),
		Symbol:            String(symbol),
		CommonOptionsData: cm_op,
	}
	AssetData.FeeData = Amount{Amount: 0, AssetID: ObjectId("1.3.0")}
	rpc.GetRequireFeeData(8, AssetData)
	st := wallet.CreateSignTransaction(8, Wallet.Default.GetActiveKey(), AssetData)
	rpc.BroadcastTransaction(st)
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
		Extensions:     []interface{}{},
		Issuer:         ObjectId(Wallet.Default.Info.ID),
		IssueToAccount: ObjectId(to_info.ID),
		AssetToIssue:   Amount{Amount: uint64(amount * precision), AssetID: ObjectId(asset_info.ID)},
	}
	issue.FeeData = Amount{Amount: 0, AssetID: ObjectId("1.3.0")}

	rpc.GetRequireFeeData(13, issue)
	st := wallet.CreateSignTransaction(13, Wallet.Default.GetActiveKey(), issue)
	rpc.BroadcastTransaction(st)
}
