package slices

import (
	"testing"
)

func TestNotContainsStr(t *testing.T) {
	slice := []string{"a", "b", "c"}
	if !ContainsStr(slice, "b") {
		t.Fail()
	}
}

func TestContainsStr(t *testing.T) {
	slice := []string{"a", "b", "c"}
	if ContainsStr(slice, "d") {
		t.Fail()
	}
}

func TestIndexStrOK(t *testing.T) {
	slice := []string{"a", "b", "c"}
	if IndexStr(slice, "b") != 1 {
		t.Fail()
	}
}

func TestIndexStrNotExist(t *testing.T) {
	slice := []string{"a", "b", "c"}
	if IndexStr(slice, "d") != -1 {
		t.Fail()
	}
}

func TestDeleteStrAt0(t *testing.T) {
	var slice []string
	if _, ok := DeleteStrAt(slice, 0); ok {
		t.Fail()
	}
}

func TestDeleteStrAt0OK(t *testing.T) {
	slice := []string{"a"}
	if result, ok := DeleteStrAt(slice, 0); !ok {
		t.Fail()
	} else {
		if len(result) != 0 {
			t.Fail()
		}
	}
}

func TestInsertStrAtFail1(t *testing.T) {
	slice := []string{"b"}
	if _, ok := InsertStr(slice, -1, "a"); ok {
		t.Fail()
	}
}

func TestInsertStrAtFail2(t *testing.T) {
	slice := []string{"a"}
	if _, ok := InsertStr(slice, 2, "b"); ok {
		t.Fail()
	}
}

func TestInsertStrAtOK0(t *testing.T) {
	slice := []string{"a", "c"}
	if result, ok := InsertStr(slice, 1, "b"); !ok {
		t.Fail()
	} else {
		if len(result) != 3 && result[1] != "b" {
			t.Fail()
		}
	}

}

func TestInsertStrAtLast(t *testing.T) {
	slice := []string{"a", "b"}
	if result, ok := InsertStr(slice, 2, "c"); !ok {
		t.Fail()
	} else {
		if len(result) != 3 && result[2] != "c" {
			t.Fail()
		}
	}
}

func TestDeleteNilStrSlice(t *testing.T) {
	var slice []string
	if _, ok := DeleteStr(slice, "a"); ok {
		t.Fail()
	}
}

func TestDeleteEmptyStrSlice(t *testing.T) {
	slice := []string{}
	if _, ok := DeleteStr(slice, "a"); ok {
		t.Fail()
	}
}

func TestDeleteOneStrSlice(t *testing.T) {
	slice := []string{"a"}
	if result, ok := DeleteStr(slice, "a"); !ok {
		t.Fail()
	} else {
		if len(result) != 0 {
			t.Fail()
		}
	}
}

func TestAddFirstStrNilSlice(t *testing.T) {
	var slice []string
	result := AddFirstStr(slice, "a")
	if len(result) != 1 {
		t.Fail()
	}
}

func TestAddFirstStr(t *testing.T) {
	slice := []string{"b"}
	result := AddFirstStr(slice, "a")
	if result[0] != "a" {
		t.Fail()
	}
}

func TestNewStrSlice(t *testing.T) {
	result := NewStrSlice()
	if result == nil || len(result) != 0 {
		t.Fail()
	}
}

func TestForeachStrNil(t *testing.T) {
	ForeachStr(nil, func(int, string) {})
}

func TestForeachStrOK(t *testing.T) {
	i := ""
	slice := []string{"a", "b", "c"}
	ForeachStr(slice, func(_ int, v string) {
		i = i + v
	})
	if i != "abc" {
		t.Fail()
	}
}
