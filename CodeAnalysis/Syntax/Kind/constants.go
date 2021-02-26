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
    OpenBraceToken SyntaxKind = "OpenBraceToken"
    CloseBraceToken SyntaxKind = "CloseBraceToken"
    BangToken SyntaxKind = "BangToken"
    BadToken SyntaxKind = "BadToken"
    IdentifierToken SyntaxKind = "IdentifierToken"
    AmpersandAmpersandToken SyntaxKind = "AmpersandAmpersandToken"
    PipePipeToken SyntaxKind = "PipePipeToken"
    EqualsEqualsToken SyntaxKind = "EqualsEqualsToken"
    LessToken SyntaxKind = "LessToken"
    LessOrEqualsToken SyntaxKind = "LessOrEqualsToken"
    GreaterToken SyntaxKind = "GreaterToken"
    GreaterOrEqualsToken SyntaxKind = "GreaterOrEqualsToken"
    EqualsToken SyntaxKind = "EqualsToken"
    BangEqualsToken SyntaxKind = "BangEqualsToken"

    //Keywords
    LetKeyword SyntaxKind = "LetKeyword"
    FalseKeyword SyntaxKind = "FalseKeyword"
    IfKeyword SyntaxKind = "IfKeyword"
    ElseKeyword SyntaxKind = "ElseKeyword"
    TrueKeyword SyntaxKind = "TrueKeyword"
    VarKeyword SyntaxKind = "VarKeyword"

    //Nodes
    CompilationUnit SyntaxKind = "CompilationUnit"

    //Statements
    BlockStatement SyntaxKind = "BlockStatement"
    VariableDeclaration SyntaxKind = "VariableDeclaration"
    ExpressionStatement SyntaxKind = "ExpressionStatement"
    IfStatement SyntaxKind = "IfStatement"
    ElseClause SyntaxKind = "ElseClause"

    //expressions
    AssignmentExpression SyntaxKind = "AssignmentExpression" 
    BinaryExpression SyntaxKind = "BinaryExpression" 
    LiteralExpression SyntaxKind = "LiteralExpression" 
    NameExpression SyntaxKind = "NameExpression" 
    ParenthesizedExpression SyntaxKind = "ParenthesizedExpression"
    UnaryExpression SyntaxKind = "UnaryExpression" 
)

func GetValues() []SyntaxKind {
    return []SyntaxKind{
        NumberToken,
        WhitespaceToken,
        PlusToken,
        EndOfFileToken,
        MinusToken,
        StarToken,
        SlashToken, 
        OpenParenthesisToken,
        CloseParenthesisToken,
        BangToken,
        BadToken,
        IdentifierToken,
        AmpersandAmpersandToken,
        PipePipeToken,
        EqualsEqualsToken,
        EqualsToken,
        BangEqualsToken,
        TrueKeyword,
        FalseKeyword,
        AssignmentExpression, 
        BinaryExpression, 
        LiteralExpression, 
        NameExpression, 
        ParenthesizedExpression,
        UnaryExpression, 
    }
}
