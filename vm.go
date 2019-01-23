package stackmachine

type VM struct {
	pc              int
	dataStack       stack
	instructionList []instruction
}

func newVM() *VM {
	return &VM{
		pc:        0,
		dataStack: make(stack, 0),
	}
}

func (vm *VM) Run(instructions []instruction) {
	vm.instructionList = instructions
	for {
		if vm.pc >= len(vm.instructionList) {
			break
		}
		ins := vm.instructionList[vm.pc]
		ins.behavior(vm)
	}
}
