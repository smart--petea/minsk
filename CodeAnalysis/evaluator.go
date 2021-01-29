package CodeAnalysis

import (
    "fmt"

    Binding "minsk/CodeAnalysis/Binding"
    "minsk/CodeAnalysis/Binding/Kind/BoundUnaryOperatorKind"
    "minsk/CodeAnalysis/Binding/Kind/BoundBinaryOperatorKind"

    "reflect"
)

type Evaluator struct {
        Root Binding.BoundExpression
        _variables map[string]interface{}
}

func NewEvaluator(root Binding.BoundExpression, variables map[string]interface{}) *Evaluator {
    return &Evaluator{
        Root: root,
        _variables: variables,
    }
}

func (e *Evaluator) Evaluate() interface{} {
    return e.evaluateExpression(e.Root)
}

func (e *Evaluator) evaluateExpression(root Binding.BoundExpression) interface{} {
    if l, ok := root.(*Binding.BoundLiteralExpression); ok {
        return l.Value
    }

    if v, ok := root.(*Binding.BoundVariableExpression); ok {
        return e._variables[v.Name]
    }

    if a, ok := root.(*Binding.BoundAssignmentExpression); ok {
        value := e.evaluateExpression(a.Expression)
        e._variables[a.Name] = value
        return value
    }

    if u, ok := root.(*Binding.BoundUnaryExpression); ok {
        operand := e.evaluateExpression(u.Operand)

        switch u.Op.Kind {
        case BoundUnaryOperatorKind.Identity:
            return operand.(int)
        case BoundUnaryOperatorKind.Negation:
            return -(operand.(int))
        case BoundUnaryOperatorKind.LogicalNegation:
            return !(operand.(bool))
        default:
            panic(fmt.Sprintf("Unexpected unary operator %s", u.Op.Kind))
        }
    }

    if b, ok := root.(*Binding.BoundBinaryExpression); ok {
        left := e.evaluateExpression(b.Left)
        right := e.evaluateExpression(b.Right)

        switch b.Op.Kind {
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

        case BoundBinaryOperatorKind.Equals:
            if b.Left.GetType() == reflect.Bool {
                return left.(bool) == right.(bool)
            } else {
                return left.(int) == right.(int)
            }

        case BoundBinaryOperatorKind.NotEquals:
            if b.Left.GetType() == reflect.Bool {
                return left.(bool) != right.(bool)
            } else {
                return left.(int) != right.(int)
            }

        default:
            panic(fmt.Sprintf("Unexpected binary operator %s", b.Op.Kind))
        }
    }

    panic(fmt.Sprintf("Unexpected node %s", root.Kind()))
}
