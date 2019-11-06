package common

import (
	"math/big"
	"os"
)

// 反转Bytearray
func ReverseBytes(res []byte) []byte {
	for from, to := 0, len(res)-1; from < to; from, to = from+1, to-1 {
		res[from], res[to] = res[to], res[from]
	}
	return res
}

// FileExisted checks whether filename exists in filesystem
func FileExisted(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func Intvar(byte_s []byte) int64 {
	if byte_s[0] == 0 {
		return 0
	}
	//i0x80 := new(big.Int).SetUint64(0x80)
	i0x7f := new(big.Int).SetUint64(0x7f)
	tmp := len(byte_s) - 1
	new_i := new(big.Int)
	i := new(big.Int).SetBytes(byte_s[tmp:])
	i = new(big.Int).And(i, i0x7f)
	for tmp > 0 {
		new_i = new(big.Int).Or(new_i, i)
		new_i = new_i.Lsh(new_i, 7)
		tmp -= 1
		i = new(big.Int).SetBytes(byte_s[tmp : tmp+1])
		i = new(big.Int).And(i, i0x7f)
	}
	new_i = new(big.Int).Or(new_i, i)
	return new_i.Int64()
}

func Varint(ui uint64) []byte {
	if ui == 0 {
		return []byte{0}
	}
	i := new(big.Int).SetUint64(ui)
	//log.Println(i.Bytes())
	i0x80 := new(big.Int).SetUint64(0x80)
	i0x7f := new(big.Int).SetUint64(0x7f)
	byte_s := []byte{}
	for i.Cmp(i0x80) == 1 {
		b := new(big.Int).And(i, i0x7f)
		//log.Println(b.)
		b = new(big.Int).Or(b, i0x80)
		byte_s = append(byte_s, b.Bytes()...)
		//log.Println(byte_s)
		i = i.Rsh(i, 7)
	}
	byte_s = append(byte_s, i.Bytes()...)
	return byte_s
}

func VarInt(si int64, base uint) []byte {
	if si == 0 {
		var byte_s [64]byte
		return byte_s[:base/8]
	}
	i := new(big.Int).SetInt64(si)
	l := uint(i.BitLen())
	byte_s := ReverseBytes(i.Bytes())
	ln := len(byte_s)
	if base > l {
		i = new(big.Int).Lsh(i, uint(((base-l)/8)*8))
	}
	byte_s = append(byte_s, i.Bytes()[ln:]...)
	if si < 0 {
		byte_s[len(byte_s)-1] = 0x80
	}
	return byte_s
}

func UintVar(byte_s []byte) uint64 {
	i := len(byte_s)
	for byte_s[i-1] == 0 {
		byte_s = byte_s[:i]
		i -= 1
	}
	return new(big.Int).SetBytes(byte_s).Uint64()
}

func VarUint(ui uint64, base uint) []byte {
	if ui == 0 {
		var byte_s [64]byte
		return byte_s[:base/8]
	}
	i := new(big.Int).SetUint64(ui)
	l := uint(i.BitLen())
	byte_s := ReverseBytes(i.Bytes())
	ln := len(byte_s)

	if base > l {
		i = new(big.Int).Lsh(i, uint(((base-l)/8)*8))
	}
	byte_s = append(byte_s, i.Bytes()[ln:]...)
	return byte_s
}
