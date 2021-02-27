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
        case BoundNodeKind.VariableDeclaration:
            e.evaluateVariableDeclaration(node.(*Binding.BoundVariableDeclaration))
        case BoundNodeKind.ExpressionStatement:
            e.evaluateExpressionStatement(node.(*Binding.BoundExpressionStatement))
        case BoundNodeKind.IfStatement:
            e.evaluateIfStatement(node.(*Binding.BoundIfStatement))
        case BoundNodeKind.WhileStatement:
            e.evaluateWhileStatement(node.(*Binding.BoundWhileStatement))
        case BoundNodeKind.ForStatement:
            e.evaluateForStatement(node.(*Binding.BoundForStatement))
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
        return equals(left, right, b.Left.GetType())

    case BoundBinaryOperatorKind.NotEquals:
       return !equals(left, right, b.Left.GetType())

    case BoundBinaryOperatorKind.Less:
        return left.(int) < right.(int)

    case BoundBinaryOperatorKind.LessOrEquals:
        return left.(int) <= right.(int)

    case BoundBinaryOperatorKind.Greater:
        return left.(int) > right.(int)

    case BoundBinaryOperatorKind.GreaterOrEquals:
        return left.(int) >= right.(int)

    default:
        panic(fmt.Sprintf("Unexpected binary operator %s", b.Op.Kind))
    }
}

func equals(left, right interface{}, ttype reflect.Kind) bool {
    switch ttype {
    case reflect.Bool:
        return left.(bool) == right.(bool)
    case reflect.Int:
        return left.(int) == right.(int)
    default:
        return false
    }
}

func (e *Evaluator) evaluateForStatement(node *Binding.BoundForStatement) {
    lowerBound := e.evaluateExpression(node.LowerBound).(int)
    upperBound := e.evaluateExpression(node.UpperBound).(int)

    for i := lowerBound; i <= upperBound; i = i + 1 {
        e.variables[node.Variable] = i
        e.evaluateStatement(node.Body)
    }
}

func (e *Evaluator) evaluateWhileStatement(node *Binding.BoundWhileStatement) {
    for e.evaluateExpression(node.Condition).(bool) {
        e.evaluateStatement(node.Body)
    }
}

func (e *Evaluator) evaluateIfStatement(node *Binding.BoundIfStatement) {
    condition := e.evaluateExpression(node.Condition).(bool)

    if condition {
        e.evaluateStatement(node.ThenStatement)
    } else if node.ElseStatement != nil {
        e.evaluateStatement(node.ElseStatement)
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

func (e *Evaluator) evaluateVariableDeclaration(node *Binding.BoundVariableDeclaration) {
    value := e.evaluateExpression(node.Initializer)
    e.variables[node.Variable] = value
    e.lastValue = value
}
