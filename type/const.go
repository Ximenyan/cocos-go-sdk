package types

import (
	"time"
)

const COCOS_ID = `1.3.0`
const TIME_FORMAT = `2006-01-02T15:04:05`
var UTCZone = time.FixedZone("UTC", 0)
var (
	DATABASE_API_ID  = 2
	BROADCAST_API_ID = 4
	HISTORY_API_ID   = 3
)

const (
	OP_TRANSFER           = 0
	OP_CREATE_ACCOUNT     = 5
	OP_VOTE               = 6
	OP_UPGRADE_ACCOUNT    = 7
	OP_CREATE_ASSET_TOKEN = 8
	OP_UPDATE_TOKEN       = 9
	OP_ISSUE_TOKEN        = 13
	OP_RESERVE_TOKEN      = 14
	OP_FUND_FEEPOOL       = 15
	OP_PROPOSAL           = 20
	OP_APPROVAL           = 21
	//OP_VESTING_CREATE     = 27
	OP_VESTING_WITHDRAW  = 27
	OP_CLAIM_FEES        = 31
	OP_CREATE_CONTRACT   = 34
	OP_INVOKE_CONTRACT   = 35
	OP_NH_CREATOR        = 37
	OP_CREATE_WORLDVIEW  = 38
	OP_RELATE_WORLDVIEW  = 39
	OP_CREATE_NH_ASSET   = 40
	OP_DEL_NH_ASSET      = 41
	OP_TRANSFER_NH_ASSET = 42
	OP_SELL_NH_ASSET     = 43
	OP_CANCEL_NH_ORDER   = 44
	OP_FILL_NHORDER      = 45
	OP_REVISE_CONTRACT   = 50
	OP_PLEDGE_GAS        = 54
)

var EMPTY_ID ObjectId = ""

func EmptyFee() Fee {
	A := Fee{FeeData: Amount{Amount: 0, AssetID: COCOS_ID}}
	return A
}

func ToExpiration(expiration string) Expiration {
	t, _ := time.Parse(TIME_FORMAT, string(expiration))
	return Expiration( time.Unix(t.Unix()+30,0).In(UTCZone).Format(TIME_FORMAT))
}

func GetExpiration() Expiration {
	return Expiration(time.Now().In(UTCZone).Format(TIME_FORMAT))
}