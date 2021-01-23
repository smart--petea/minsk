package BoundBinaryOperatorKind

type BoundBinaryOperatorKind string

const (
    Addition BoundBinaryOperatorKind = "Addition"
    Subtraction BoundBinaryOperatorKind = "Subtraction"
    Multiplication BoundBinaryOperatorKind = "Multiplication"
    Division BoundBinaryOperatorKind = "Division"
    LogicalAnd BoundBinaryOperatorKind = "LogicalAnd"
    LogicalOr BoundBinaryOperatorKind = "LogicalOr"
    Equals BoundBinaryOperatorKind = "Equals"
    NotEquals BoundBinaryOperatorKind = "NotEquals"
)
