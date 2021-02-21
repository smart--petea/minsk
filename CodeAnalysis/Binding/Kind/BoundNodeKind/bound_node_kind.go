package BoundNodeKind

type BoundNodeKind string

const (
    //Statements
    BlockStatement BoundNodeKind = "BlockStatement"
    ExpressionStatement BoundNodeKind = "ExpressionStatement"
    VariableDeclaration BoundNodeKind = "VariableDeclaration"

    //Expressions
    UnaryExpression BoundNodeKind = "UnaryExpression"
    LiteralExpression BoundNodeKind = "LiteralExpression"
    BinaryExpression  BoundNodeKind = "BinaryExpression"
    VariableExpression  BoundNodeKind = "VariableExpression"
    AssignmentExpression BoundNodeKind = "AssignmentExpression"
)
