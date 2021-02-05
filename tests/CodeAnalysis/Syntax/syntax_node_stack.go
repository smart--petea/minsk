package SyntaxTest

import (
    "minsk/CodeAnalysis/Syntax"
)

type syntaxNodeStack struct {
    stack []Syntax.SyntaxNode
}

func (stack *syntaxNodeStack) Count() int {
    return len(stack.stack)
}

func (stack *syntaxNodeStack) Push(node Syntax.SyntaxNode) {
    stack.stack = append(stack.stack, node)
}

func (stack *syntaxNodeStack) Pop() Syntax.SyntaxNode {
    node := stack.stack[len(stack.stack) - 1]
    stack.stack = stack.stack[:len(stack.stack) - 1]

    return node
}
