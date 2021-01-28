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
    EqualsEqualsToken SyntaxKind = "EqualsEqualsToken"
    EqualsToken SyntaxKind = "EqualsEqualsToken"
    BangEqualsToken SyntaxKind = "BangEqualsToken"

    //Keywords
    TrueKeyword SyntaxKind = "TrueKeyword"
    FalseKeyword SyntaxKind = "FalseKeyword"

    //expressions
    AssignmentExpression SyntaxKind = "AssignmentExpression" 
    BinaryExpression SyntaxKind = "BinaryExpression" 
    LiteralExpression SyntaxKind = "LiteralExpression" 
    NameExpression SyntaxKind = "NameExpression" 
    ParenthesizedExpression SyntaxKind = "ParenthesizedExpression"
    UnaryExpression SyntaxKind = "UnaryExpression" 
)
