package CocosSDK

import (
	"cocos-go-sdk/rpc"
)

/*查询Block*/
func GetBlock(block_hight int64) *rpc.Block {
	return rpc.GetBlock(block_hight)
}

/*查询Blocks*/
func GetBlocks(block_hights []int) *[]rpc.Block {
	return rpc.GetBlocks(block_hights)
}

/*查询BlockHeader*/
func GetBlockHeader(block_hight int) *rpc.BlockHeader {
	return rpc.GetBlockHeader(block_hight)
}

/*查询交易*/
func GetTransactionById(txId string) *rpc.TransactinInfo {
	return rpc.GetTransactionById(txId)
}
