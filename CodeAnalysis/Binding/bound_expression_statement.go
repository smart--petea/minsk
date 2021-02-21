package Binding

import (
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
)

type BoundExpressionStatement struct {
    Expression BoundExpression
}

func NewBoundExpressionStatement(expression BoundExpression) *BoundExpressionStatement {
    return &BoundExpressionStatement{
        Expression: expression,
    }
}

func (b *BoundExpressionStatement) Kind() BoundNodeKind.BoundNodeKind {
    return BoundNodeKind.ExpressionStatement
}
