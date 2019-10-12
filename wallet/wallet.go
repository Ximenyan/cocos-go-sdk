package wallet

import (
	"cocos-go-sdk/common"
	"cocos-go-sdk/rpc"
	. "cocos-go-sdk/type"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

type Wallet struct {
	path     string              `json:"-"`
	Accounts map[string]*Account `json:"Accounts"`
	Default  *Account            `json:"-"`
	lock     *sync.Mutex         `json:"-"`
}

//创建钱包
func CreateWallet() *Wallet {
	w := &Wallet{
		path:     "./wallet.dat",
		Accounts: make(map[string]*Account),
		lock:     &sync.Mutex{},
	}
	w.LoadWallet(w.path)
	return w
}

func (w *Wallet) initWallet() (err error) {
	w.Lock()
	defer w.Unlock()
	msh, err := ioutil.ReadFile(w.path)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(msh, w); err != nil {
		return err
	}
	return nil
}

//加载钱包
func (w *Wallet) LoadWallet(path string) (err error) {
	w.Lock()
	defer w.Unlock()
	if w.path != path {
		w.Save()
	}
	msh, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(msh, w); err != nil {
		return err
	}
	w.path = path
	return nil
}

//增加账户从私钥
func (w *Wallet) AddAccountByPrivateKey(prkWif string, password string) (err error) {
	w.Lock()
	defer w.Unlock()
	prk := PrkFromBase58String(prkWif)
	encryptWif, _ := EncryptKey(prkWif, []byte(password))
	puk := prk.GetPublicKey().ToBase58String()
	info := rpc.GetAccountInfoByPublicKey(puk)
	if info == nil {
		return errors.New("the key is not find!")
	}
	if _, succes := w.Accounts[info.Name]; !succes {
		w.Accounts[info.Name] = &Account{
			Name:     info.Name,
			KeyPairs: []KeyPair{},
			Info:     info,
		}
	}
	if !w.Accounts[info.Name].VerificationPassword(password) {
		return errors.New("the password is error!")
	}
	if active_puk, success := info.Active.KeyAuths[0][0].(string); success &&
		puk == active_puk &&
		w.Accounts[info.Name].GetActiveKey() == nil {
		activePair := KeyPair{
			Type:        "active",
			PubKey:      puk,
			EncryptWif:  encryptWif,
			Private_Key: prk,
		}
		memoPair := KeyPair{
			Type:        "memo",
			PubKey:      puk,
			EncryptWif:  encryptWif,
			Private_Key: prk,
		}
		w.Accounts[info.Name].KeyPairs = append(w.Accounts[info.Name].KeyPairs, []KeyPair{activePair, memoPair}...)
	} else if owner_puk, success := info.Owner.KeyAuths[0][0].(string); success &&
		puk == owner_puk &&
		w.Accounts[info.Name].GetOwnerKey() == nil {
		log.Println(info.Owner.KeyAuths[0][0].(string))
		ownerPair := KeyPair{
			Type:        "owner",
			PubKey:      puk,
			EncryptWif:  encryptWif,
			Private_Key: prk,
		}
		w.Accounts[info.Name].KeyPairs = append(w.Accounts[info.Name].KeyPairs, ownerPair)
	} /*else {
		unknowPair := KeyPair{
			Type:        "unknow",
			PubKey:      puk,
			EncryptWif:  encryptWif,
			Private_Key: prk,
		}
		w.Accounts[info.Name].KeyPairs = append(w.Accounts[info.Name].KeyPairs, unknowPair)
	}*/
	w.Save()
	return
}

//创建账户
func (w *Wallet) CreateAccount(name string, password string) (err error) {
	w.Lock()
	defer w.Unlock()
	if w.Default.Info == nil {
		w.Default.Info = rpc.GetAccountInfoByName(w.Default.Name)
	}
	w.Accounts[name] = CreateAccount(w.Default.GetActiveKey(), name, password, w.Default.Info.ID) //append(w.Accounts, CreateAccount(name, password))
	w.Save()
	return
}

//保存钱包到文件
func (this *Wallet) Save() error {
	data, err := json.Marshal(this)
	if err != nil {
		return err
	}
	if common.FileExisted(this.path) {
		filename := this.path
		err := ioutil.WriteFile(filename, data, 0644)
		if err != nil {
			return err
		}
		return os.Rename(filename, this.path)
	} else {
		return ioutil.WriteFile(this.path, data, 0644)
	}
}

//加锁
func (w *Wallet) Lock() (err error) {
	w.lock.Lock()
	return
}

//解锁
func (w *Wallet) Unlock() (err error) {
	w.lock.Unlock()
	return
}

//判断钱包账户是否为空
func (w *Wallet) IsEmpty() bool {
	if len(w.Accounts) <= 0 {
		return false
	}
	return true
}

//Transfer
func (w *Wallet) Transfer(to, symbol string, value uint64) error {
	t := CreateTransaction(w.Default.GetActiveKey(), w.Default.Name, to, symbol, value)
	rpc.GetRequireFeeData(0, t)
	st := CreateSignTransaction(0, w.Default.GetActiveKey(), t)
	return rpc.BroadcastTransaction(st)
}

//upgrade_account

func (w *Wallet) UpgradeAccount(name string) error {
	info := rpc.GetAccountInfoByName(name)
	t := CreateUpgradeAccount(name, info.ID)
	rpc.GetRequireFeeData(7, t)
	st := CreateSignTransaction(7, w.Default.GetActiveKey(), t)
	return rpc.BroadcastTransaction(st)
}

//SetDefaultAccount
func (w *Wallet) SetDefaultAccount(name, password string) error {
	w.Lock()
	defer w.Unlock()
	if acc, succes := w.Accounts[name]; succes {
		for i := 0; i < len(w.Accounts[name].KeyPairs); i++ {
			wif, _ := DecryptKey(w.Accounts[name].KeyPairs[i].EncryptWif, []byte(password))
			w.Accounts[name].KeyPairs[i].Private_Key = PrkFromWifString(wif)
			if w.Accounts[name].KeyPairs[i].PubKey != w.Accounts[name].KeyPairs[i].Private_Key.GetPublicKey().ToBase58String() {
				return errors.New("password error!")
			}
		}
		w.Default = acc
		return nil
	}
	return errors.New("no account name:" + name)
}
