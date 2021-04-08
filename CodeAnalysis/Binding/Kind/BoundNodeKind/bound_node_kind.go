package BoundNodeKind

type BoundNodeKind string

const (
    //Statements
    BlockStatement BoundNodeKind = "BlockStatement"
    ConditionalGotoStatement BoundNodeKind = "ConditionalGotoStatement"
    ExpressionStatement BoundNodeKind = "ExpressionStatement"
    IfStatement BoundNodeKind = "IfStatement"
    WhileStatement BoundNodeKind = "WhileStatement"
    ForStatement BoundNodeKind = "ForStatement"
    GotoStatement BoundNodeKind = "GotoStatement"
    LabelStatement BoundNodeKind = "LabelStatement"
    VariableDeclaration BoundNodeKind = "VariableDeclaration"

    //Expressions
    UnaryExpression BoundNodeKind = "UnaryExpression"
    LiteralExpression BoundNodeKind = "LiteralExpression"
    ConversionExpression BoundNodeKind = "ConversionExpression"
    BinaryExpression  BoundNodeKind = "BinaryExpression"
    VariableExpression  BoundNodeKind = "VariableExpression"
    AssignmentExpression BoundNodeKind = "AssignmentExpression"
    ErrorExpression BoundNodeKind = "ErrorExpression"
    CallExpression BoundNodeKind = "CallExpression"
)
