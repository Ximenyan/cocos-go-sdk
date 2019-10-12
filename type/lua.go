package types

import (
	"fmt"
	"reflect"
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
	} else {
		t = 4
	}
	return &LuaObject{Type: t, Value: v}
}
