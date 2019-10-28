package common

import (
	"testing"
)

func TestVarUint(t *testing.T) {

	t.Log(VarUint(10000000, 64))
}

func TestVarInt(t *testing.T) {
	t.Log(Varint2(1))
}

func TestVarUint16(t *testing.T) {
	t.Log(VarUint(9537, 16))
}
