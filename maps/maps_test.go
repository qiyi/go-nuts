package maps

import "testing"

func TestGetStr(t *testing.T) {
    m := NewMap()
    m["a"] = "1"
    result, _ := GetStr(m, "a")
    if result != "1" {
        t.Fail()
    }
}
