package properties

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestLoad(t *testing.T) {
	bytes, err := ioutil.ReadFile("C:/Users/bphan/Documents/workspace/src/github.com/qiyi/go-nuts/test.properties")
	if err != nil {
		t.Errorf("Read file failed.{}", err)
	}

	properties, err := Load(bytes)
	if err != nil {
		t.Fail()
	}
	fmt.Println(properties.Get("k1"))
	fmt.Println(properties.Get("k2"))
	fmt.Println(properties.Get("k3"))
	fmt.Println(properties.Get("k4"))
	fmt.Println(properties.Get("k5"))
	fmt.Println(properties.Get("k6"))

	properties.Store("C:/Users/bphan/Documents/workspace/src/github.com/qiyi/go-nuts/test2.properties")
}
