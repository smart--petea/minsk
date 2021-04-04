package CodeAnalysis

import (
    "fmt"

    Binding "minsk/CodeAnalysis/Binding"
    "minsk/CodeAnalysis/Binding/Kind/BoundUnaryOperatorKind"
    "minsk/CodeAnalysis/Binding/Kind/BoundBinaryOperatorKind"
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
    "minsk/CodeAnalysis/Symbols"

    "log"
)

type Evaluator struct {
        Root *Binding.BoundBlockStatement
        variables map[*Symbols.VariableSymbol]interface{}
        lastValue interface{}
}

func NewEvaluator(root *Binding.BoundBlockStatement, variables map[*Symbols.VariableSymbol]interface{}) *Evaluator {
    return &Evaluator{
        Root: root,
        variables: variables,
    }
}

func (e *Evaluator) Evaluate() interface{} {
    labelToIndex := make(map[*Binding.BoundLabel]int)

    for i := range e.Root.Statements {
        if l, ok := e.Root.Statements[i].(*Binding.BoundLabelStatement); ok {
            labelToIndex[l.Label] = i + 1
        }
    }

    for index := 0; index < len(e.Root.Statements); {
        s := e.Root.Statements[index]
        switch s.Kind() {
            case BoundNodeKind.VariableDeclaration:
                e.evaluateVariableDeclaration(s.(*Binding.BoundVariableDeclaration))
                index = index + 1

            case BoundNodeKind.ExpressionStatement:
                e.evaluateExpressionStatement(s.(*Binding.BoundExpressionStatement))
                index = index + 1

            case BoundNodeKind.GotoStatement:
                gs := s.(*Binding.BoundGotoStatement)
                index = labelToIndex[gs.Label]

            case BoundNodeKind.ConditionalGotoStatement:
                cgs := s.(*Binding.BoundConditionalGotoStatement)
                condition := e.evaluateExpression(cgs.Condition).(bool)
                if condition && !cgs.JumpIfFalse || !condition && cgs.JumpIfFalse {
                    index = labelToIndex[cgs.Label]
                } else {
                    index = index + 1
                }

            case BoundNodeKind.LabelStatement:
                index = index + 1

            default:
                panic(fmt.Sprintf("Unexpected node %s", s.Kind()))
        }
    }

    return e.lastValue
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
    log.Printf("evaluateLiteralExpression %+v", l)
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
    log.Printf("evaluateUnaryExpression %+v operand=%+v kind=%+v", operand, u.Operand, u.Op.Kind)

    switch u.Op.Kind {
    case BoundUnaryOperatorKind.Identity:
        return operand.(int)
    case BoundUnaryOperatorKind.Negation:
        return -(operand.(int))
    case BoundUnaryOperatorKind.LogicalNegation:
        return !(operand.(bool))
    case BoundUnaryOperatorKind.OnesComplement:
        switch val := operand.(type) {
        case bool:
            if val {
                return false
            }

            return true

        case int:
            return ^val
        }


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
        if b.GetType() == Symbols.TypeSymbolInt {
            return left.(int) + right.(int)
        } else {
            return left.(string) + right.(string)
        }

    case BoundBinaryOperatorKind.Subtraction:
        return left.(int) - right.(int)

    case BoundBinaryOperatorKind.Multiplication:
        return left.(int) * right.(int)

    case BoundBinaryOperatorKind.Division:
        return left.(int) / right.(int)

    case BoundBinaryOperatorKind.BitwiseAnd:
        if b.GetType() == Symbols.TypeSymbolInt {
            return left.(int) & right.(int)
        }

        var leftOp, rightOp int
        if left.(bool) {
            leftOp = 1
        }

        if right.(bool) {
            rightOp = 1
        }

        if (leftOp & rightOp) > 0 {
            return true
        }

        return false

    case BoundBinaryOperatorKind.BitwiseOr:
        if b.GetType() == Symbols.TypeSymbolInt {
            return left.(int) | right.(int)
        }

        var leftOp, rightOp int
        if left.(bool) {
            leftOp = 1
        }

        if right.(bool) {
            rightOp = 1
        }

        if (leftOp | rightOp) > 0 {
            return true
        }

        return false

    case BoundBinaryOperatorKind.BitwiseXor:
        if b.GetType() == Symbols.TypeSymbolInt {
            return left.(int) ^ right.(int)
        }

        var leftOp, rightOp int
        if left.(bool) {
            leftOp = 1
        }

        if right.(bool) {
            rightOp = 1
        }

        if (leftOp ^ rightOp) > 0 {
            return true
        }

        return false

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

func equals(left, right interface{}, ttype *Symbols.TypeSymbol) bool {
    switch ttype {
    case Symbols.TypeSymbolBool:
        return left.(bool) == right.(bool)
    case Symbols.TypeSymbolInt:
        return left.(int) == right.(int)
    default:
        return false
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
