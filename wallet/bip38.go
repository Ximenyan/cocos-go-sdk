package wallet

import (
	"bytes"
	"crypto/aes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"math/big"

	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/scrypt"
)

//base58编码
var b58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

func Base58CheckEncode(version, payload []byte) []byte {
	s := append(version, payload...)
	checksum := checkSum(s)
	result := append(s, checksum...)
	return Base58Encode(result)
}

func Base58Encode(input []byte) []byte {
	var result []byte

	x := big.NewInt(0).SetBytes(input)

	base := big.NewInt(int64(len(b58Alphabet)))
	zero := big.NewInt(0)

	mod := &big.Int{}
	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod) // 对x取余数
		result = append(result, b58Alphabet[mod.Int64()])
	}

	ReverseBytes(result)

	for _, b := range input {

		if b == 0x00 {
			result = append([]byte{b58Alphabet[0]}, result...)
		} else {
			break
		}
	}
	return result

}

//字节数组的反转
func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}
func Base58Decode(input []byte) []byte {
	result := big.NewInt(0)
	zeroBytes := 0
	for _, b := range input {
		if b == '1' {
			zeroBytes++
		} else {
			break
		}
	}

	payload := input[zeroBytes:]

	for _, b := range payload {
		charIndex := bytes.IndexByte(b58Alphabet, b) //反推出余数

		result.Mul(result, big.NewInt(58)) //之前的结果乘以58

		result.Add(result, big.NewInt(int64(charIndex))) //加上这个余数

	}

	decoded := result.Bytes()

	decoded = append(bytes.Repeat([]byte{0x00}, zeroBytes), decoded...)
	return decoded
}

// BIP38 encrypt implement
func EncryptKey(privateWif string, passKey []byte) (encryptWif string, err error) {
	// TODO: test compressed private WIF
	prk := PrkFromBase58String(privateWif)
	puk := prk.GetUnCompressedPubkey()

	addr := puk.GetSha256Address()
	new_addr := Base58CheckEncode([]byte{0}, addr)
	sum := checkSum(new_addr)
	key, err := scrypt.Key(passKey, sum, 16384, 8, 8, 64)
	if err != nil {
		return
	}
	data := encrypt(prk.PrivKey, key[:32], key[32:])
	buf := append([]byte{0x01, 0x42, 0xC0}, sum...)
	buf = append(buf, data...)
	buf = append(buf, checkSum(buf)...)
	encryptWif = string(Base58Encode(buf))
	return
}

func DecryptKey(encryptWif string, passKey []byte) (privateWif string, err error) {
	byte_s := Base58Decode([]byte(encryptWif))
	byte_s = byte_s[2:]
	flag_byte := byte_s[0:1]
	byte_s = byte_s[1:]
	if flag_byte[0] != 0xc0 {
		return ``, errors.New("privateWif error !!!!")
	}
	sum := byte_s[0:4]
	byte_s = byte_s[4 : len(byte_s)-4]
	key, err := scrypt.Key(passKey, sum, 16384, 8, 8, 64)
	derf1 := key[0:32]
	derf2 := key[32:64]
	enf1 := byte_s[0:16]
	enf2 := byte_s[16:32]

	c, err := aes.NewCipher(derf2)
	var decf2 = make([]byte, 16)
	var decf1 = make([]byte, 16)
	c.Decrypt(decf2, enf2)
	c.Decrypt(decf1, enf1)
	privraw := append(decf1, decf2...)
	b_int_1, _ := big.NewInt(0).SetString(hex.EncodeToString(privraw), 16)
	b_int_2, _ := big.NewInt(0).SetString(hex.EncodeToString(derf1), 16)
	b_int_raw := b_int_1.Xor(b_int_1, b_int_2)
	rip := ripemd160.New()
	ripemd160.New()
	rip.Write(b_int_raw.Bytes())
	new_sum := rip.Sum([]byte{})[:4]
	privateWif = string(Base58Encode(append(b_int_raw.Bytes(), new_sum...)))
	return
}

func checkSum(buf []byte) []byte {
	h := sha256.Sum256(buf)
	h = sha256.Sum256(h[:])
	return h[:4]
}

func encrypt(pk, f1, f2 []byte) []byte {
	c, _ := aes.NewCipher(f2)

	for i, _ := range f1 {
		f1[i] ^= pk[i]
	}

	var dst = make([]byte, 48)
	c.Encrypt(dst, f1[:16])
	c.Encrypt(dst[16:], f1[16:])
	dst = dst[:32]

	return dst
}
