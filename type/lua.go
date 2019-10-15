package types

import (
<<<<<<< HEAD
	"fmt"
	"reflect"
=======
	"cocos-go-sdk/common"
	"fmt"
	"reflect"
	"strconv"
>>>>>>> dev
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
<<<<<<< HEAD
	Type  uint
	Value interface{}
}

func (o LuaObject) String() string {
	if o.Type == STRING {
		return fmt.Sprintf(`[%d,{"v":"%s"}]`, o.Type, o.Value)
	} else if o.Type == NUMBER {
		return fmt.Sprintf(`[%d,{"v":%f}]`, o.Type, o.Value)
	} else if o.Type == INT {
		return fmt.Sprintf(`[%d,{"v":%d}]`, o.Type, o.Value)
	} else {
		fmt.Println(reflect.TypeOf(o.Value).String())
		return fmt.Sprintf(`[%d,{"v":%s}]`, o.Type, o.Value)
	}
}

func CreateLuaObject(v interface{}) *LuaObject {
	var t uint
	type_str := reflect.TypeOf(v).Name()
	if strings.Index(type_str, "int") == 0 {
		t = 0
	} else if strings.Index(type_str, "float") == 0 {
		t = 1
	} else if type_str == "string" {
		t = 2
	} else if type_str == "bool" {
		t = 3
=======
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

func (o LuaObject) MarshalJSON() ([]byte, error) {
	var str string
	if o.Type == STRING {
		str = fmt.Sprintf(`[%d,{"v":"%s"}]`, o.Type, o.Value)
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
>>>>>>> dev
	} else {
		t = 4
	}
	return &LuaObject{Type: t, Value: v}
}
<<<<<<< HEAD
=======

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
>>>>>>> dev
