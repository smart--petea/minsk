package Util

type Stack struct {
    stack []interface{}
}

func (stack *Stack) Count() int {
    return len(stack.stack)
}

func (stack *Stack) Push(elem interface{}) {
    stack.stack = append(stack.stack, elem)
}

func (stack *Stack) Pop() interface{} {
    elem := stack.stack[len(stack.stack) - 1]
    stack.stack = stack.stack[:len(stack.stack) - 1]

    return elem
}

func NewStack() *Stack {
    return &Stack{}
}
