package utils

type StringStack struct {
	bytes []string
	count int
}

func NewStringStack() StringStack {
	return StringStack{bytes: []string{}}
}

func (stack *StringStack) Empty() bool {
	return stack.count == 0
}
func (stack *StringStack) Push(b string) {
	stack.bytes = append(stack.bytes, b)
	stack.count++
}

func (stack *StringStack) Pop() string {
	pop := stack.bytes[stack.count-1]
	stack.bytes = stack.bytes[:stack.count-1]
	stack.count--

	return pop
}
