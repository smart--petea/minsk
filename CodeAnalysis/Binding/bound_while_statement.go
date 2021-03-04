package Binding

import (
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
    "minsk/Util"
)

type BoundWhileStatement struct {
    *Util.ChildrenProvider

    Condition BoundExpression
    Body BoundStatement
}

func NewBoundWhileStatement(condition BoundExpression, body BoundStatement) *BoundWhileStatement {
    return &BoundWhileStatement{
        ChildrenProvider: Util.NewChildrenProvider(condition, body),

        Condition: condition,
        Body: body,
    }
}

func (b *BoundWhileStatement) Kind() BoundNodeKind.BoundNodeKind {
    return BoundNodeKind.WhileStatement
}
