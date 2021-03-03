package Binding

import (
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
)

type BoundExpressionStatement struct {
    *boundNodeChildren

    Expression BoundExpression
}

func NewBoundExpressionStatement(expression BoundExpression) *BoundExpressionStatement {
    return &BoundExpressionStatement{
        boundNodeChildren: newBoundNodeChildren(expression),

        Expression: expression,
    }
}

func (b *BoundExpressionStatement) Kind() BoundNodeKind.BoundNodeKind {
    return BoundNodeKind.ExpressionStatement
}
