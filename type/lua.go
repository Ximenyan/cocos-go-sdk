package types

import (
	"CocosSDK/common"
    "encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

const (
	INT = iota
	NUMBER
	STRING
	BOOL
	TABLE
	FUNCTION
)

type LuaObject struct {
	Type  uint64
	Value interface{}
}

func (o LuaObject) GetBytes() []byte {
	var byte_s []byte
	switch o.Type {
	case STRING:
		byte_s = append(common.Varint(o.Type), String(o.Value.(string)).GetBytes()...)
	case INT:
		byte_s = append(common.Varint(o.Type), common.VarInt(int64(o.Value.(int)), 64)...)
	case BOOL:
		byte_s = append(common.Varint(o.Type), 0)
		if o.Value.(bool) {
			byte_s[1] = 1
		}
	}
	return byte_s
}

type Table struct {
	Dict  map[string]interface{}
	Array []interface{}
}

func (o *LuaObject) GetTable() *Table {
	if o.Type != TABLE {
		return nil
	}
	if m, ok := o.Value.(map[*LuaObject]*LuaObject); ok {
		tab := &Table{Dict: make(map[string]interface{}), Array: []interface{}{}}
		for k, v := range m {
			if k.Type == INT {
				if v.Type == TABLE {
					tab.Array = append(tab.Array, v.GetTable())
				} else {
					tab.Array = append(tab.Array, v.Value)
				}
			}
			if k.Type == STRING {
				if v.Type == TABLE {
					tab.Dict[k.Value.(string)] = v.GetTable()
				} else {
					tab.Dict[k.Value.(string)] = v.Value
				}
			}
		}
		return tab
	}
	return nil

}

func (o *LuaObject) UnmarshalJSON(data []byte) (err error) {
	tmp := []interface{}{}
	if err = json.Unmarshal(data, &tmp); err != nil {
		return
	}
	if len(tmp) <= 0 {
		o.Type = TABLE
		o.Value = make(map[*LuaObject]*LuaObject)
		return
	}
	if o_type, ok := tmp[0].(float64); ok {
		m := make(map[string]interface{})
		if byte_s, err := json.Marshal(tmp[1]); err == nil {
			if err = json.Unmarshal(byte_s, &m); err != nil {
				return err
			}
		} else {
			return err
		}
		o.Type = uint64(o_type)
		o_value := m[`v`]
		switch o.Type {
		case INT:
			o.Value = uint64(o_value.(float64))
			return
		case NUMBER:
			o.Value = o_value.(float64)
			return
		case STRING:
			o.Value = o_value.(string)
			return
		case BOOL:
			o.Value = o_value.(bool)
			return
		case TABLE:
			o.Value = make(map[*LuaObject]*LuaObject) //o_value.()
			for _, item := range o_value.([]interface{}) {
				//log.Println("sda:", o.Value)
				k := item.([]interface{})[0].(map[string]interface{})[`key`]
				v := item.([]interface{})[1]
				o_item_k := new(LuaObject)
				o_item_v := new(LuaObject)
				byte_s_k, _ := json.Marshal(k)
				byte_s_v, _ := json.Marshal(v)
				o_item_k.UnmarshalJSON(byte_s_k)
				o_item_v.UnmarshalJSON(byte_s_v)
				o.Value.(map[*LuaObject]*LuaObject)[o_item_k] = o_item_v
			}
			return
		}
	} else {
		o.Type = TABLE
		o.Value = make(map[*LuaObject]*LuaObject)
		for _, item := range tmp {
			k := item.([]interface{})[0].(map[string]interface{})[`key`]
			v := item.([]interface{})[1]
			o_item_k := new(LuaObject)
			o_item_v := new(LuaObject)
			byte_s_k, _ := json.Marshal(k)
			byte_s_v, _ := json.Marshal(v)
			o_item_k.UnmarshalJSON(byte_s_k)
			o_item_v.UnmarshalJSON(byte_s_v)
			o.Value.(map[*LuaObject]*LuaObject)[o_item_k] = o_item_v
		}
		return
	}
	err = errors.New("the lua object type error!!!")
	return
}

func (o LuaObject) MarshalJSON() ([]byte, error) {
	var str string
	if o.Type == STRING {
		str_tmp := strings.Replace(o.Value.(string), `"`, `\"`, -1)
		str = fmt.Sprintf(`[%d,{"v":"%s"}]`, o.Type, str_tmp)
	} else if o.Type == NUMBER {
		str = fmt.Sprintf(`[%d,{"v":%f}]`, o.Type, o.Value)
	} else if o.Type == INT {
		str = fmt.Sprintf(`[%d,{"v":%d}]`, o.Type, o.Value)
	} else if o.Type == BOOL {
		if o.Value.(bool) {
			str = fmt.Sprintf(`[%d,{"v":"true"}]`, o.Type)
		} else {
			str = fmt.Sprintf(`[%d,{"v":"false"}]`, o.Type)
		}
	} else if o.Type == TABLE {
		str = `[4,{"v":[`
		for k, v := range o.Value.(map[*LuaObject]*LuaObject) {
			byte_s_k, _ := k.MarshalJSON()
			byte_s_v, _ := v.MarshalJSON()
			str_item := fmt.Sprintf(`[{"key":%s},%s],`, string(byte_s_k), string(byte_s_v))
			str += str_item
		}
		str = str[:len(str)-1] + `]}]`
	} else {
		fmt.Println(reflect.TypeOf(o.Value).String())
		str = fmt.Sprintf(`[%d,{"v":%s}]`, o.Type, o.Value)
	}
	return []byte(str), nil
}

func CreateLuaObject(v interface{}) *LuaObject {
	var t uint64
	type_str := reflect.TypeOf(v).Name()
	if strings.Index(type_str, "int") == 0 {
		t = 2
		v = strconv.Itoa(v.(int))
	} else if strings.Index(type_str, "uint") == 0 {
		t = 0
	} else if strings.Index(type_str, "float") == 0 {
		t = 2
		v = strconv.FormatFloat(v.(float64), 'f', -1, 64)
	} else if type_str == "string" {
		t = 2
	} else if type_str == "bool" {
		t = 2
		v = strconv.FormatBool(v.(bool))
	} else {
		t = 4
	}
	return &LuaObject{Type: t, Value: v}
}

type ValueList []*LuaObject

func (o ValueList) GetBytes() []byte {
	byte_s := common.Varint(uint64(len(o)))
	for i := 0; i < len(o); i++ {
		byte_s = append(byte_s, o[i].GetBytes()...)
	}
	return byte_s
}

func CreateValueList(values []interface{}) ValueList {
	value_list := []*LuaObject{}
	for i := 0; i < len(values); i++ {
		value_list = append(value_list, CreateLuaObject(values[i]))
	}
	return value_list
}
