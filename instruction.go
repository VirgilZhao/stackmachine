package stackmachine

import (
	"strconv"
	"strings"
)

type instruction struct {
	name     string
	arity    int
	values   []token
	behavior func(*VM)
}

var keywords = map[string]instruction{
	"push": instruction{
		name:     "push",
		arity:    1,
		behavior: pushInstruction,
	},
	"add": instruction{
		name:     "add",
		arity:    0,
		behavior: addInstruction,
	},
	"sub": instruction{
		name:     "sub",
		arity:    0,
		behavior: subInstruction,
	},
	"mul": instruction{
		name:     "mul",
		arity:    0,
		behavior: mulInstruction,
	},
	"div": instruction{
		name:     "div",
		arity:    0,
		behavior: divInstruction,
	},
	"lt": instruction{
		name:     "lt",
		arity:    0,
		behavior: lessThanInstruction,
	},
	"gt": instruction{
		name:     "gt",
		arity:    0,
		behavior: greaterThanInstruction,
	},
	"le": instruction{
		name:     "le",
		arity:    0,
		behavior: lessEqualInstruction,
	},
	"ge": instruction{
		name:     "ge",
		arity:    0,
		behavior: greaterEqualInstruction,
	},
	"jp": instruction{
		name:     "jp",
		arity:    1,
		behavior: jumpInstruction,
	},
	"jpn": instruction{
		name:     "jpn",
		arity:    1,
		behavior: jumpIfNotInstruction,
	},
}

func pushInstruction(vm *VM) {
	ins := vm.instructionList[vm.pc]
	token := ins.values[0]
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
	case tokenTypeBoolean:
		if token.value == "true" {
			vm.dataStack.Push(true)
		} else {
			vm.dataStack.Push(false)
		}
		vm.pc++
	default:
		panic("not valid token type")
	}
}

func addInstruction(vm *VM) {
	two := vm.dataStack.Pop()
	one := vm.dataStack.Pop()
	switch tone := one.(type) {
	case int64:
		switch ttwo := two.(type) {
		case int64:
			vm.dataStack.Push(tone + ttwo)
		default:
			panic("not valid data for add")
		}
	case string:
		switch ttwo := two.(type) {
		case string:
			vm.dataStack.Push(tone + ttwo)
		default:
			panic("not valid data for add")
		}
	default:
		panic("not valid data for add")
	}
	vm.pc++
}

func subInstruction(vm *VM) {
	two := vm.dataStack.Pop()
	one := vm.dataStack.Pop()
	switch tone := one.(type) {
	case int64:
		switch ttwo := two.(type) {
		case int64:
			vm.dataStack.Push(tone + ttwo)
		default:
			panic("not valid data for sub")
		}
	default:
		panic("not valid data for sub")
	}
	vm.pc++
}

func mulInstruction(vm *VM) {
	two := vm.dataStack.Pop()
	one := vm.dataStack.Pop()
	switch tone := one.(type) {
	case int64:
		switch ttwo := two.(type) {
		case int64:
			vm.dataStack.Push(tone * ttwo)
		default:
			panic("not valid data for sub")
		}
	default:
		panic("not valid data for sub")
	}
	vm.pc++
}

func divInstruction(vm *VM) {
	two := vm.dataStack.Pop()
	one := vm.dataStack.Pop()
	switch tone := one.(type) {
	case int64:
		switch ttwo := two.(type) {
		case int64:
			vm.dataStack.Push(tone / ttwo)
		default:
			panic("not valid data for sub")
		}
	default:
		panic("not valid data for sub")
	}
	vm.pc++
}

func equalInstruction(vm *VM) {
	compareInstruction(vm, compareEqual)
}

func lessThanInstruction(vm *VM) {
	compareInstruction(vm, compareLessThan)
}

func greaterThanInstruction(vm *VM) {
	compareInstruction(vm, compareGreaterThan)
}

func lessEqualInstruction(vm *VM) {
	compareInstruction(vm, compareLessEqual)
}

func greaterEqualInstruction(vm *VM) {
	compareInstruction(vm, compareGreaterEqual)
}

const (
	compareEqual = iota
	compareLessThan
	compareGreaterThan
	compareLessEqual
	compareGreaterEqual
)

func compareInstruction(vm *VM, compare int) {
	two := vm.dataStack.Pop()
	one := vm.dataStack.Pop()
	switch tone := one.(type) {
	case int64:
		switch ttwo := two.(type) {
		case int64:
			vm.dataStack.Push(intCompare(tone, ttwo, compare))
		default:
			panic("not valid data for compare operator")
		}
	case string:
		switch ttwo := two.(type) {
		case string:
			vm.dataStack.Push(stringCompare(tone, ttwo, compare))
		default:
			panic("not valid data for compare operator")
		}
	case bool:
		switch ttwo := two.(type) {
		case bool:
			vm.dataStack.Push(boolCompare(tone, ttwo, compare))
		default:
			panic("not valid data for compare operator")
		}
	default:
		panic("not valid data for compare operator")
	}
	vm.pc++
}

func intCompare(a, b int64, cmp int) bool {
	switch cmp {
	case compareEqual:
		return a == b
	case compareLessThan:
		return a < b
	case compareGreaterThan:
		return a > b
	case compareLessEqual:
		return a <= b
	case compareGreaterEqual:
		return a >= b
	default:
		panic("not valid compare operator")
	}
}

func stringCompare(a, b string, cmp int) bool {
	switch cmp {
	case compareEqual:
		return a == b
	case compareLessThan:
		return a < b
	case compareGreaterThan:
		return a > b
	case compareLessEqual:
		return a <= b
	case compareGreaterEqual:
		return a >= b
	default:
		panic("not valid compare operator")
	}
}

func boolCompare(a, b bool, cmp int) bool {
	switch cmp {
	case compareEqual:
		return a == b
	default:
		panic("not valid compare operator")
	}
}

func jumpInstruction(vm *VM) {
	ins := vm.instructionList[vm.pc]
	tk := ins.values[0]
	if jumpCount, err := strconv.ParseInt(tk.value, 10, 64); err == nil {
		vm.pc += int(jumpCount) + 1
	} else {
		panic("not valid jump param")
	}
}

func jumpIfNotInstruction(vm *VM) {
	ins := vm.instructionList[vm.pc]
	tk := ins.values[0]
	if jumpCount, err := strconv.ParseInt(tk.value, 10, 64); err == nil {
		val := vm.dataStack.Pop()
		jump := false
		switch tval := val.(type) {
		case int64:
			if tval != 1 {
				jump = true
			}
		case bool:
			if !tval {
				jump = true
			}
		default:
			panic("not valid jump condition")
		}
		if jump {
			vm.pc += int(jumpCount) + 1
		} else {
			vm.pc++
		}
	} else {
		panic("not valid jump if not param")
	}
}
