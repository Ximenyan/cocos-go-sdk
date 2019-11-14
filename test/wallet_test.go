package main

import (
	"CocosSDK/rpc"
	"CocosSDK/wallet"
	"testing"
)

func TestBase58toPrk(t *testing.T) {
	prk := wallet.PrkFromBase58String("5JSs8GhR5e1GCPG6B3PRxeGYUajzEdF4FtbEDbQHbQqv5bwm5yk")
	t.Log(prk.ToBase58String())
	puk := prk.GetPublicKey()
	if puk.ToBase58String() != "COCOS6FZYBGR1DuPsgYpUSN3fgQtbASRuomLrjT5SqzEjCqJVKEpJuw" {
		t.Fatal("TestBase58toPrk error......")
	}
}

func TestAccount(t *testing.T) {
	rpc.InitClient("47.93.62.96", 8049, false)
	w := wallet.CreateWallet("asd123")
	w.CreateAccount("ximenyan", []byte("asd123"))
	w.AddAccountByPrivateKey("5Jfi4ESi9MoN7jGMR1bRB9NRUwJ1duPqiGpmzxhdCQJCZQAaLvq", []byte("sddd1234"))
	w.Save("./wallet.dat")
}
