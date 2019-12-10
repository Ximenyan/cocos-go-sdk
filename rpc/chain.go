package rpc

import (
	"CocosSDK/common"
	. "CocosSDK/type"
	"encoding/hex"
	"encoding/json"
	"log"
	"math/big"
	"time"

	"github.com/tidwall/gjson"
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
type TXInBlockInfo struct {
	ID         string `json:"id"`
	BlockNum   int64  `json:"block_num"`
	TrxInBlock int    `json:"trx_in_block"`
	TrxHash    string `json:"trx_hash"`
}

func GetTransactionInBlock(txId string) *TXInBlockInfo {
	req := CreateRpcRequest(CALL,
		[]interface{}{0, `get_transaction_in_block_info`,
			[]interface{}{txId}})
	if resp, err := Client.Send(req); err == nil {
		tx_info := &TXInBlockInfo{}
		if err = resp.GetInterface(tx_info); err == nil {
			return tx_info
		}
	}
	return nil
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

func GetBlock(block int64) *Block {
	req := CreateRpcRequest(CALL,
		[]interface{}{0, `get_block`,
			[]interface{}{block}})
	block_info := &Block{}
	if resp, err := Client.Send(req); err == nil {
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
	var amount uint64
	if str, b := v.Balance.Amount.(string); b {
		byte_s, _ := hex.DecodeString(str)
		byte_s = common.ReverseBytes(byte_s)
		amount = new(big.Int).SetBytes(byte_s).Uint64()
	} else {
		amount = uint64(v.Balance.Amount.(float64))
	}
	if byte_s, err := json.Marshal(v.Policy); err == nil {
		policy_js := gjson.ParseBytes(byte_s)
		update_time := policy_js.Get("1.coin_seconds_earned_last_update").String()
		vesting_seconds := policy_js.Get("1.vesting_seconds").Int()
		t, _ := time.Parse(TIME_FORMAT, update_time)
		if time.Now().Unix()-t.In(UTCZone).Unix() < vesting_seconds {
			amount = (amount * uint64(time.Now().Unix()-t.In(UTCZone).Unix())) / (24 * 60 * 60)

		}
		return amount
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
	ID                             string `json:"id"`
	HeadBlockNumber                int    `json:"head_block_number"`
	HeadBlockID                    string `json:"head_block_id"`
	Time                           string `json:"time"`
	CurrentWitness                 string `json:"current_witness"`
	CurrentTransactionCount        int    `json:"current_transaction_count"`
	NextMaintenanceTime            string `json:"next_maintenance_time"`
	LastBudgetTime                 string `json:"last_budget_time"`
	WitnessBudget                  BigInt `json:"witness_budget"`
	AccountsRegisteredThisInterval int    `json:"accounts_registered_this_interval"`
	RecentlyMissedCount            int    `json:"recently_missed_count"`
	CurrentAslot                   int    `json:"current_aslot"`
	RecentSlotsFilled              string `json:"recent_slots_filled"`
	DynamicFlags                   int    `json:"dynamic_flags"`
	LastIrreversibleBlockNum       int    `json:"last_irreversible_block_num"`
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

type Votings struct {
	ID         string `json:"id"`
	Parameters struct {
		CurrentFees struct {
			Parameters         [][]interface{} `json:"parameters"`
			Scale              int             `json:"scale"`
			MaximunHandlingFee int             `json:"maximun_handling_fee"`
		} `json:"current_fees"`
		BlockInterval                 int `json:"block_interval"`
		MaintenanceInterval           int `json:"maintenance_interval"`
		MaintenanceSkipSlots          int `json:"maintenance_skip_slots"`
		CommitteeProposalReviewPeriod int `json:"committee_proposal_review_period"`
		MaximumBlockSize              int `json:"maximum_block_size"`
		MaximumTimeUntilExpiration    int `json:"maximum_time_until_expiration"`
		MaximumProposalLifetime       int `json:"maximum_proposal_lifetime"`
		MaximumAssetFeedPublishers    int `json:"maximum_asset_feed_publishers"`
		WitnessNumberOfElection       int `json:"witness_number_of_election"`
		CommitteeNumberOfElection     int `json:"committee_number_of_election"`
		MaximumAuthorityMembership    int `json:"maximum_authority_membership"`
		CashbackGasPeriodSeconds      int `json:"cashback_gas_period_seconds"`
		CashbackVbPeriodSeconds       int `json:"cashback_vb_period_seconds"`
		CashbackVotePeriodSeconds     int `json:"cashback_vote_period_seconds"`
		WitnessPayPerBlock            int `json:"witness_pay_per_block"`
		WitnessPayVestingSeconds      int `json:"witness_pay_vesting_seconds"`

		WorkerBudgetPerDay               BigInt        `json:"worker_budget_per_day"`
		AccountsPerFeeScale              int           `json:"accounts_per_fee_scale"`
		AccountFeeScaleBitshifts         int           `json:"account_fee_scale_bitshifts"`
		MaxAuthorityDepth                int           `json:"max_authority_depth"`
		MaximumRunTimeRatio              int           `json:"maximum_run_time_ratio"`
		MaximumNhAssetOrderExpiration    int           `json:"maximum_nh_asset_order_expiration"`
		AssignedTaskLifeCycle            int           `json:"assigned_task_life_cycle"`
		CrontabSuspendThreshold          int           `json:"crontab_suspend_threshold"`
		CrontabSuspendExpiration         int           `json:"crontab_suspend_expiration"`
		WitnessCandidateFreeze           string        `json:"witness_candidate_freeze"`
		CommitteeCandidateFreeze         string        `json:"committee_candidate_freeze"`
		CandidateAwardBudget             string        `json:"candidate_award_budget"`
		CommitteePercentOfCandidateAward int           `json:"committee_percent_of_candidate_award"`
		UnsuccessfulCandidatesPercent    int           `json:"unsuccessful_candidates_percent"`
		Extensions                       []interface{} `json:"extensions"`
	} `json:"parameters"`
	NextAvailableVoteID    int      `json:"next_available_vote_id"`
	ActiveCommitteeMembers []string `json:"active_committee_members"`
	ActiveWitnesses        []string `json:"active_witnesses"`
}
type VotingInfo struct {
	ID                    string        `json:"id"`
	WitnessAccount        string        `json:"witness_account"`
	LastAslot             int           `json:"last_aslot"`
	SigningKey            string        `json:"signing_key"`
	PayVb                 string        `json:"pay_vb"`
	VoteID                string        `json:"vote_id"`
	TotalVotes            string        `json:"total_votes"`
	URL                   string        `json:"url"`
	TotalMissed           int           `json:"total_missed"`
	LastConfirmedBlockNum int           `json:"last_confirmed_block_num"`
	WorkStatus            bool          `json:"work_status"`
	NextMaintenanceTime   string        `json:"next_maintenance_time"`
	Supporters            []interface{} `json:"supporters"`
}

func (o *Votings) GetInfo() [][]VotingInfo {
	witnesses_params := []interface{}{o.ActiveWitnesses}
	committee_params := []interface{}{o.ActiveCommitteeMembers}

	votingInfos := [][]VotingInfo{}
	req := CreateRpcRequest(CALL,
		[]interface{}{0, `get_objects`, witnesses_params})
	if resp, err := Client.Send(req); err == nil {
		VotingInfos := []VotingInfo{}
		if err = resp.GetInterface(&VotingInfos); err == nil {
			votingInfos = append(votingInfos, VotingInfos)
		}
	}
	req = CreateRpcRequest(CALL,
		[]interface{}{0, `get_objects`, committee_params})
	if resp, err := Client.Send(req); err == nil {
		VotingInfos := []VotingInfo{}
		if err = resp.GetInterface(&VotingInfos); err == nil {
			votingInfos = append(votingInfos, VotingInfos)
		}
	}
	return votingInfos
}

func GetVotingInfo() *Votings {
	req := CreateRpcRequest(CALL,
		[]interface{}{0, `get_objects`,
			[]interface{}{[]string{"2.0.0"}}})
	if resp, err := Client.Send(req); err == nil {
		VotingInfos := []*Votings{}
		if err = resp.GetInterface(&VotingInfos); err == nil {
			return VotingInfos[0]
		}
	}
	return nil
}

func GetObject(id string) *gjson.Result {

	req := CreateRpcRequest(CALL,
		[]interface{}{0, `get_objects`,
			[]interface{}{[]string{id}}})
	if resp, err := Client.Send(req); err == nil {
		ress := []interface{}{}
		if err := resp.GetInterface(&ress); err == nil {
			if byte_s, err := json.Marshal(ress[0]); err == nil {
				js := gjson.ParseBytes(byte_s)
				return &js
			}
		}
	}
	return nil
}

func GetCurrentFees() (m map[int64]gjson.Result) {
	defer func() {
		if recover() != nil {
			m = nil
		}
	}()
	req := CreateRpcRequest(CALL,
		[]interface{}{0, `get_objects`,
			[]interface{}{[]string{"2.0.0"}}})
	if resp, err := Client.Send(req); err == nil {
		VotingInfos := []*Votings{}
		if err = resp.GetInterface(&VotingInfos); err == nil {
			if byte_s, err := json.Marshal(VotingInfos[0].Parameters.CurrentFees.Parameters); err == nil {
				log.Println(string(byte_s))
				c_fees := gjson.ParseBytes(byte_s)
				m := make(map[int64]gjson.Result)
				for _, pair := range c_fees.Array() {
					m[pair.Get("0").Int()] = pair.Get("1")
				}
				return m
			}
		}
	}
	return nil
}
