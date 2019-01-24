package stackmachine

import (
	"io/ioutil"
	"testing"
)

func TestVMRun(t *testing.T) {
	data, err := ioutil.ReadFile("/Users/yuzhao/Documents/test.s")
	if err != nil {
		panic(err)
	}
	l := newLexer(string(data))
	c := parseCode(l)
	vm := newVM()
	vm.Run(c)
}
