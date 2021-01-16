package Binding

import (
    "minsk/CodeAnalysis/Binding/Kind/BoundUnaryOperatorKind"
    "minsk/CodeAnalysis/Binding/Kind/BoundBinaryOperatorKind"
    "minsk/CodeAnalysis/Binding/TypeCarrier"

    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/CodeAnalysis/Syntax"

    "minsk/Util"

    "fmt"
)

type Binder struct {
    Util.Diagnostic
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

    var value int

    if val, ok := literalSyntax.LiteralToken.Value().(int); ok {
        value = val
    }

    return NewBoundLiteralExpression(value)
}

func (b *Binder) BindUnaryExpression(syntax Syntax.ExpressionSyntax) BoundExpression {
    unarySyntax := syntax.(*Syntax.UnaryExpressionSyntax)
    boundOperand := b.BindExpression(unarySyntax.Operand)
    boundOperatorKind := b.BindUnaryOperatorKind(unarySyntax.OperatorNode.Kind(), boundOperand.GetTypeCarrier()) 

    if boundOperatorKind == "" {
        b.AddDiagnostic("Unary operator '%+v' is not defined for type %T", unarySyntax.OperatorNode, boundOperand.GetTypeCarrier()) 
        return boundOperand;
    }

    return NewBoundUnaryExpression(boundOperatorKind, boundOperand)
}

func (b *Binder) BindBinaryExpression(syntax Syntax.ExpressionSyntax) BoundExpression {
    binarySyntax := (syntax).(*Syntax.BinaryExpressionSyntax)

    boundLeft := b.BindExpression(binarySyntax.Left)
    boundRight := b.BindExpression(binarySyntax.Right)
    boundOperatorKind := b.BindBinaryOperatorKind(binarySyntax.OperatorNode.Kind(), boundLeft.GetTypeCarrier(), boundRight.GetTypeCarrier()) 

    if boundOperatorKind == "" {
        b.AddDiagnostic("Binary operator '%+v' is not defined for types %T and %T", binarySyntax.OperatorNode, boundLeft.GetTypeCarrier(), boundRight.GetTypeCarrier()) 
        return boundLeft;
    }

    return NewBoundBinaryExpression(boundLeft, boundOperatorKind, boundRight)
}

func (b *Binder) BindUnaryOperatorKind(kind SyntaxKind.SyntaxKind, typeCarrier TypeCarrier.TypeCarrier) BoundUnaryOperatorKind.BoundUnaryOperatorKind {
    if TypeCarrier.IsInt(typeCarrier) == false {
        return ""
    } 

    switch kind {
    case SyntaxKind.PlusToken:
        return BoundUnaryOperatorKind.Identity
    case SyntaxKind.MinusToken:
        return BoundUnaryOperatorKind.Negation
    default:
        panic(fmt.Sprintf("Unexpected unary operator %s", kind))
    }
}

func (b *Binder) BindBinaryOperatorKind(kind SyntaxKind.SyntaxKind, leftTypeCarrier, rightTypeCarrier TypeCarrier.TypeCarrier) BoundBinaryOperatorKind.BoundBinaryOperatorKind {
    if TypeCarrier.IsInt(leftTypeCarrier) == false || TypeCarrier.IsInt(rightTypeCarrier) == false {
        return ""
    } 

    switch kind {
    case SyntaxKind.PlusToken:
        return BoundBinaryOperatorKind.Addition
    case SyntaxKind.MinusToken:
        return BoundBinaryOperatorKind.Subtraction
    case SyntaxKind.StarToken:
        return BoundBinaryOperatorKind.Multiplication
    case SyntaxKind.SlashToken:
        return BoundBinaryOperatorKind.Division
    default:
        panic(fmt.Sprintf("Unexpected binary operator %s", kind))
    }
}
