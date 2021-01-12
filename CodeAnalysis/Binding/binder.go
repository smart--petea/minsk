package Binding

type BoundNodeKind struct

const (
    UnaryExpression BoundNodeKind = "UnaryExpression"
    LiteralExpression BoundNodeKind = "LiteralExpression"
    BinaryExpression  BoundNodeKind = "BinaryExpression"
)

type BoundNode interface {
    Kind() BoundNodeKind
}

type BoundExpression interface {
    BoundNode

    Type() Type
}

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

func (b *BoundUnaryExpression) Type() Type {
    return b.Operand.Type
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

func (b *BoundLiteralExpression) Type() Type {
    return b.Value.Type()
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

func (b *BoundBinaryExpression) Type() Type {
    return b.Left.Type()
}

type BoundBinaryOperatorKind string

const (
    Addition BoundBinaryOperatorKind = "Addition"
    Subtraction BoundBinaryOperatorKind = "Subtraction"
    Multiplication BoundBinaryOperatorKind = "Multiplication"
    Division BoundBinaryOperatorKind = "Division"
)

type Binder struct {}

func (b *Binder) BindExpression(syntax ExpressionSyntax) BoundExpression {
    switch syntax.Kind() {
    case SyntaxKind.UnaryExpression:
        return b.BindUnaryExpression(UnaryExpressionSyntax(syntax))
    case SyntaxKind.BinaryExpression:
        return b.BindBinaryExpression(BinaryExpressionSyntax(syntax))
    case SyntaxKind.LiteralExpression:
        return b.BindLiteralExpression(LiteralExpressionSyntax(syntax))
    default:
        panic(fmt.Sprintf("Unexpected syntax %s", syntax.Kind()))
    }
}

func (b *Binder) BindLiteralExpression(syntax LiteralExpressionSyntax) BoundExpression {
    var value int

    if val, ok := syntax.LiteralToken.Value.(int); ok {
        value = val
    }


    return NewBoundLiteralExpression(value)
}

func (b *Binder) BindUnaryExpression(syntax UnaryExpressionSyntax) BoundExpression {
    boundOperand := b.BindExpression(syntax.Operand)
    boundOperatorKind := b.BinaryUnaryOperatorKind(syntax.OperatorToken.Kind, boundOperator.Type()) 

    return NewBoundUnaryExpression(boundOperatorKind, boundOperand)
}

func (b *Binder) BindBinaryExpression(syntax BinaryExpressionSyntax) BoundExpression {
    boundLeft := b.BindExpression(syntax.Left)
    boundOperatorKind := b.BinaryUnaryOperatorKind(syntax.OperatorToken.Kind) 
    boundRight := b.BindExpression(syntax.Right)

    return NewBoundBinaryExpression(boundLeft, boundOperatorKind, boundRight)
}

func (b *Binder) BindUnaryOperatorKind(kind SyntaxKind, operandType Type) BoundUnaryOperatorKind {
    //todo if not type of int return nil
    if operandType 
    switch kind {
    case SyntaxKind.PlusToken:
        return Identity
    case SyntaxKind.MinusToken:
        return Negation
    default:
        panic(fmt.Sprintf("Unexpected unary operator %s", kind))
    }
}

func (b *Binder) BindBinaryOperatorKind(kind SyntaxKind) BoundBinaryOperatorKind {
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
