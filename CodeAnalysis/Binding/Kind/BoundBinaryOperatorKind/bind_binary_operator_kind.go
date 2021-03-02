package BoundBinaryOperatorKind

type BoundBinaryOperatorKind string

const (
    Addition BoundBinaryOperatorKind = "Addition"
    Subtraction BoundBinaryOperatorKind = "Subtraction"
    Multiplication BoundBinaryOperatorKind = "Multiplication"
    Division BoundBinaryOperatorKind = "Division"
    BitwiseAnd BoundBinaryOperatorKind = "BitwiseAnd"
    BitwiseXor BoundBinaryOperatorKind = "BitwiseXor"
    BitwiseOr BoundBinaryOperatorKind = "BitwiseOr"
    LogicalAnd BoundBinaryOperatorKind = "LogicalAnd"
    LogicalOr BoundBinaryOperatorKind = "LogicalOr"
    Equals BoundBinaryOperatorKind = "Equals"
    NotEquals BoundBinaryOperatorKind = "NotEquals"
    Less BoundBinaryOperatorKind = "Less"
    LessOrEquals BoundBinaryOperatorKind = "LessOrEquals"
    Greater BoundBinaryOperatorKind = "Greater"
    GreaterOrEquals BoundBinaryOperatorKind = "GreaterOrEquals"
)
