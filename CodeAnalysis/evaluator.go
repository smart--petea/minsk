package CodeAnalysis

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
    return e.evaluateExpression(e.Root)
}

func (e *Evaluator) evaluateExpression(root ExpressionSyntax) int {
    if n, ok := root.(*NumberExpressionSyntax); ok {
        return n.NumberToken.Value().(int)
    }

    if b, ok := root.(*BinaryExpressionSyntax); ok {
        left := e.evaluateExpression(b.Left)
        right := e.evaluateExpression(b.Right)

        switch b.OperatorNode.Kind() {
        case PlusToken:
            return left + right
        case MinusToken:
            return left - right
        case StarToken:
            return left * right
        case SlashToken:
            return left / right
        default:
            panic(fmt.Sprintf("Unexpected binary operator %s", b.OperatorNode.Kind()))
        }
    }

    if p, ok := root.(*ParenthesizedExpressionSyntax); ok {
        result := e.evaluateExpression(p.Expression)

        return result
    }

    if s, ok := root.(*SyntaxToken); ok && s.Kind() == NumberToken{
        return s.Value().(int)
    }


    panic(fmt.Sprintf("Unexpected node %s", root.Kind()))
}
