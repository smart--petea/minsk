package Binding

import (
    "minsk/CodeAnalysis/Binding/BoundUnaryOperator"
    "minsk/CodeAnalysis/Binding/BoundBinaryOperator"

    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/CodeAnalysis/Syntax"

    "minsk/Util"

    "fmt"
    "reflect"
)

type Binder struct {
    Util.DiagnosticBag
    _variables map[string]interface{}
}

func NewBinder(variables map[string]interface{}) *Binder {
    return &Binder{
        _variables: variables,
    }
}

func (b *Binder) BindExpression(syntax Syntax.ExpressionSyntax) BoundExpression {
    switch syntax.Kind() {
    case SyntaxKind.ParenthesizedExpression:
        return b.BindParenthesizedExpression(syntax)

    case SyntaxKind.LiteralExpression:
        return b.BindLiteralExpression(syntax)

    case SyntaxKind.NameExpression:
        return b.BindNameExpression(syntax)

    case SyntaxKind.AssignmentExpression:
        return b.BindAssignmentExpression(syntax)

    case SyntaxKind.UnaryExpression:
        return b.BindUnaryExpression(syntax)

    case SyntaxKind.BinaryExpression:
        return b.BindBinaryExpression(syntax)

    default:
        panic(fmt.Sprintf("Unexpected syntax %s", syntax.Kind()))
    }
}

func (b *Binder) BindAssignmentExpression(syntax Syntax.ExpressionSyntax) BoundExpression {
    assignmentExpressionSyntax := syntax.(*Syntax.AssignmentExpressionSyntax)

    name := string(assignmentExpressionSyntax.IdentifierToken.Runes)
    boundExpression := b.BindExpression(assignmentExpressionSyntax.Expression)

    var defaultValue interface{}
    switch boundExpression.GetType() {
    case reflect.Int:
        defaultValue = 0
    case reflect.Bool:
        defaultValue = false
    }

    if defaultValue == nil {
        panic(fmt.Sprintf("Unsuported variable type: %s", boundExpression.GetType()))
    }

    b._variables[name] = defaultValue
    return NewBoundAssignmentExpression(name, boundExpression)
}

func (b *Binder) BindNameExpression(syntax Syntax.ExpressionSyntax) BoundExpression {
    nameExpressionSyntax := syntax.(*Syntax.NameExpressionSyntax)
    name := string(nameExpressionSyntax.IdentifierToken.Runes)

    var value interface{}
    var ok bool
    if value, ok = b._variables[name]; !ok {
        b.ReportUndefinedName(nameExpressionSyntax.IdentifierToken.Span(), name)
        return NewBoundLiteralExpression(0)
    }

    return NewBoundVariableExpression(name, reflect.TypeOf(value).Kind())
}

func (b *Binder) BindParenthesizedExpression(syntax Syntax.ExpressionSyntax) BoundExpression {
    pS := syntax.(*Syntax.ParenthesizedExpressionSyntax)
    return b.BindExpression(pS.Expression)
}

func (b *Binder) BindLiteralExpression(syntax Syntax.ExpressionSyntax) BoundExpression {
    literalSyntax := syntax.(*Syntax.LiteralExpressionSyntax)

    var value interface{}

    switch literalSyntax.LiteralToken.Kind() {
    case SyntaxKind.TrueKeyword:
        value = true
    case SyntaxKind.FalseKeyword:
        value = false
    case SyntaxKind.IdentifierToken:
        value = string(literalSyntax.LiteralToken.Runes)
    default:
        if val, ok := literalSyntax.Value().(int); ok {
            value = val
        }
    }

    return NewBoundLiteralExpression(value)
}

func (b *Binder) BindUnaryExpression(syntax Syntax.ExpressionSyntax) BoundExpression {
    unarySyntax := syntax.(*Syntax.UnaryExpressionSyntax)
    boundOperand := b.BindExpression(unarySyntax.Operand)
    boundOperator := BoundUnaryOperator.Bind(unarySyntax.OperatorNode.Kind(), boundOperand.GetType()) 

    if boundOperator == nil {
        syntaxToken := unarySyntax.OperatorNode.(*Syntax.SyntaxToken)
        b.ReportUndefinedUnaryOperator(syntaxToken.Span(), syntaxToken.Runes, boundOperand.GetType())
        return boundOperand;
    }

    return NewBoundUnaryExpression(boundOperator, boundOperand)
}

func (b *Binder) BindBinaryExpression(syntax Syntax.ExpressionSyntax) BoundExpression {
    binarySyntax := (syntax).(*Syntax.BinaryExpressionSyntax)

    boundLeft := b.BindExpression(binarySyntax.Left)
    boundRight := b.BindExpression(binarySyntax.Right)
    boundOperator := BoundBinaryOperator.Bind(binarySyntax.OperatorNode.Kind(), boundLeft.GetType(), boundRight.GetType()) 

    if boundOperator == nil {
        syntaxToken := binarySyntax.OperatorNode.(*Syntax.SyntaxToken)
        b.ReportUndefinedBinaryOperator(syntaxToken.Span(), syntaxToken.Runes, boundLeft.GetType(), boundRight.GetType())

        return boundLeft;
    }

    return NewBoundBinaryExpression(boundLeft, boundOperator, boundRight)
}
