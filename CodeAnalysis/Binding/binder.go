package Binding

import (
    "minsk/CodeAnalysis/Binding/BoundUnaryOperator"
    "minsk/CodeAnalysis/Binding/BoundBinaryOperator"

    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/CodeAnalysis/Syntax"

    "minsk/Util"

    "fmt"
)

type Binder struct {
    Util.DiagnosticBag
}

func NewBinder() *Binder {
    return &Binder{}
}

func (b *Binder) BindExpression(syntax Syntax.ExpressionSyntax) BoundExpression {
    switch syntax.Kind() {
    case SyntaxKind.UnaryExpression:
        return b.BindUnaryExpression(syntax)
    case SyntaxKind.BinaryExpression:
        return b.BindBinaryExpression(syntax)
    case SyntaxKind.LiteralExpression:
        return b.BindLiteralExpression(syntax)
    case SyntaxKind.ParenthesizedExpression:
        return b.BindParenthesizedExpression(syntax)
    default:
        panic(fmt.Sprintf("Unexpected syntax %s", syntax.Kind()))
    }
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
