package CodeAnalysis

import (
    "fmt"

    Binding "minsk/CodeAnalysis/Binding"
    "minsk/CodeAnalysis/Binding/Kind/BoundUnaryOperatorKind"
    "minsk/CodeAnalysis/Binding/Kind/BoundBinaryOperatorKind"
)

type Evaluator struct {
        Root Binding.BoundExpression
}

func NewEvaluator(root Binding.BoundExpression) *Evaluator {
    return &Evaluator{
        Root: root,
    }
}

func (e *Evaluator) Evaluate() interface{} {
    return e.evaluateExpression(e.Root)
}

func (e *Evaluator) evaluateExpression(root Binding.BoundExpression) interface{} {
    if l, ok := root.(*Binding.BoundLiteralExpression); ok {
        return l.Value
    }

    if u, ok := root.(*Binding.BoundUnaryExpression); ok {
        operand := e.evaluateExpression(u.Operand)

        switch u.OperatorKind {
        case BoundUnaryOperatorKind.Identity:
            return operand.(int)
        case BoundUnaryOperatorKind.Negation:
            return -(operand.(int))
        case BoundUnaryOperatorKind.LogicalNegation:
            return !(operand.(bool))
        default:
            panic(fmt.Sprintf("Unexpected unary operator %s", u.OperatorKind))
        }
    }

    if b, ok := root.(*Binding.BoundBinaryExpression); ok {
        left := e.evaluateExpression(b.Left)
        right := e.evaluateExpression(b.Right)

        switch b.OperatorKind {
        case BoundBinaryOperatorKind.Addition:
            return left.(int) + right.(int)

        case BoundBinaryOperatorKind.Subtraction:
            return left.(int) - right.(int)

        case BoundBinaryOperatorKind.Multiplication:
            return left.(int) * right.(int)

        case BoundBinaryOperatorKind.Division:
            return left.(int) / right.(int)

        case BoundBinaryOperatorKind.LogicalAnd:
            return left.(bool) && right.(bool)

        case BoundBinaryOperatorKind.LogicalOr:
            return left.(bool) || right.(bool)

        default:
            panic(fmt.Sprintf("Unexpected binary operator %s", b.OperatorKind))
        }
    }

    /*todo
    if s, ok := root.(*Syntax.SyntaxToken); ok && s.Kind() == SyntaxKind.NumberToken{
        return s.Value().(int)
    }
    */

    panic(fmt.Sprintf("Unexpected node %s", root.Kind()))
}
