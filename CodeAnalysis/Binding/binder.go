package Binding

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    Syntax "minsk/CodeAnalysis/Syntax"

    "fmt"
)

type BoundNodeKind string

const (
    UnaryExpression BoundNodeKind = "UnaryExpression"
    LiteralExpression BoundNodeKind = "LiteralExpression"
    BinaryExpression  BoundNodeKind = "BinaryExpression"
)

type BoundNode interface {
    Kind() BoundNodeKind
}

type TypeCarrier interface{}

type BoundExpression interface {
    BoundNode

    GetTypeCarrier() TypeCarrier
}

//todo move diagnostic in a separate class. should be nested also for lexer and parser

type BoundUnaryExpression struct {
    Operand BoundExpression
    OperatorKind BoundUnaryOperatorKind
}

type BoundUnaryOperatorKind string

const (
    Identity BoundUnaryOperatorKind = "Indentity"
    Negation BoundUnaryOperatorKind = "Negation"
)

func NewBoundUnaryExpression(operatorKind BoundUnaryOperatorKind, operand BoundExpression) *BoundUnaryExpression {
    return &BoundUnaryExpression{
        Operand: operand,
        OperatorKind: operatorKind,
    }
}

func (b *BoundUnaryExpression) GetTypeCarrier() TypeCarrier {
    return b.Operand.GetTypeCarrier()
}

func (b *BoundUnaryExpression) Kind() BoundNodeKind {
    return UnaryExpression
}

type BoundLiteralExpression struct {
    Value interface{}
}

func NewBoundLiteralExpression(value interface{}) *BoundLiteralExpression {
    return &BoundLiteralExpression{
        Value: value,
    }
}

func (b *BoundLiteralExpression) Kind() BoundNodeKind {
    return LiteralExpression
}

func (b *BoundLiteralExpression) GetTypeCarrier() TypeCarrier {
    return TypeCarrier(b.Value)
}

type BoundBinaryExpression struct {
    Left BoundExpression
    OperatorKind BoundBinaryOperatorKind
    Right BoundExpression
}

func NewBoundBinaryExpression(left BoundExpression, operatorKind BoundBinaryOperatorKind, right BoundExpression) *BoundBinaryExpression {
    return &BoundBinaryExpression{
        Left: left,
        OperatorKind: operatorKind,
        Right: right,
    }
}

func (b *BoundBinaryExpression) Kind() BoundNodeKind {
   return BinaryExpression 
}

func (b *BoundBinaryExpression) GetTypeCarrier() TypeCarrier {
    return b.Left.GetTypeCarrier()
}

type BoundBinaryOperatorKind string

const (
    Addition BoundBinaryOperatorKind = "Addition"
    Subtraction BoundBinaryOperatorKind = "Subtraction"
    Multiplication BoundBinaryOperatorKind = "Multiplication"
    Division BoundBinaryOperatorKind = "Division"
)

type Binder struct {
    Diagnostics []string
}

func NewBinder() *Binder {
    return &Binder{}
}

func (b *Binder) AddDiagnostic(format string, args ...interface{}) {
    b.Diagnostics = append(b.Diagnostics, fmt.Sprintf(format, args...))
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
        b.AddDiagnostic("Unary operator '%+v' is not defined for type %T", unarySyntax.OperatorNode, boundOperand.GetTypeCarrier()) //todo look for access to runes
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
        b.AddDiagnostic("Binary operator '%+v' is not defined for types %T and %T", binarySyntax.OperatorNode, boundLeft.GetTypeCarrier(), boundRight.GetTypeCarrier()) //todo find access to runes
        return boundLeft;
    }

    return NewBoundBinaryExpression(boundLeft, boundOperatorKind, boundRight)
}

func (b *Binder) BindUnaryOperatorKind(kind SyntaxKind.SyntaxKind, typeCarrier TypeCarrier) BoundUnaryOperatorKind {
    if isInt(typeCarrier) == false {
        return ""
    } 

    switch kind {
    case SyntaxKind.PlusToken:
        return Identity
    case SyntaxKind.MinusToken:
        return Negation
    default:
        panic(fmt.Sprintf("Unexpected unary operator %s", kind))
    }
}

func (b *Binder) BindBinaryOperatorKind(kind SyntaxKind.SyntaxKind, leftTypeCarrier, rightTypeCarrier TypeCarrier) BoundBinaryOperatorKind {
    if isInt(leftTypeCarrier) == false || isInt(rightTypeCarrier) == false {
        return ""
    } 

    switch kind {
    case SyntaxKind.PlusToken:
        return Addition
    case SyntaxKind.MinusToken:
        return Subtraction
    case SyntaxKind.StarToken:
        return Multiplication
    case SyntaxKind.SlashToken:
        return Division
    default:
        panic(fmt.Sprintf("Unexpected binary operator %s", kind))
    }
}

func isInt(val interface{}) bool {
    switch val.(type) {
    case int, int32, int64:
        return true
    default:
        return false
    }
}
