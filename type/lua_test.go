package types

import (
	"encoding/json"
)
import (
	"testing"
)

func TestLuaObj(t *testing.T) {
	str := `[[{"key":[2,{"v":"event_list"}]},[4,{"v":[[{"key":[0,{"v":1}]},[2,{"v":"{\"URL\":\"eeee\",\"DELAY\":\"2342dfd\",\"STATUS\":0,\"METHOD\":\"rrrr44\",\"CALLBACK\":\"3453\",\"PARAMS\":\"4445\"}"}]],[{"key":[0,{"v":2}]},[2,{"v":"{\"URL\":\"123\",\"DELAY\":\"2123123123\",\"STATUS\":0,\"ID\":2,\"METHOD\":\"2333\",\"CALLBACK\":\"333\",\"PARAMS\":\"444\"}"}]],[{"key":[0,{"v":3}]},[2,{"v":"{\"URL\":\"2ewqe\",\"DELAY\":\"531451\",\"STATUS\":0,\"ID\":3,\"METHOD\":\"234234\",\"CALLBACK\":\"23423534fghgh\",\"PARAMS\":\"erterdcsfcfdfgg\"}"}]],[{"key":[0,{"v":4}]},[2,{"v":"{\"URL\":\"xx\",\"DELAY\":\"weeq\",\"STATUS\":0,\"ID\":4,\"METHOD\":\"xxx\",\"CALLBACK\":\"ddd\",\"PARAMS\":\"sss\"}"}]],[{"key":[0,{"v":5}]},[2,{"v":"{\"URL\":\"xx\",\"DELAY\":\"weeq\",\"STATUS\":0,\"ID\":5,\"METHOD\":\"xxx\",\"CALLBACK\":\"ddd\",\"PARAMS\":\"sss\"}"}]]]}]],[{"key":[2,{"v":"now_id"}]},[0,{"v":5}]]]`
	l := &LuaObject{}
	json.Unmarshal([]byte(str), l)
	byte_s, _ := json.Marshal(l)
	t.Log(string(byte_s))
}
