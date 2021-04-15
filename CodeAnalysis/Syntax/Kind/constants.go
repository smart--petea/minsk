package Kind

const (
    //tokens
    NumberToken SyntaxKind = "NumberToken"
    WhitespaceToken SyntaxKind = "WhitespaceToken"
    StringToken SyntaxKind = "StringToken"
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
    CommaToken SyntaxKind = "CommaToken"
    ColonToken SyntaxKind = "ColonToken"
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
    TildeToken SyntaxKind = "TildeToken"
    HatToken SyntaxKind = "HatToken"
    AmpersandToken SyntaxKind = "AmpersandToken"
    PipeToken SyntaxKind = "PipeToken"

    //Keywords
    LetKeyword SyntaxKind = "LetKeyword"
    ToKeyword SyntaxKind = "ToKeyword"
    FalseKeyword SyntaxKind = "FalseKeyword"
    IfKeyword SyntaxKind = "IfKeyword"
    ElseKeyword SyntaxKind = "ElseKeyword"
    TrueKeyword SyntaxKind = "TrueKeyword"
    VarKeyword SyntaxKind = "VarKeyword"
    WhileKeyword SyntaxKind = "WhileKeyword"
    ForKeyword SyntaxKind = "ForKeyword"
    FunctionKeyword SyntaxKind = "FunctionKeyword"

    //Nodes
    CompilationUnit SyntaxKind = "CompilationUnit"
    Member SyntaxKind = "Member"
    GlobalStatement SyntaxKind = "GlobalStatement"
    FunctionDeclaration SyntaxKind = "FunctionDeclaration"
    Parameter SyntaxKind = "Parameter"
    TypeClause SyntaxKind = "TypeClause"
    ElseClause SyntaxKind = "ElseClause"

    //Statements
    BlockStatement SyntaxKind = "BlockStatement"
    VariableDeclaration SyntaxKind = "VariableDeclaration"
    ExpressionStatement SyntaxKind = "ExpressionStatement"
    IfStatement SyntaxKind = "IfStatement"
    WhileStatement SyntaxKind = "WhileStatement"
    ForStatement SyntaxKind = "ForStatement"

    //expressions
    AssignmentExpression SyntaxKind = "AssignmentExpression" 
    CallExpression SyntaxKind = "CallExpression" 
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
        StringToken,
        PlusToken,
        EndOfFileToken,
        MinusToken,
        StarToken,
        SlashToken, 
        OpenParenthesisToken,
        CloseParenthesisToken,
        BangToken,
        CommaToken,
        ColonToken,
        BadToken,
        IdentifierToken,
        AmpersandAmpersandToken,
        PipePipeToken,
        EqualsEqualsToken,
        EqualsToken,
        BangEqualsToken,
        TrueKeyword,
        ToKeyword,
        FalseKeyword,
        AssignmentExpression, 
        CallExpression, 
        BinaryExpression, 
        LiteralExpression, 
        NameExpression, 
        ParenthesizedExpression,
        UnaryExpression, 
    }
}
