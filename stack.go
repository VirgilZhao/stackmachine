package stackmachine

type stackItemValue interface{}

type stack []stackItemValue

func (s stack) Cap() int {
	return cap(s)
}

func (s stack) Get(idx int) (stackItemValue, bool) {
	if len(s) < idx+1 {
		return nil, false
	}
	return s[len(s)-idx-1], true
}

func (s *stack) Push(value stackItemValue) {
	*s = append(*s, value)
}

func (s *stack) Pop() stackItemValue {
	targetStack := *s
	if len(targetStack) == 0 {
		panic("stack is already empty!")
	}
	value := targetStack[len(targetStack)-1]
	*s = targetStack[:len(targetStack)-1]
	return value
}
