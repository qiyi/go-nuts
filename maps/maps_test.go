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

func TestGetStrMap(t *testing.T) {
	m := NewMap()
	m["n"] = map[string]interface{}{"a": 1}
	if result, ok := GetStrMap(m, "n"); !ok {
		t.Fail()
	} else {
		if result["a"] != 1 {
			t.Fail()
		}
	}
}

func TestGetStrMapInvalidType(t *testing.T) {
	m := NewMap()
	m["n"] = map[interface{}]interface{}{"a": 1}
	if _, ok := GetStrMap(m, "n"); ok {
		t.Fail()
	}
}

func TestGetStrMapFail(t *testing.T) {
	m := NewMap()
	if _, ok := GetStrMap(m, "n"); ok {
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

func TestVisitWithNil(t *testing.T) {
	if _, err := Visit(nil); err == nil {
		t.Fail()
	}
}

func TestWithNilPath(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": 1,
	}
	if r, err := Visit(m); err != nil {
		t.Fail()
	} else {
		if v, ok := r.(map[interface{}]interface{}); !ok {
			t.Fail()
		} else {
			if v["a"] != 1 {
				t.Fail()
			}
		}
	}
}

func TestVisit(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": 1,
	}
	if r, err := Visit(m, "a"); err != nil {
		t.Fail()
	} else {
		if r != 1 {
			t.Fail()
		}
	}
}

func TestVisitStr(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": "b",
	}

	if r, err := VisitStr(m, "a"); err != nil {
		t.Error("Get str value failed.")
	} else {
		if r != "b" {
			t.Error("Get str value incorrect.")
		}
	}

}

func TestVisitStrFail(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": 1,
	}

	if _, err := VisitStr(m, "a"); err == nil {
		t.Fail()
	}
}

func TestVisitStrNonExist(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": "b",
	}

	if _, err := VisitStr(m, "c"); err == nil {
		t.Fail()
	}
}

func TestVisitInt(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": 1,
	}

	if r, err := VisitInt(m, "a"); err != nil {
		t.Fail()
	} else {
		if r != 1 {
			t.Fail()
		}
	}
}

func TestVisitIntFail(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": "b",
	}

	if _, err := VisitInt(m, "a"); err == nil {
		t.Fail()
	}
}

func TestVisitIntNonExist(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": 1,
	}

	if _, err := VisitInt(m, "b"); err == nil {
		t.Fail()
	}
}

func TestVisitBool(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": true,
	}

	if r, err := VisitBool(m, "a"); err != nil {
		t.Fail()
	} else {
		if r != true {
			t.Fail()
		}
	}
}

func TestVisitBoolFail(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": "b",
	}

	if _, err := VisitBool(m, "a"); err == nil {
		t.Fail()
	}
}

func TestVisitBoolNonExist(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": true,
	}

	if _, err := VisitBool(m, "b"); err == nil {
		t.Fail()
	}
}

func TestVisitSlice(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": []interface{}{
			1, 2,
		},
	}

	if r, err := VisitSlice(m, "a"); err != nil {
		t.Error("Get slice failed")
	} else {
		if r[0] != 1 || r[1] != 2 {
			t.Error("Slice value incorrect")
		}
	}
}

func TestVisitSliceIncorrectType(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": []int{
			1, 2,
		},
	}

	if _, err := VisitSlice(m, "a"); err == nil {
		t.Fail()
	}
}

func TestVisitSliceNonExist(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": []interface{}{
			1, 2,
		},
	}

	if _, err := VisitSlice(m, "b"); err == nil {
		t.Fail()
	}
}

func TestVisitFromSlice(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": []interface{}{
			1, 2,
		},
	}

	if r, err := VisitInt(m, "a", 0); err != nil {
		t.Fail()
	} else {
		if r != 1 {
			t.Fail()
		}
	}
}

func TestVisitFromSliceOOIndex(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": []interface{}{
			1, 2,
		},
	}

	if _, err := VisitInt(m, "a", 2); err == nil {
		t.Fail()
	}
}

func TestVisitFromSliceNonIntIndex(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": []interface{}{
			1, 2,
		},
	}

	if _, err := VisitInt(m, "a", "b"); err == nil {
		t.Fail()
	}
}

func TestVisitStrSlice(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": []string{
			"c", "d",
		},
	}

	if r, err := VisitStrSlice(m, "a"); err != nil {
		t.Fail()
	} else {
		if r[0] != "c" || r[1] != "d" {
			t.Fail()
		}
	}
}

func TestVisitStrSliceIncorrectType(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": []int{
			1, 2,
		},
	}

	if _, err := VisitStrSlice(m, "a"); err == nil {
		t.Fail()
	}
}

func TestVisitStrSliceNonExist(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": []int{
			1, 2,
		},
	}

	if _, err := VisitStrSlice(m, "b"); err == nil {
		t.Fail()
	}
}

func TestVisitFromStrSlice(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": []string{
			"b", "c",
		},
	}

	if r, err := VisitStr(m, "a", 0); err != nil {
		t.Fail()
	} else {
		if r != "b" {
			t.Fail()
		}
	}
}

func TestVisitFromStrSliceOOIndex(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": []string{
			"b", "c",
		},
	}

	if _, err := VisitStr(m, "a", 2); err == nil {
		t.Fail()
	}
}

