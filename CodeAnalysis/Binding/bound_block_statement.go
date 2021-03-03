package Binding

import (
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
)

type BoundBlockStatement struct {
    *boundNodeChildren

    Statements []BoundStatement
}

func NewBoundBlockStatement(statements []BoundStatement) *BoundBlockStatement {
    var boundNodes []BoundNode
    for _, statement := range statements {
        boundNodes = append(boundNodes, BoundNode(statement))
    }

    return &BoundBlockStatement{
        boundNodeChildren: newBoundNodeChildren(boundNodes...),

        Statements: statements,
    }
}

func (b *BoundBlockStatement) Kind() BoundNodeKind.BoundNodeKind {
    return BoundNodeKind.BlockStatement
}
