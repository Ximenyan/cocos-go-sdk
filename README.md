# Golang SDK For Cocos-BCX
## 安装

```
git clone https://github.com/Ximenyan/cocos-go-sdk.git
```
或者

```
go get github.com/Ximenyan/cocos-go-sdk
```

## 使用

```
import (
	sdk "cocos-go-sdk"
	"fmt"
)

func main() {
    //初始化SDK 
    //节点host port 是否ssl
	sdk.InitSDK("47.93.62.96", 8049, false)
}

```

## API接口
- [钱包相关API](###钱包相关API)
- [账户相关API](###钱包相关API)
- [资产相关API](###钱包相关API)
- [合约相关API](###钱包相关API)

### 钱包相关API

#### 加载钱包
```
方法：sdk.Wallet.LoadWallet(file_path)
参数：file_path 钱包路径
```

#### 设置账户

```
方法：sdk.Wallet.SetDefaultAccount(name, password string) error
参数：name     账户名
      password 账户密码
```
#### 导入私钥

```
方法：sdk.Wallet.AddAccountByPrivateKey(prkWif， password ) error
参数：prkWif   私钥
      password 密码
 
```
#### 创建账户

```
方法：sdk.Wallet.CreateAccount(name , password) error
参数：name   账户名
      password 密码
```


### 账户相关API

### 资产相关API

### 合约相关API