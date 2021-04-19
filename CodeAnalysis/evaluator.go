package CodeAnalysis

import (
    "fmt"

    Binding "minsk/CodeAnalysis/Binding"
    "minsk/CodeAnalysis/Binding/Kind/BoundUnaryOperatorKind"
    "minsk/CodeAnalysis/Binding/Kind/BoundBinaryOperatorKind"
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
    "minsk/CodeAnalysis/Symbols"
    "minsk/CodeAnalysis/Symbols/BuiltinFunctions"
    "minsk/Util"
    "minsk/Util/Convert"
)

type Evaluator struct {
        Root *Binding.BoundBlockStatement

        globals Symbols.MapVariableSymbol
        locals *Symbols.StackMapVariableSymbol

        lastValue interface{}
        random *Util.Random
        functionBodies map[*Symbols.FunctionSymbol]*BoundBlockStatement
}

func NewEvaluator(functionBodies map[*Symbols.FunctionSymbol]*BoundBlockStatement, root *Binding.BoundBlockStatement, globals Symbols.MapVariableSymbol) *Evaluator {
    return &Evaluator{
        Root: root,
        globals: globals,
        locals: Symbols.NewStackMapVariableSymbol(),
        functionBodies: functionBodies,
    }
}

func (e *Evaluator) evaluateStatement(body *Binding.BoundBlockStatement) interface{} {
    labelToIndex := make(map[*Binding.BoundLabel]int)

    for i := range body.Statements {
        if l, ok := body.Statements[i].(*Binding.BoundLabelStatement); ok {
            labelToIndex[l.Label] = i + 1
        }
    }

    for index := 0; index < len(body.Statements); {
        s := body.Statements[index]
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

func (e *Evaluator) Evaluate() interface{} {
    return e.evaluateStatement(e.Root)
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
        case BoundNodeKind.CallExpression:
            return e.evaluateCallExpression(root.(*Binding.BoundCallExpression))
        case BoundNodeKind.ConversionExpression:
            return e.evaluateConversionExpression(root.(*Binding.BoundConversionExpression))
        default:
            panic(fmt.Sprintf("Unexpected node %s", root.Kind()))
    }
}

func (e *Evaluator) evaluateConversionExpression(node *Binding.BoundConversionExpression) interface{} {
    value := e.evaluateExpression(node.Expression)
    if node.GetType() == Symbols.TypeSymbolBool {
        return Convert.ToBoolean(value)
    } else if node.GetType() == Symbols.TypeSymbolInt {
        return Convert.ToInt(value)
    } else if node.GetType() == Symbols.TypeSymbolString {
        return Convert.ToString(value)
    } else {
        panic(fmt.Sprintf("Unexpected type %s", node.GetType()))
    }
}

func (e *Evaluator) evaluateCallExpression(node *Binding.BoundCallExpression) interface{} {
    if node.Function == BuiltinFunctions.Input {
        var x string
        fmt.Scanf("%s", &x)
        return x
    } else if node.Function == BuiltinFunctions.Print {
        message, _ := e.evaluateExpression(node.Arguments[0]).(string)
        fmt.Println(message)
        return nil
    } else if node.Function == BuiltinFunctions.Rnd {
        max, ok := e.evaluateExpression(node.Arguments[0]).(int)
        if !ok {
            panic("Arguments 0 can't be converted in int")
        }

        if e.random == nil {
            e.random = Util.NewRandom()
        }

        return e.random.Next(max)
    } else {
        locals := Symbols.NewMapVariableSymbol()

        args := make([]interface{}, len(node.Arguments))
        for i, _ := range node.Arguments {
            parameter := node.Function.Parameters[i]
            value = e.evaluateExpression(node.Arguments)
            locals.Add(parameter, value)
        }

        e.locals.Push(locals)
        statement := e.functionBodies[node.Function]

        return e.EvaluateStatement(statement)
    }
}

func (e *Evaluator) evaluateLiteralExpression(l *Binding.BoundLiteralExpression) interface{} {
    return l.Value
}

func (e *Evaluator) evaluateVariableExpression(v *Binding.BoundVariableExpression) interface{} {
    if v.Variable.Kind != SymbolKind.GlobalVariable {
        locals := e.locals.Peek()
        return locals[v.Variable]
    } else {
        return e.globals[v.Variable]
    }
}

func (e *Evaluator) evaluateAssignmentExpression(a *Binding.BoundAssignmentExpression) interface{} {
    value := e.evaluateExpression(a.Expression)

    if a.Variable.Kind == SymbolKind.GlobalVariable {
        e.globals[a.Variable] = value
    } else {
        locals := e.locals.Peek()
        locals[a.Variable] = value
    }

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
    e.globals[node.Variable] = value
    e.lastValue = value
}
