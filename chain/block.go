package chain

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
