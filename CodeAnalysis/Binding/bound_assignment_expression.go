package Binding

import (
    "reflect"
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
)

type BoundAssignmentExpression struct {
    Name string
    Expression BoundExpression
}

func NewBoundAssignmentExpression(name string, expression BoundExpression) *BoundAssignmentExpression {
    return &BoundAssignmentExpression{
        Name: name,
        Expression: expression,
    }
}

func (b *BoundAssignmentExpression) GetType() reflect.Kind {
    return b.Expression.GetType()
}

func (b *BoundAssignmentExpression) Kind() BoundNodeKind.BoundNodeKind {
    return BoundNodeKind.AssignmentExpression
}
