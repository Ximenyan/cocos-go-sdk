package common

import (
	"testing"
)

func TestVarint(t *testing.T) {

	t.Log(Varint(1234431))
}

func TestIntVar(t *testing.T) {

	t.Log(Intvar([]byte{255, 171, 75}))
}
func TestVarUint(t *testing.T) {

	t.Log(VarUint(10000000, 64))
}
