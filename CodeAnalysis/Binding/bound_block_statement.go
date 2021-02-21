package Binding

import (
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
)

type BoundBlockStatement struct {
    Statements []BoundStatement
}

func NewBoundBlockStatement(statements []BoundStatement) *BoundBlockStatement {
    return &BoundBlockStatement{
        Statements: statements,
    }
}

func (b *BoundBlockStatement) Kind() BoundNodeKind.BoundNodeKind {
    return BoundNodeKind.BlockStatement
}
