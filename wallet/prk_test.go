package wallet

import (
	"CocosSDK/crypto/secp256k1"
	//	"CocosSDK/rpc"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"strconv"
	"testing"
)

//const TEST_NET = "47.93.62.96"
//const LOCAL = "192.168.0.166"

//var _ = rpc.InitClient(TEST_NET, 8049, false)

func TestMemo(t *testing.T) {
	puk := PukFromBase58String("COCOS7X5HPYLUicec1HLK9LxWWyZhAEx2RxgNWKPuK3MZ9GQDfrRUe3")
	t.Log(hex.EncodeToString(puk))
}

func TestPrk(t *testing.T) {
	//elliptic.Curve
	//elliptic.TestOnCurve()
	message := "ximenyan1111"
	prk := PrkFromBase58String("5HqGVLJ2zN5tJw7JNKsc8e25Gadn5GYiwkEde3XrhUN3D1pZx2P")

	puk := PukFromBase58String("COCOS6qF5SzyYRcnEiHjF3zF143LSyyvGRcs1tYMc5iBTvWJGHhb39V")
	puk.UnCompressed()
	x, y := puk.GetPoint()
	cure := secp256k1.S256()
	x, y = cure.ScalarMult(x, y, prk.PrivKey)

	sha := sha512.New()
	byte_s := x.Bytes()
	sha.Write(byte_s)
	resss := sha.Sum(nil)

	noce := strconv.FormatUint(GetNonce(), 10)
	seed := noce + hex.EncodeToString(resss)
	sha.Reset()

	sha.Write([]byte(seed))
	seed_digest := sha.Sum(nil)
	s256 := sha256.New()
	s256.Write([]byte(message))
	checksum := s256.Sum(nil)
	byte_s_msg := append(checksum[0:4], []byte(message)...)
	num := 16 - len(byte_s_msg)%16
	for i := 0; i < num; i++ {
		byte_s_msg = append(byte_s_msg, byte(num))
	}
	block, _ := aes.NewCipher(seed_digest[0:32])
	m := cipher.NewCBCEncrypter(block, seed_digest[32:48])
	m.CryptBlocks(byte_s_msg, byte_s_msg)
	t.Log(hex.EncodeToString(byte_s_msg))
}
