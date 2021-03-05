package Binding

import (
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
    "minsk/Util"
)

type BoundExpressionStatement struct {
    *Util.ChildrenProvider

    Expression BoundExpression
}

func NewBoundExpressionStatement(expression BoundExpression) *BoundExpressionStatement {
    return &BoundExpressionStatement{
        ChildrenProvider: Util.NewChildrenProvider(expression),

        Expression: expression,
    }
}

func (b *BoundExpressionStatement) Kind() BoundNodeKind.BoundNodeKind {
    return BoundNodeKind.ExpressionStatement
}

//todo
func (b *BoundExpressionStatement) GetProperties() []*BoundNodeProperty {
    return []*BoundNodeProperty{}
}
