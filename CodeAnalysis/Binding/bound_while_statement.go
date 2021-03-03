package Binding

import (
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
)

type BoundWhileStatement struct {
    *boundNodeChildren

    Condition BoundExpression
    Body BoundStatement
}

func NewBoundWhileStatement(condition BoundExpression, body BoundStatement) *BoundWhileStatement {
    return &BoundWhileStatement{
        boundNodeChildren: newBoundNodeChildren(condition, body),

        Condition: condition,
        Body: body,
    }
}

func (b *BoundWhileStatement) Kind() BoundNodeKind.BoundNodeKind {
    return BoundNodeKind.WhileStatement
}
