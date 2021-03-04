package Binding

import (
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
    "minsk/Util"
)

type BoundBlockStatement struct {
    *Util.ChildrenProvider

    Statements []BoundStatement
}

func NewBoundBlockStatement(statements []BoundStatement) *BoundBlockStatement {
    var boundNodes []interface{}
    for _, statement := range statements {
        boundNodes = append(boundNodes, statement)
    }

    return &BoundBlockStatement{
        ChildrenProvider: Util.NewChildrenProvider(boundNodes...),

        Statements: statements,
    }
}

func (b *BoundBlockStatement) Kind() BoundNodeKind.BoundNodeKind {
    return BoundNodeKind.BlockStatement
}
