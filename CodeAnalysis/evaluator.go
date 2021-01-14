package CodeAnalysis

import (
    "fmt"

    Syntax "minsk/CodeAnalysis/Syntax"
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
)

type Evaluator struct {
        Root BoundExpression
}

func NewEvaluator(root BoundExpression) *Evaluator {
    return &Evaluator{
        Root: root,
    }
}

func (e *Evaluator) Evaluate() int {
    return e.evaluateExpression(e.Root)
}

func (e *Evaluator) evaluateExpression(root BoundExpression) int {
    if l, ok := root.(*BoundLiteralExpression); ok {
        return l.Value().(int)
    }

    if u, ok := root.(*BoundUnaryExpression); ok {
        operand := e.evaluateExpression(u.Operand)

        switch u.OperatorKind.Kind() {
        case BoundUnaryOperatorKind.Identity:
            return operand
        case BoundUnaryOperatorKind.Negation:
            return -operand
        default:
            panic(fmt.Sprintf("Unexpected unary operator %s", u.OperatorKind.Kind()))
        }
    }

    if b, ok := root.(*BoundBinaryExpression); ok {
        left := e.evaluateExpression(b.Left)
        right := e.evaluateExpression(b.Right)

        switch b.OperatorKind {
        case BoundBinaryOperatorKind.Addition:
            return left + right
        case BoundBinaryOperatorKind.Subtraction:
            return left - right
        case BoundBinaryOperatorKind.Multiplication:
            return left * right
        case BoundBinaryOperatorKind.Division:
            return left / right
        default:
            panic(fmt.Sprintf("Unexpected binary operator %s", b.OperatorKind))
        }
    }

    if s, ok := root.(*Syntax.SyntaxToken); ok && s.Kind() == SyntaxKind.NumberToken{
        return s.Value().(int)
    }

    panic(fmt.Sprintf("Unexpected node %s", root.Kind()))
}
