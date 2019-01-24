package stackmachine

type VM struct {
	pc        int
	dataStack stack
	code      *code
	labelDict map[string]int
}

func newVM() *VM {
	return &VM{
		pc:        0,
		dataStack: make(stack, 0),
	}
}

func (vm *VM) Run(c *code) {
	vm.code = c
	for {
		if vm.pc >= len(vm.code.instructionList) {
			break
		}
		ins := vm.code.instructionList[vm.pc]
		ins.behavior(vm)
	}
}
