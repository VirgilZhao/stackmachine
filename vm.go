package stackmachine

import (
	"strconv"
	"strings"
)

type VM struct {
	pc        int
	dataStack stack
}

func newVM() *VM {
	return &VM{
		pc:        0,
		dataStack: make(stack, 0),
	}
}

func (vm *VM) Execute(c *code) {
	for {
		if vm.pc >= len(c.tokens) {
			break
		}
		token := c.tokens[vm.pc]
		switch token.tokenType {
		case tokenTypeNumber:
			if strings.HasPrefix(token.value, "0x") || strings.HasPrefix(token.value, "0X") {
				if s, err := strconv.ParseInt(token.value, 0, 64); err == nil {
					vm.dataStack.Push(s)
				} else {
					panic("0x number convert error")
				}
			} else {
				if s, err := strconv.ParseInt(token.value, 10, 64); err == nil {
					vm.dataStack.Push(s)
				} else {
					panic("number convert error")
				}
			}
			vm.pc++
		case tokenTypeString:
			vm.dataStack.Push(token.value)
			vm.pc++
		case tokenTypeKeyword:
			if ins, ok := keywords[token.value]; ok {
				ins.behavior(vm)
			} else {
				panic("not valid keyword")
			}
		}
	}
}
