package CodeAnalysis

import (
    "fmt"

//    Syntax "minsk/CodeAnalysis/Syntax"
    Binding "minsk/CodeAnalysis/Binding"
 //   SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
)

type Evaluator struct {
        Root Binding.BoundExpression
}

func NewEvaluator(root Binding.BoundExpression) *Evaluator {
    return &Evaluator{
        Root: root,
    }
}

func (e *Evaluator) Evaluate() int {
    return e.evaluateExpression(e.Root)
}

func (e *Evaluator) evaluateExpression(root Binding.BoundExpression) int {
    if l, ok := root.(*Binding.BoundLiteralExpression); ok {
        return l.Value.(int)
    }

    if u, ok := root.(*Binding.BoundUnaryExpression); ok {
        operand := e.evaluateExpression(u.Operand)

        switch u.OperatorKind {
        case Binding.Identity:
            return operand
        case Binding.Negation:
            return -operand
        default:
            panic(fmt.Sprintf("Unexpected unary operator %s", u.OperatorKind))
        }
    }

    if b, ok := root.(*Binding.BoundBinaryExpression); ok {
        left := e.evaluateExpression(b.Left)
        right := e.evaluateExpression(b.Right)

        switch b.OperatorKind {
        case Binding.Addition:
            return left + right
        case Binding.Subtraction:
            return left - right
        case Binding.Multiplication:
            return left * right
        case Binding.Division:
            return left / right
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