func TestVisitFromStrSliceNonIntIndex(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": []string{
			"b", "c",
		},
	}

	if _, err := VisitStr(m, "a", "b"); err == nil {
		t.Fail()
	}
}

func TestVisitMap(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": map[interface{}]interface{}{
			"b": 1,
		},
	}

	if r, err := VisitMap(m, "a"); err != nil {
		t.Fail()
	} else {
		if r["b"] != 1 {
			t.Fail()
		}
	}
}

func TestVisitMapIncorrectType(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": map[interface{}]int{
			"b": 1,
		},
	}

	if _, err := VisitMap(m, "a"); err == nil {
		t.Fail()
	}
}

func TestVisitMapNonExist(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": map[interface{}]interface{}{
			"b": 1,
		},
	}

	if _, err := VisitMap(m, "c"); err == nil {
		t.Fail()
	}
}

func TestVisitFromMap(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": 1,
	}

	if r, err := VisitInt(m, "a"); err != nil {
		t.Fail()
	} else if r != 1 {
		t.Fail()
	}
}

func TestVisitFromMapNonExist(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": 1,
	}

	if _, err := VisitInt(m, "b"); err == nil {
		t.Fail()
	}
}

func TestVisitStrMap(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": map[string]interface{}{
			"b": 1,
		},
	}

	if r, err := VisitStrMap(m, "a"); err != nil {
		t.Fail()
	} else {
		if r["b"] != 1 {
			t.Fail()
		}
	}
}

func TestVisitStrMapIncorrectType(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": map[interface{}]int{
			"b": 1,
		},
	}

	if _, err := VisitStrMap(m, "a"); err == nil {
		t.Fail()
	}
}

func TestVisitStrMapNonExist(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": map[string]interface{}{
			"b": 1,
		},
	}

	if _, err := VisitStrMap(m, "a", "c"); err == nil {
		t.Fail()
	}
}

func TestVisitFromStrMap(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": map[string]interface{}{
			"b": 1,
		},
	}

	if r, err := VisitInt(m, "a", "b"); err != nil {
		t.Fail()
	} else if r != 1 {
		t.Fail()
	}
}

func TestVisitFromStrMapNonExist(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": map[string]interface{}{
			"b": 1,
		},
	}

	if _, err := VisitInt(m, "a", "c"); err == nil {
		t.Fail()
	}
}

func TestVisitMapSlice(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": []interface{}{
			map[interface{}]interface{}{"b": 1},
			map[interface{}]interface{}{"c": 2},
		},
	}

	if r, err := VisitMapSlice(m, "a"); err != nil {
		t.Fail()
	} else {
		if r[0]["b"] != 1 {
			t.Fail()
		}
	}
}

func TestVisitMapSlice2(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": []map[interface{}]interface{}{
			map[interface{}]interface{}{"b": 1},
			map[interface{}]interface{}{"c": 2},
		},
	}
	if r, err := VisitMapSlice(m, "a"); err != nil {
		t.Fail()
	} else {
		if r[0]["b"] != 1 {
			t.Fail()
		}
	}
}

func TestVisitMapSliceFail(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": 1,
	}
	if _, err := VisitMapSlice(m, "a"); err == nil {
		t.Fail()
	}
}

func TestVisitFromMapSlice(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": []map[interface{}]interface{}{
			map[interface{}]interface{}{
				"b": 1,
			},
		},
	}

	if r, err := VisitInt(m, "a", 0, "b"); err != nil {
		t.Fail()
	} else if r != 1 {
		t.Fail()
	}
}

func TestVisitFromMapSliceFail(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": []map[interface{}]interface{}{
			map[interface{}]interface{}{
				"b": 1,
			},
		},
	}

	if _, err := VisitInt(m, "a", 1, "b"); err == nil {
		t.Fail()
	}
}

func TestVisitStrMapSlice(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": []interface{}{
			map[string]interface{}{"b": 1},
			map[string]interface{}{"c": 2},
		},
	}

	if r, err := VisitStrMapSlice(m, "a"); err != nil {
		t.Fail()
	} else {
		if r[0]["b"] != 1 {
			t.Fail()
		}
	}
}

func TestVisitStrMapSlice2(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": []map[string]interface{}{
			map[string]interface{}{"b": 1},
			map[string]interface{}{"c": 2},
		},
	}

	if r, err := VisitStrMapSlice(m, "a"); err != nil {
		t.Fail()
	} else {
		if r[0]["b"] != 1 {
			t.Fail()
		}
	}
}

func TestVisitStrMapSliceFail(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": 1,
	}

	if _, err := VisitStrMapSlice(m, "a"); err == nil {
		t.Fail()
	}
}

func TestVisitFromStrMapSlice(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": []map[string]interface{}{
			map[string]interface{}{"b": 1},
		},
	}

	if r, err := VisitInt(m, "a", 0, "b"); err != nil {
		t.Fail()
	} else if r != 1 {
		t.Fail()
	}
}

func TestVisitFromStrMapSliceFali(t *testing.T) {
	m := map[interface{}]interface{}{
		"a": []map[string]interface{}{
			map[string]interface{}{"b": 1},
		},
	}

	if _, err := VisitInt(m, "a", 1, "b"); err == nil {
		t.Fail()
	}
}

func TestVisitFail(t *testing.T) {
	type x struct{}
	m := map[interface{}]interface{}{
		"a": x{},
	}
	if _, err := Visit(m, "a", "x"); err == nil {
		t.Fail()
	}
}
