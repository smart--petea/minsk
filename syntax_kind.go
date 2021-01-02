package minsk

type SyntaxKind string

const (
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
    BinaryExpression SyntaxKind = "BinaryExpression" 
    NumberExpression SyntaxKind = "NumberExpression" 
    ParenthesizedExpression SyntaxKind = "ParenthesizedExpression"
)

