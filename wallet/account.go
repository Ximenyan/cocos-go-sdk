package wallet

import (
	"cocos-go-sdk/rpc"
)

type KeyPair struct {
	Type        string     `json:"type"`
	PubKey      string     `json:"pubkey"`
	EncryptWif  string     `json:"encrypt_wif"`
	Private_Key PrivateKey `json:"-"`
}

type Account struct {
	Name     string           `json:"name"`
	KeyPairs []KeyPair        `json:"key_pair"`
	Info     *rpc.AccountInfo `json:"-"`
}

func (acc Account) GetActiveKey() *PrivateKey {
	for _, pair := range acc.KeyPairs {
		if pair.Type == `active` {
			return &pair.Private_Key
		}
	}
	return nil
}

func (acc Account) GetOwnerKey() *PrivateKey {
	for _, pair := range acc.KeyPairs {
		if pair.Type == `owner` {
			return &pair.Private_Key
		}
	}
	return nil
}
func (acc Account) GetMemoKey() *PrivateKey {
	for _, pair := range acc.KeyPairs {
		if pair.Type == `memo` {
			return &pair.Private_Key
		}
	}
	return nil
}
func (acc Account) VerificationPassword(password string) bool {
	for i := 0; i < len(acc.KeyPairs); i++ {
		wif, _ := DecryptKey(acc.KeyPairs[i].EncryptWif, []byte(password))
		acc.KeyPairs[i].Private_Key = PrkFromWifString(wif)
		if acc.KeyPairs[i].PubKey != acc.KeyPairs[i].Private_Key.GetPublicKey().ToBase58String() {
			return false
		}
	}
	return true
}
func CreateAccount(prk *PrivateKey, name string, password string, registrar string) *Account {
	active_PrivateKey := CreatePrivateKeyFromSeed(name + "active" + password)
	active_PubKey := active_PrivateKey.GetPublicKey().ToBase58String()
	active_EncryptWif, _ := EncryptKey(active_PrivateKey.ToBase58String(), []byte(password))
	owner_PrivateKey := CreatePrivateKeyFromSeed(name + "owner" + password)
	owner_PubKey := owner_PrivateKey.GetPublicKey().ToBase58String()
	owner_EncryptWif, _ := EncryptKey(owner_PrivateKey.ToBase58String(), []byte(password))
	activePair := KeyPair{
		Type:        "active",
		PubKey:      active_PubKey,
		EncryptWif:  active_EncryptWif,
		Private_Key: active_PrivateKey,
	}
	memoPair := KeyPair{
		Type:        "memo",
		PubKey:      active_PubKey,
		EncryptWif:  active_EncryptWif,
		Private_Key: active_PrivateKey,
	}
	ownerPair := KeyPair{
		Type:        "owner",
		PubKey:      owner_PubKey,
		EncryptWif:  owner_EncryptWif,
		Private_Key: owner_PrivateKey,
	}
	keys := make(map[string]KeyPair)
	keys["active"] = activePair
	keys["owner"] = ownerPair
	keys["memo"] = activePair
	acc := &Account{
		Name:     name,
		KeyPairs: []KeyPair{activePair, memoPair, ownerPair},
	}
	u := CreateRegisterData(active_PubKey, owner_PubKey, name, registrar, registrar)
	st := CreateSignTransaction(5, prk, u)
	rpc.BroadcastTransaction(st)
	return acc
}
