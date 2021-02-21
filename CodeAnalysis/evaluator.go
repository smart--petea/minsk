package CodeAnalysis

import (
    "fmt"

    Binding "minsk/CodeAnalysis/Binding"
    "minsk/CodeAnalysis/Binding/Kind/BoundUnaryOperatorKind"
    "minsk/CodeAnalysis/Binding/Kind/BoundBinaryOperatorKind"
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
    "minsk/Util"

    "reflect"
)

type Evaluator struct {
        Root Binding.BoundStatement
        variables map[*Util.VariableSymbol]interface{}
        lastValue interface{}
}

func NewEvaluator(root Binding.BoundStatement, variables map[*Util.VariableSymbol]interface{}) *Evaluator {
    return &Evaluator{
        Root: root,
        variables: variables,
    }
}

func (e *Evaluator) Evaluate() interface{} {
    e.evaluateStatement(e.Root)

    return e.lastValue
}

func (e *Evaluator) evaluateStatement(node Binding.BoundStatement) {
    switch node.Kind() {
        case BoundNodeKind.BlockStatement:
            e.evaluateBlockStatement(node.(*Binding.BoundBlockStatement))
        case BoundNodeKind.ExpressionStatement:
            e.evaluateExpressionStatement(node.(*Binding.BoundExpressionStatement))
        default:
            panic(fmt.Sprintf("Unexpected node %s", node.Kind()))
    }
}

func (e *Evaluator) evaluateExpression(root Binding.BoundExpression) interface{} {
    switch root.Kind() {
        case BoundNodeKind.LiteralExpression:
            return e.evaluateLiteralExpression(root.(*Binding.BoundLiteralExpression))
        case BoundNodeKind.VariableExpression:
            return e.evaluateVariableExpression(root.(*Binding.BoundVariableExpression))
        case BoundNodeKind.AssignmentExpression:
            return e.evaluateAssignmentExpression(root.(*Binding.BoundAssignmentExpression))
        case BoundNodeKind.UnaryExpression:
            return e.evaluateUnaryExpression(root.(*Binding.BoundUnaryExpression))
        case BoundNodeKind.BinaryExpression:
            return e.evaluateBinaryExpression(root.(*Binding.BoundBinaryExpression))
        default:
            panic(fmt.Sprintf("Unexpected node %s", root.Kind()))
    }
}

func (e *Evaluator) evaluateLiteralExpression(l *Binding.BoundLiteralExpression) interface{} {
    return l.Value
}

func (e *Evaluator) evaluateVariableExpression(v *Binding.BoundVariableExpression) interface{} {
    return e.variables[v.Variable]
}

func (e *Evaluator) evaluateAssignmentExpression(a *Binding.BoundAssignmentExpression) interface{} {
    value := e.evaluateExpression(a.Expression)
    e.variables[a.Variable] = value
    return value
}

func (e *Evaluator) evaluateUnaryExpression(u *Binding.BoundUnaryExpression) interface{} {
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

func (e *Evaluator) evaluateBinaryExpression(b *Binding.BoundBinaryExpression) interface{} {
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

func (e *Evaluator) evaluateBlockStatement(node *Binding.BoundBlockStatement) {
    for _, statement := range node.Statements {
        e.evaluateStatement(statement)
    }
}

func (e *Evaluator) evaluateExpressionStatement(node *Binding.BoundExpressionStatement) {
    e.lastValue = e.evaluateExpression(node.Expression)
}
