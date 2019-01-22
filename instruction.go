package stackmachine

type instruction struct {
	name     string
	behavior func(*VM)
}

var keywords = map[string]instruction{
	"add": instruction{
		name:     "add",
		behavior: addInstruction,
	},
}

func addInstruction(vm *VM) {
	one := vm.dataStack.Pop()
	two := vm.dataStack.Pop()
	switch tone := one.(type) {
	case int64:
		switch ttwo := two.(type) {
		case int64:
			vm.dataStack.Push(tone + ttwo)
		default:
			panic("not valid data for add function")
		}
	case string:
		switch ttwo := two.(type) {
		case string:
			vm.dataStack.Push(tone + ttwo)
		default:
			panic("not valid data for add function")
		}
	default:
		panic("not valid data for add function")
	}
	vm.pc++
}
