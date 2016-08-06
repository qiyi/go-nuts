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


