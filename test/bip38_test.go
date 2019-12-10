package main

import (
	"CocosSDK/wallet"
	"testing"
)

func TestDecryptKey(t *testing.T) {
	wif, _ := wallet.DecryptKey(`6PRVFAxyDm8knMyBKsn4oL9KMsSC52gHoSFPBSUNa7vgxnAAMz1UN4kBeU`, []byte("07d5f5a9a01723a99b871beb455f2f1f2c34e12f39edfa1431c6ab4e18c60d9a"))

	if wif == `2h3RYjBFE4s8vCcu9wLJbynjGxWwoKiCNGfFyGEpNx6mz3nR97` {
		t.Log("pass")
	} else {
		t.Log(wif)
		t.Error(`error::`, wif)
	}
	t.Log(wallet.PrkFromWifString(wif).ToBase58String())

}

func TestEncryptKey(t *testing.T) {
	encryptWif, _ := wallet.EncryptKey("5KWL1EBGKgdrPvatwMPEPeiuwsnS9wBnvBLSV9Jz4Fojwdra1iA", []byte("07d5f5a9a01723a99b871beb455f2f1f2c34e12f39edfa1431c6ab4e18c60d9a"))
	if encryptWif == `6PRVFAxyDm8knMyBKsn4oL9KMsSC52gHoSFPBSUNa7vgxnAAMz1UN4kBeU` {
		t.Log("pass")
	} else {
		t.Error(`error::`, encryptWif)
	}
}
