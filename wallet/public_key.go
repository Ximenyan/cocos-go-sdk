package wallet

import (
	"CocosSDK/crypto/base58-go"
	"CocosSDK/crypto/secp256k1"
	"crypto/sha256"
	"math/big"

	"golang.org/x/crypto/ripemd160"
)

type PublicKey []byte

const Prefix = "COCOS"

func (puk PublicKey) GetSha256Address() []byte {
	sum := sha256.Sum256(puk)
	r := ripemd160.New()
	r.Write(sum[:])
	sum_rip := r.Sum([]byte{})
	return sum_rip
}

func (puk PublicKey) UnCompressed() PublicKey {
	x, y := puk.GetPoint()
	pubkey := append([]byte{4},
		append(x.Bytes(),
			y.Bytes()...)...)
	return pubkey
}

func (puk PublicKey) GetPoint() (x, y *big.Int) {
	xy := secp256k1.GetXY(puk)
	x, y = xy.X.GetBig(), xy.Y.GetBig()
	return
}
func (puk PublicKey) ToBase58String() string {
	rip := ripemd160.New()
	rip.Write(puk)
	temps_2 := rip.Sum([]byte{})
	data := append(puk, temps_2[0:4]...)
	bi := new(big.Int).SetBytes(data).String()
	encoded, _ := base58.BitcoinEncoding.Encode([]byte(bi))
	return "COCOS" + string(encoded)
}

func PukFromBase58String(base58Str string) PublicKey {
	byte_s, _ := base58.BitcoinEncoding.Decode([]byte(base58Str)[5:])
	big_i, _ := new(big.Int).SetString(string(byte_s), 10)
	data := big_i.Bytes()
	puk := data[0 : len(data)-4]
	return puk
}
