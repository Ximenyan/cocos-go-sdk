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
[简单DEMO:CocosGoSdkDemo](https://github.com/Ximenyan/CocosGoSdkDemo)
## API接口
- [钱包相关API](#钱包相关API)
- [账户相关API](#账户相关API)
- [资产相关API](#资产相关API)
- [合约相关API](#合约相关API)
- [部分查询API](#部分查询API)
### 钱包相关API

#### 加载钱包
```
方法：sdk.Wallet.LoadWallet(file_path)
参数：file_path 钱包路径
```

#### 设置默认账户

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
#### 删除账户

```
方法：sdk.Wallet.DeleteAccountByName(name ...string) (err error)
参数：name   账户名
```
#### 创建账户

```
方法：sdk.Wallet.CreateAccount(name , password) error
参数：name   账户名
      password 密码
```
#### 保存钱包

```
方法：sdk.Wallet.SaveAs(path string) error
参数：path   保存路径
```

### 账户相关API

#### 升级终身账户

```
方法：sdk.Wallet.UpgradeAccount(name string) error
参数：name   账户名
```
#### 注册NH开发者

```
方法：sdk.Wallet.RegisterNhAssetCreator(name string) error
参数：name   账户名
```

### 资产相关API

#### 创建Token
```
方法：sdk.CreateAsset(symbol string, max_supply, precision, uint64) error
参数：
	symbol   token简写
	max_supplay 最大发行量
	precision   精度
	
```

#### Token发行
```
方法：sdk.IssueToken(symbol, issue_to_account string, amount float64) error
参数：
	symbol              token简写
	issue_to_account    接受Token的账户
	amount	            发行数量
```
#### Toke更新
```
方法：sdk.UpdateToken(symbol string, max_supply, precision uint64) error 
参数：
    new_issuer 新发行人
	symbol   token简写
	max_supplay 最大发行量
	precision   精度
	
```

#### Token销毁
```
方法：sdk.ReserveToken(symbol,  amount float64) error
参数：
	symbol              token简写
	amount	            销毁数量
```
#### 注资手续费池
```
方法：sdk.TokenFundFeePool(symbol string, amount float64) error
参数：
	symbol              token简写
	amount	            注资金额
```
#### 领取累计的手续费
```
方法：sdk.ClaimFees(symbol string, value float64) error
参数：
	symbol              token简写
	value	            提取数量
```
#### Token转账
```
方法：sdk.Wallet.Transfer(to, symbol, memo string, value float64) error
参数：
参数：
	symbol              token简写
	to                  接受Token的账户
	value	            发行数量
	memo		    备注
```
#### 创建Vesting Balance
```
方法：sdk.CreateVestingBalance(symbol string, amount float64) error
参数：
	symbol              token简写
	amount	            质押数量
```
#### 领取Vesting奖励
```
方法：sdk.WithdrawVestingBalance(balance_id string) error
参数：
	balance_id              Vesting Balance Id
```

#### 创建世界观
```
方法：sdk.CreateWorldView(name string) error 
参数：
	name              世界观名称
```
#### 提议关联世界观
```
方法：sdk.RelateWorldView(world_view string) error
参数：
	world_view              世界观名称
```
#### 批准关联世界观的提议
```
方法：sdk.ApprovalsProposal(proposal_id string) error
参数：
	proposal_id              提议的ID
```
#### 创建NH资产
```
方法：sdk.CreateNhAsset(asset_symbol, world_view, owner_name, base_describe string) error
参数：
参数：
	asset_symbol        当前NH资产交易时，使用的资产符号
	owner_name          接受账户
	world_view	    世界观
	base_describe       基础属性
```

#### NH资产转账
```
方法：sdk.TransferNhAsset(to_name, asset_id string) error
参数：
参数：
	to_name               接收账户
	asset_id              NH资产ID
```

#### NH资产删除
```
方法：sdk.DeleteNhAsset(asset_id string) error
参数：
参数：
	asset_id              NH资产ID
```

#### 创建卖出NH资产订单

```
方法：sdk.SellNhAsset(otcaccount_name, asset_id, memo, pending_order_fee_asset, price_asset string, pending_order_fee_amount, price_amount uint64) error 
参数：
参数：
	otcaccount_name：OTC交易平台账户，用于收取挂单费用
        pending_order_fee_amount：挂单费用数量，用户向OTC平台账户支付的挂单费用
	pending_order_fee_asset：挂单费用所用资产符号或ID，用户向OTC平台账户支付的挂单费用
	asset_id：NH资产ID
	memo：挂单备注信息
	price_amount：商品挂单价格数量
	price_asset：商品挂单价格所用资产符号或ID
```

#### 撤销NH资产卖出单

```
方法：sdk.CancelNhAssetOrder(order_id string) error 
参数：
参数：
	order_id  ：订单ID
```

#### 吃单，买入NH资产

```
方法：FillNhAsset(order_id string) error
参数：
参数：
	order_id  ：订单ID
```

### 合约相关API


#### 创建合约1

```
方法：sdk.CreateContractByFile(c_name, c_auth, path string) error
参数：
参数：
	c_name  ：合约名
	c_auth  ：合约权限(一对公私钥中的公钥publicKey)
	path    ：合约（lua 代码）在本地的存放路径
```
#### 创建合约2

```
方法：sdk.CreateContract(c_name, c_auth, data string) error
参数：
参数：
	c_name  ：合约名
	c_auth  ：合约权限(一对公私钥中的公钥publicKey)
	data    ：合约（lua 代码）
```

#### 更新合约1

```
方法：sdk.ReviseContractByFile(c_name, path string) error
参数：
参数：
	c_name  ：合约名
	path    ：合约（lua 代码）在本地的存放路径
```
#### 更新合约2

```
方法：sdk.ReviseContract(c_name, data string) error
参数：
参数：
	c_name  ：合约名
	data    ：合约（lua 代码）
```
#### 合约调用

```
方法：sdk.InvokeContract(contract_name, func_name string, args ...interface{})
参数：
参数：
	contract_name ： 合约名
	func_name ： 调用方法名
	args  ：参数
```
### 部分查询API

#### 查询订单信息

```
方法：GetNhAssetOrderInfo(id string) NhAssetOrderInfo
参数：
参数：
	id ： 订单ID
	func_name ： 调用方法名
	args  ：参数
```


#### owner订单列表

```
方法： GetAccountNhAssetOrderList(owner_name string, page, page_size int) OrdersList
```
#### 查询订单列表

```
方法： GetNhAssetOrderList(asset_name, world_view string, page, page_size int) OrdersList
```
#### owner订单列表

```
方法： GetAccountNhAssetOrderList(owner_name string, page, page_size int) OrdersList
```
#### owner资产列表

```
方法： GetNhAssetList(acc_name string, page, page_size, _type int, world_view ...string) AssetsList
```

#### 查询区块

```
方法：  GetBlock(block_hight int) Block
```
#### 查询交易信息

```
方法： GetTransactionById(txId string) TransactinInf
```
#### 查询账户Balance

```
方法： GetAccountBalances(acc_name string) *[]rpc.Balance
```
#### 查询链上所有token信息

```
方法： sdk.GetAllTokenInfo() []*rpc.TokenInfo
```


#### 查询收到的所有提议

```
方法： sdk.GetAllProposals(acct_id string) *[]rpc.Proposal
```

#### 查询 某条提议
```
方法： sdk.GetAllProposal(proposal_id string) *[]rpc.Proposal
```

#### 通过Symbol查询token信息
```
方法： sdk.GetTokenInfoBySymbol(symbol string) *rpc.TokenInfo 
```

#### 通过id查询token信息
```
方法： sdk.GetTokenInfoById(id string) *rpc.TokenInfo 
```

#### 查询BlockHeader
```
方法： sdk.GetBlockHeader(block_hight int) *rpc.BlockHeader
```

#### 查询Contract
```
方法： sdk.GetContract(contract_name string) *rpc.Contract
```
#### 查询账户待提取的奖励
```
方法： sdk.GetVestingBalances(acct_name string) []rpc.VestingBalances
```
#### 查询账户操作记录
```
方法： sdk.GetAccountHistorys(acct_name string) []interface{} 
```

#### 获取市场限价单交易历史
```
方法： sdk.GetFillOrderHistory(asset_id, _asset_id string, limit uint64) []interface{}
```

#### 查询某个时间段的交易市场行情
```
方法： sdk.GetMarketHistory(asset_id, _asset_id, start, end string, limit uint64) []interface{}
```

