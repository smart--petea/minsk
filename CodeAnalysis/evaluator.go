package CodeAnalysis

import (
    "fmt"

    "minsk/CodeAnalysis/SyntaxKind"
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
    if n, ok := root.(*LiteralExpressionSyntax); ok {
        return n.LiteralToken.Value().(int)
    }

    if u, ok := root.(*UnaryExpressionSyntax); ok {
        operand := e.evaluateExpression(u.Operand)

        switch u.OperatorNode.Kind() {
        case SyntaxKind.PlusToken:
            return operand
        case SyntaxKind.MinusToken:
            return -operand
        default:
            panic(fmt.Sprintf("Unexpected unary operator %s", u.OperatorNode.Kind()))
        }
    }

    if b, ok := root.(*BinaryExpressionSyntax); ok {
        left := e.evaluateExpression(b.Left)
        right := e.evaluateExpression(b.Right)

        switch b.OperatorNode.Kind() {
        case SyntaxKind.PlusToken:
            return left + right
        case SyntaxKind.MinusToken:
            return left - right
        case SyntaxKind.StarToken:
            return left * right
        case SyntaxKind.SlashToken:
            return left / right
        default:
            panic(fmt.Sprintf("Unexpected binary operator %s", b.OperatorNode.Kind()))
        }
    }

    if p, ok := root.(*ParenthesizedExpressionSyntax); ok {
        result := e.evaluateExpression(p.Expression)

        return result
    }

    if s, ok := root.(*SyntaxToken); ok && s.Kind() == SyntaxKind.NumberToken{
        return s.Value().(int)
    }


    panic(fmt.Sprintf("Unexpected node %s", root.Kind()))
}
