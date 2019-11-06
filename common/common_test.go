package common

import (
	"testing"
)

func TestVarint(t *testing.T) {

	t.Log(Varint(0x81))
}
func TestVarUint(t *testing.T) {

	t.Log(VarUint(10000000, 64))
}
