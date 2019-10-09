package chain

import (
	"cocos-go-sdk/common"
	"encoding/hex"
	"math/big"
	"testing"
)

func Test1(t *testing.T) {
	byte_s, _ := hex.DecodeString(`004de2a4d5d6a22518c9c2a174b08d778952667c`)
	t.Log(new(big.Int).SetBytes(common.ReverseBytes(byte_s[4:8])).Uint64())
	//3453533023
}
