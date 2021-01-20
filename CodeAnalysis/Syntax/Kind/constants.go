package Kind

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
    BangToken SyntaxKind = "BangToken"
    BadToken SyntaxKind = "BadToken"
    IdentifierToken SyntaxKind = "IdentifierToken"
    AmpersandAmpersandToken SyntaxKind = "AmpersandAmpersandToken"
    PipePipeToken SyntaxKind = "PipePipeToken"

    //Keywords
    TrueKeyword SyntaxKind = "TrueKeyword"
    FalseKeyword SyntaxKind = "FalseKeyword"

    //expressions
    BinaryExpression SyntaxKind = "BinaryExpression" 
    UnaryExpression SyntaxKind = "UnaryExpression" 
    LiteralExpression SyntaxKind = "LiteralExpression" 
    ParenthesizedExpression SyntaxKind = "ParenthesizedExpression"
)
