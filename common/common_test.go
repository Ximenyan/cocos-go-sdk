package common

import (
	"testing"
)

func TestVarint(t *testing.T) {

	t.Log(Varint(1234431))
}

func TestVarUint(t *testing.T) {

	t.Log(VarUint(1573395781, 32))
	t.Log(VarUint(1573424581, 32))
}
func TestUintVar(t *testing.T) {

	t.Log(UintVar([]byte{197, 141, 200, 93}))
}
