package BoundNodeKind

type BoundNodeKind string

const (
    UnaryExpression BoundNodeKind = "UnaryExpression"
    LiteralExpression BoundNodeKind = "LiteralExpression"
    BinaryExpression  BoundNodeKind = "BinaryExpression"
    VariableExpression  BoundNodeKind = "VariableExpression"
    AssignmentExpression BoundNodeKind = "AssignmentExpression"
)
