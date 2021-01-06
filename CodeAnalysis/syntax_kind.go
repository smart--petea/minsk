package CodeAnalysis

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
    LiteralExpression SyntaxKind = "LiteralExpression" 
    ParenthesizedExpression SyntaxKind = "ParenthesizedExpression"
)

