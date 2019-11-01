package types

import (
	"time"
)

const COCOS_ID = `1.3.0`
const TIME_FORMAT = `2006-01-02T15:04:05`

var (
	DATABASE_API_ID  = 2
	BROADCAST_API_ID = 4
	HISTORY_API_ID   = 3
)

const (
	OP_TRANSFER         = 0
	OP_UPGRADE_ACCOUNT  = 7
	OP_CREATE_ACCOUNT   = 5
	OP_UPDATE_TOKEN     = 9
	OP_ISSUE_TOKEN      = 13
	OP_RESERVE_TOKEN    = 14
	OP_FUND_FEEPOOL     = 15
	OP_PROPOSAL         = 21
	OP_APPROVAL         = 22
	OP_VESTING_CREATE   = 31
	OP_VESTING_WITHDRAW = 32
	OP_CLAIM_FEES       = 39
	OP_CREATE_CONTRACT  = 34
	OP_INVOKE_CONTRACT  = 44
	OP_NH_CREATOR       = 46
	OP_DEL_NH_ASSET     = 50
	OP_SELL_NH_ASSET    = 52
	OP_CANCEL_NH_ORDER  = 53
	OP_FILL_NHORDER     = 54
	OP_REVISE_CONTRACT  = 59
)

var EMPTY_ID ObjectId = ""

func EmptyFee() Fee {
	A := Fee{FeeData: Amount{Amount: 0, AssetID: COCOS_ID}}
	return A
}

func GetExpiration() Expiration {
	return Expiration(time.Unix(time.Now().Unix(), 0).Format(TIME_FORMAT))
}
