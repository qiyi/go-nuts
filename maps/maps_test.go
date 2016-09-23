package maps

import (
	"fmt"
	"testing"
)

func TestGetWithNil(t *testing.T) {
	if result, ok := Get(nil, "n"); ok || result != nil {
		t.Fail()
	}
}

func TestGetFail(t *testing.T) {
	m := NewMap()
	if result, ok := Get(m, "n"); result != nil || ok {
		t.Fail()
	}
}

func TestGet(t *testing.T) {
	m := NewMap()
	m["n"] = 1
	if result, ok := Get(m, "n"); result != 1 || !ok {
		t.Fail()
	}
}

func TestGetStr(t *testing.T) {
	m := NewMap()
	m["n"] = "1"
	if result, ok := GetStr(m, "n"); result != "1" || !ok {
		t.Fail()
	}
}

func TestGetStrFail(t *testing.T) {
	m := NewMap()
	if _, ok := GetStr(m, "n"); ok {
		t.Fail()
	}
}

func TestGetStrInvalidType(t *testing.T) {
	m := NewMap()
	m["n"] = 1
	if _, ok := GetStr(m, "n"); ok {
		t.Fail()
	}
}

func TestGetInt(t *testing.T) {
	m := NewMap()
	m["n"] = 1
	if result, ok := GetInt(m, "n"); !ok || result != 1 {
		t.Fail()
	}
}

func TestGetIntFail(t *testing.T) {
	m := NewMap()
	m["n"] = "1"
	if _, ok := GetInt(m, "n"); ok {
		t.Fail()
	}
}

func TestGetBool(t *testing.T) {
	m := NewMap()
	m["n"] = true
	if result, ok := GetBool(m, "n"); !ok || result != true {
		t.Fail()
	}
}

func TestGetBoolInvalidType(t *testing.T) {
	m := NewMap()
	m["n"] = 1
	if _, ok := GetBool(m, "n"); ok {
		t.Fail()
	}
}

func TestGetBoolFail(t *testing.T) {
	m := NewMap()
	if _, ok := GetBool(m, "n"); ok {
		t.Fail()
	}
}

func TestGetSlice(t *testing.T) {
	m := NewMap()
	m["n"] = []interface{}{1, 2}
	if result, ok := GetSlice(m, "n"); !ok {
		t.Fail()
	} else {
		if len(result) != 2 || result[0] != 1 || result[1] != 2 {
			t.Fail()
		}
	}
}

func TestGetSliceInvalidType(t *testing.T) {
	m := NewMap()
	m["n"] = []int{1, 2}
	if _, ok := GetSlice(m, "n"); ok {
		t.Fail()
	}
}

func TestGetSliceFail(t *testing.T) {
	m := NewMap()
	if _, ok := GetSlice(m, "n"); ok {
		t.Fail()
	}
}

func TestGetStrSlice(t *testing.T) {
	m := NewMap()
	m["n"] = []string{"a", "b"}
	if result, ok := GetStrSlice(m, "n"); !ok {
		t.Fail()
	} else {
		if len(result) != 2 || result[0] != "a" || result[1] != "b" {
			t.Fail()
		}
	}
}

func TestGetStrSlice2(t *testing.T) {
	m := NewMap()
	m["n"] = []interface{}{"a", "b"}
	if result, ok := GetStrSlice(m, "n"); !ok {
		t.Fail()
	} else {
		if len(result) != 2 || result[0] != "a" || result[1] != "b" {
			t.Fail()
		}
	}
}

func TestGetStrSliceInvalidType(t *testing.T) {
	m := NewMap()
	m["n"] = []int{1, 2}
	if _, ok := GetStrSlice(m, "n"); ok {
		t.Fail()
	}
}

func TestGetStrSliceFail(t *testing.T) {
	m := NewMap()
	if _, ok := GetStrSlice(m, "n"); ok {
		t.Fail()
	}
}

func TestGetMapSlice(t *testing.T) {
	m := NewMap()
	m["n"] = []interface{}{
		map[interface{}]interface{}{
			"a": 1,
		},
		map[interface{}]interface{}{
			"b": 2,
		},
	}

	if result, ok := GetMapSlice(m, "n"); !ok {
		t.Fail()
	} else {
		if len(result) != 2 || result[0]["a"] != 1 || result[1]["b"] != 2 {
			t.Fail()
		}
	}
}

func TestGetMapSlice2(t *testing.T) {
	m := NewMap()
	m["n"] = []map[interface{}]interface{}{
		map[interface{}]interface{}{
			"a": 1,
		},
		map[interface{}]interface{}{
			"b": 2,
		},
	}
	if result, ok := GetMapSlice(m, "n"); !ok {
		t.Fail()
	} else {
		if len(result) != 2 || result[0]["a"] != 1 || result[1]["b"] != 2 {
			t.Fail()
		}
	}
}

func TestGetMapSliceInvalidType(t *testing.T) {
	m := NewMap()
	m["n"] = []map[string]interface{}{
		map[string]interface{}{
			"a": 1,
		},
		map[string]interface{}{
			"b": 2,
		},
	}
	if result, ok := GetMapSlice(m, "n"); ok {
		fmt.Println(result)
		t.Fail()
	}
}

func TestGetMapSliceFail(t *testing.T) {
	m := NewMap()
	if _, ok := GetMapSlice(m, "n"); ok {
		t.Fail()
	}
}

func TestGetMap(t *testing.T) {
	m := NewMap()
	m["n"] = map[interface{}]interface{}{"a": 1}
	if result, ok := GetMap(m, "n"); !ok {
		t.Fail()
	} else {
		if result["a"] != 1 {
			t.Fail()
		}
	}
}

func TestGetMapInvalidType(t *testing.T) {
	m := NewMap()
	m["n"] = map[string]interface{}{"a": 1}
	if _, ok := GetMap(m, "n"); ok {
		t.Fail()
	}
}

func TestGetMapFail(t *testing.T) {
	m := NewMap()
	if _, ok := GetMap(m, "n"); ok {
		t.Fail()
	}
}

func TestCloneMap(t *testing.T) {
	m := NewMap()
	m["n"] = 1
	result := CloneMap(m)
	if m["n"] != result["n"] {
		t.Fail()
	}
	m["n"] = 2
	if m["n"] == result["n"] {
		t.Fail()
	}
}

func TestStrMap(t *testing.T) {
	m := NewStrMap()
	m["n"] = "m"
	result := CloneStrMap(m)
	if m["n"] != result["n"] {
		t.Fail()
	}
	m["n"] = "n"
	if m["n"] == result["n"] {
		t.Fail()
	}
}

func TestAsMap(t *testing.T) {
	m := NewStrMap()
	m["m"] = 1
	result := AsMap(m)
	if 1 != result["m"] {
		t.Fail()
	}
}

func TestAsStrMap(t *testing.T) {
	m := NewMap()
	m["m"] = 1
	if result, ok := AsStrMap(m); !ok || 1 != result["m"] {
		t.Fail()
	}
}

func TestAsStrMapFail(t *testing.T) {
	m := NewMap()
	m[1] = "m"
	if result, ok := AsStrMap(m); ok {
		fmt.Println(result)
		t.Fail()
	}

}
