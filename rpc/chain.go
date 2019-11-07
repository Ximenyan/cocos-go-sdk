package rpc

import (
	"cocos-go-sdk/common"
	. "cocos-go-sdk/type"
	"encoding/hex"
	"encoding/json"
	"log"
	"math/big"
)

const CALL = `call`
const EMPTY = ""

func GetRequireFeeData(opID int, t OpData) (interface{}, error) {
	fees := &[]interface{}{}
	req := CreateRpcRequest(CALL,
		[]interface{}{0, `get_required_fees`,
			[]interface{}{[]interface{}{[]interface{}{opID, t}}, "1.3.0"}})
	if resp, err := Client.Send(req); err == nil {
		var iter interface{}
		var byte_s []byte
		iter = resp.Result
		for iter != nil {
			amount := &Amount{}
			if byte_s, err = json.Marshal(iter); err == nil {
				if err = json.Unmarshal(byte_s, amount); err == nil {
					t.SetFee(amount.Amount)
					return amount, nil
				}
				if err = json.Unmarshal(byte_s, fees); err == nil {
					iter = (*fees)[0]
				}
			}
		}
		return resp.Result, err
	} else {
		return nil, err
	}
}

type TransactinInfo struct {
	RefBlockNum      int             `json:"ref_block_num"`
	RefBlockPrefix   int64           `json:"ref_block_prefix"`
	Expiration       string          `json:"expiration"`
	Operations       [][]interface{} `json:"operations"`
	Extensions       []interface{}   `json:"extensions"`
	Signatures       []string        `json:"signatures"`
	OperationResults [][]interface{} `json:"operation_results"`
}

func GetTransactionById(txId string) *TransactinInfo {
	req := CreateRpcRequest(CALL,
		[]interface{}{0, `get_transaction_by_id`,
			[]interface{}{txId}})
	if resp, err := Client.Send(req); err == nil {
		tx_info := &TransactinInfo{}
		if err = resp.GetInterface(tx_info); err == nil {
			return tx_info
		}
	}
	return nil
}

type Block struct {
	Previous              string          `json:"previous"`
	Timestamp             string          `json:"timestamp"`
	Witness               string          `json:"witness"`
	TransactionMerkleRoot string          `json:"transaction_merkle_root"`
	Extensions            []interface{}   `json:"extensions"`
	WitnessSignature      string          `json:"witness_signature"`
	BlockID               string          `json:"block_id"`
	Transactions          [][]interface{} `json:"transactions"`
}

func GetBlock(block int) *Block {
	req := CreateRpcRequest(CALL,
		[]interface{}{0, `get_block`,
			[]interface{}{block}})
	block_info := &Block{}
	if resp, err := Client.Send(req); err == nil {
		//log.Println(resp)
		if err = resp.GetInterface(block_info); err == nil {
			return block_info
		}
	}
	return nil
}

func GetBlocks(blocks []int) *[]Block {
	req := CreateRpcRequest(CALL,
		[]interface{}{0, `get_block`,
			blocks})
	blocks_info := &[]Block{}
	if resp, err := Client.Send(req); err == nil {
		if err = resp.GetInterface(blocks_info); err == nil {
			return blocks_info
		}
	}
	return nil
}

type BlockHeader struct {
	Extensions            []interface{} `json:"extensions"`
	Previous              string        `json:"previous"`
	Timestamp             string        `json:"timestamp"`
	TransactionMerkleRoot string        `json:"transaction_merkle_root"`
	Witness               string        `json:"witness"`
}

func GetBlockHeader(block int) *BlockHeader {
	req := CreateRpcRequest(CALL,
		[]interface{}{DATABASE_API_ID, `get_block_header`,
			[]interface{}{block}})
	header := &BlockHeader{}
	if resp, err := Client.Send(req); err == nil {
		if err = resp.GetInterface(header); err == nil {
			return header
		}
	}
	return nil
}

type VestingBalances struct {
	Balance struct {
		Amount  interface{} `json:"amount"`
		AssetID string      `json:"asset_id"`
	} `json:"balance"`
	ID     string        `json:"id"`
	Owner  string        `json:"owner"`
	Policy []interface{} `json:"policy"`
}

func (v VestingBalances) GetBalanceAmount() uint64 {
	if str, b := v.Balance.Amount.(string); b {
		byte_s, _ := hex.DecodeString(str)
		byte_s = common.ReverseBytes(byte_s)
		return new(big.Int).SetBytes(byte_s).Uint64()
	} else {
		return uint64(v.Balance.Amount.(float64))
	}
	log.Panicln("VestingBalances  GetBalanceAmount Error!!!")
	return 0
}
func GetVestingBalancesByName(acct_name string) []VestingBalances {
	acct_info := GetAccountInfoByName(acct_name)
	req := CreateRpcRequest(CALL,
		[]interface{}{DATABASE_API_ID, `get_vesting_balances`,
			[]interface{}{acct_info.ID}})
	balances := &[]VestingBalances{}
	if resp, err := Client.Send(req); err == nil {
		if err = resp.GetInterface(balances); err == nil {
			return *balances
		}
	}
	return nil
}

type DynamicGlobalProperties struct {
	AccountsRegistered_thisInterval int    `json:"accounts_registered_this_interval"`
	CurrentAslot                    int    `json:"current_aslot"`
	CurrentTransactionCount         int    `json:"current_transaction_count"`
	CurrentWitness                  string `json:"current_witness"`
	DynamicFlags                    int    `json:"dynamic_flags"`
	HeadBlockID                     string `json:"head_block_id"`
	HeadBlockNumber                 int    `json:"head_block_number"`
	ID                              string `json:"id"`
	LastBudgetTime                  string `json:"last_budget_time"`
	LastIrreversib_leBlockNum       int    `json:"last_irreversible_block_num"`
	NextMaintenanceTime             string `json:"next_maintenance_time"`
	RecentSlotsFilled               string `json:"recent_slots_filled"`
	RecentlyMissedCount             int    `json:"recently_missed_count"`
	Time                            string `json:"time"`
	WitnessBudget                   int    `json:"witness_budget"`
}

func GetDynamicGlobalProperties() *DynamicGlobalProperties {
	dgp := &DynamicGlobalProperties{}
	req := CreateRpcRequest(CALL,
		[]interface{}{0, `get_dynamic_global_properties`,
			[]interface{}{}})
	if resp, err := Client.Send(req); err == nil {
		if err = resp.GetInterface(dgp); err == nil {
			return dgp
		}
	}
	return nil
}

func (p *DynamicGlobalProperties) Get_ref_block_num() uint64 {
	return uint64(p.HeadBlockNumber & 0xffff)
}

func (p *DynamicGlobalProperties) Get_ref_block_prefix() uint64 {
	byte_s, _ := hex.DecodeString(p.HeadBlockID)
	ref_block_prefix := new(big.Int).SetBytes(common.ReverseBytes(byte_s[4:8])).Uint64()
	return ref_block_prefix
}
