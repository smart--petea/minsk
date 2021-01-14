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

type Binder struct {
    Diagnostics []string
}

func (b *Binder) AddDiagnostic(format string, args ...interface{}) {
    b.Diagnostics = append(b.Diagnostics, fmt.Sprintf(format, args...))
}

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

    if boundOperatorKind == nil {
        p.AddDiagnostic("Unary operator '%s' is not defined for type %v", string(syntax.OperatorToken.Runes), boundOperand.Type)
        return boundOperand;
    }

    return NewBoundUnaryExpression(boundOperatorKind.Value(), boundOperand)
}

func (b *Binder) BindBinaryExpression(syntax BinaryExpressionSyntax) BoundExpression {
    boundLeft := b.BindExpression(syntax.Left)
    boundRight := b.BindExpression(syntax.Right)
    boundOperatorKind := b.BinaryUnaryOperatorKind(syntax.OperatorToken.Kind, boundLeft.Type(), boundRight.Type()) 

    if boundOperatorKind == nil {
        p.AddDiagnostic("Binary operator '%s' is not defined for types %v and %v", string(syntax.OperatorToken.Runes), boundLeft.Type, boundRight.Type)
        return boundLeft;
    }

    return NewBoundBinaryExpression(boundLeft, boundOperatorKind.Value(), boundRight)
}

func (b *Binder) BindUnaryOperatorKind(kind SyntaxKind, operandType Type) BoundUnaryOperatorKind {

    //todo if not type of int return nil

    switch kind {
    case SyntaxKind.PlusToken:
        return Identity
    case SyntaxKind.MinusToken:
        return Negation
    default:
        panic(fmt.Sprintf("Unexpected unary operator %s", kind))
    }
}

func (b *Binder) BindBinaryOperatorKind(kind SyntaxKind, leftType, rightType Type) BoundBinaryOperatorKind {
    //if leftType is not int or rightType is not nil return nil

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
