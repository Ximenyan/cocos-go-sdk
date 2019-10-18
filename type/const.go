package types

import (
	"time"
)

const COCOS_ID = `1.3.0`
const TIME_FORMAT = `2006-01-02T15:04:05`
const (
	DATABASE_API_ID = 2
)
const (
	OP_CREATE_CONTRACT = 43
	OP_PROPOSAL        = 21
	OP_APPROVAL        = 22
	OP_CLAIM_FEES      = 39
)

func EmptyFee() Fee {
	A := Fee{FeeData: Amount{Amount: 0, AssetID: COCOS_ID}}
	return A
}

func GetExpiration() Expiration {
	return Expiration(time.Unix(time.Now().Unix(), 0).Format(TIME_FORMAT))
}
