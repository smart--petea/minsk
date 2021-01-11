package SyntaxKind

type SyntaxKind string

const (
    //tokens
    NumberToken SyntaxKind = "NumberToken"
    WhitespaceToken SyntaxKind = "WhitespaceToken"
    PlusToken SyntaxKind = "PlusToken"
    EndOfFileToken SyntaxKind = "EndOfFileToken"
    MinusToken SyntaxKind = "MinusToken"
    StarToken SyntaxKind = "StarToken"
    SlashToken SyntaxKind = "SlashToken"
    OpenParenthesisToken SyntaxKind = "OpenParenthisToken"
    CloseParenthesisToken SyntaxKind = "CloseParenthisToken"
    BadToken SyntaxKind = "BadToken"

    //expressions
    BinaryExpression SyntaxKind = "BinaryExpression" 
    UnaryExpression SyntaxKind = "UnaryExpression" 
    LiteralExpression SyntaxKind = "LiteralExpression" 
    ParenthesizedExpression SyntaxKind = "ParenthesizedExpression"
)

func (kind SyntaxKind) GetUnaryOperatorPrecedence() int {
    switch kind {
        case PlusToken: 
            return 3
        case MinusToken: 
            return 3

        default:
            return 0
    }
}

func (kind SyntaxKind) GetBinaryOperatorPrecedence() int {
    switch kind {
        case StarToken:
            return 2
        case SlashToken:
            return 2

        case PlusToken: 
            return 1
        case MinusToken: 
            return 1

        default:
            return 0
    }
}
