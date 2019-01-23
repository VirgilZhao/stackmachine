package stackmachine

import "fmt"

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
	printStack(*s)
}

func (s *stack) Pop() stackItemValue {
	targetStack := *s
	if len(targetStack) == 0 {
		panic("stack is already empty!")
	}
	value := targetStack[len(targetStack)-1]
	*s = targetStack[:len(targetStack)-1]
	printStack(*s)
	return value
}

func printStack(s stack) {
	for _, val := range s {
		switch item := val.(type) {
		case string:
			fmt.Printf("[\"%s\"]", item)
		case int64:
			fmt.Printf("[%d]", item)
		case bool:
			fmt.Printf("[%t]", val)
		}
	}
	if len(s) == 0 {
		fmt.Println("[]")
	} else {
		fmt.Println()
	}
}
