package BoundUnaryOperatorKind

type BoundUnaryOperatorKind string

const (
    Identity BoundUnaryOperatorKind = "Indentity"
    Negation BoundUnaryOperatorKind = "Negation"
    OnesComplement BoundUnaryOperatorKind = "OnesComplement"
    LogicalNegation BoundUnaryOperatorKind = "LogicalNegation"
)
