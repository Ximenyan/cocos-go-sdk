package rpc

import . "cocos-go-sdk/type"

const CALL = `call`
const EMPTY = ""

func GetRequireFeeData(opID int, t OpData) *Amount {
	fee := &[]*Amount{}
	req := CreateRpcRequest(CALL,
		[]interface{}{0, `get_required_fees`,
			[]interface{}{[]interface{}{[]interface{}{opID, t}}, "1.3.0"}})
	if resp, err := Client.Send(req); err == nil {
		if err = resp.GetInterface(fee); err == nil {
			t.SetFee((*fee)[0].Amount)
			return (*fee)[0]
		}
		return nil
	}

	return nil
}

func GetTransactionById(txId string) {
	//fee := &[]*Amount{}
	req := CreateRpcRequest(CALL,
		[]interface{}{0, `get_transaction_by_id`,
			[]interface{}{txId}})
	if _, err := Client.Send(req); err == nil {
		//if err = resp.GetInterface(fee); err == nil {
		//	return
		//}
		return
	}

	return
}
