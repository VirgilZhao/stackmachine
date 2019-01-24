package stackmachine

type code struct {
	instructionList []instruction
	labelMap        map[string]int
}

func newCode() *code {
	return &code{
		instructionList: make([]instruction, 0),
		labelMap:        make(map[string]int),
	}
}
