package minsk

import (
    "fmt"
)

type Evaluator struct {
        Root ExpressionSyntax
}

func NewEvaluator(root ExpressionSyntax) *Evaluator {
    return &Evaluator{
        Root: root,
    }
}

func (e *Evaluator) Evaluate() int {
    return p.evaluateExpression(e.Root)
}

func (e *Evaluator) evaluateExpression(root ExpressionSyntax) int {
    switch adjusted := root.(type) {
    case NumberExpressionSyntax:
        return adjusted.NumberToken.Value().(int)

    case BinaryExpressionSyntax:
        left := e.evaluateExpression(adjusted.Left)
        right := e.evaluateExpression(adjusted.Right)

        switch adjusted.OperatorNode.Kind() {
        case PlusToken:
            return left + right
        case MinusToken:
            return left - right
        case StarToken:
            return left * right
        case SlashToken:
            return left / right
        default:
            panic(fmt.Sprintf("Unexpected binary operator %s", adjusted.OperatorNode.Kind()))
        }
    default:
        panic(fmt.Sprintf("Unexpected node %s", adjusted.Kind()))
    }
}
