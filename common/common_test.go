package common

import (
	"testing"
)

func TestVarint(t *testing.T) {

	t.Log(Varint(1234431))
}

<<<<<<< HEAD
func TestIntVar(t *testing.T) {

	t.Log(Intvar([]byte{255, 171, 75}))
}
func TestVarUint(t *testing.T) {

	t.Log(VarUint(10000000, 64))
=======
func TestVarUint(t *testing.T) {

	t.Log(VarUint(1573395781, 32))
	t.Log(VarUint(1573424581, 32))
}
func TestUintVar(t *testing.T) {

	t.Log(UintVar([]byte{197, 141, 200, 93}))
>>>>>>> master
}
