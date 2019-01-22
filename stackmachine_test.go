package stackmachine

import (
	"io/ioutil"
	"testing"
)

func TestCode(t *testing.T) {
	data, err := ioutil.ReadFile("/Users/yuzhao/Documents/test.s")
	if err != nil {
		panic(err)
	}
	newCode(string(data))
}

func TestVMExecute(t *testing.T) {
	data, err := ioutil.ReadFile("/Users/yuzhao/Documents/test.s")
	if err != nil {
		panic(err)
	}
	c := newCode(string(data))
	vm := newVM()
	vm.Execute(c)
}
