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
	OP_UPDATE_TOKEN     = 9
	OP_ISSUE_TOKEN      = 13
	OP_RESERVE_TOKEN    = 14
	OP_FUND_FEEPOOL     = 15
	OP_PROPOSAL         = 21
	OP_APPROVAL         = 22
	OP_VESTING_CREATE   = 31
	OP_VESTING_WITHDRAW = 32
	OP_CLAIM_FEES       = 39
	OP_CREATE_CONTRACT  = 43
	OP_INVOKE_CONTRACT  = 44
	OP_REVISE_CONTRACT  = 59
)

func EmptyFee() Fee {
	A := Fee{FeeData: Amount{Amount: 0, AssetID: COCOS_ID}}
	return A
}

func GetExpiration() Expiration {
	return Expiration(time.Unix(time.Now().Unix(), 0).Format(TIME_FORMAT))
}
