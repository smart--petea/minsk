package Binding

import (
    "minsk/CodeAnalysis/Binding/BoundBinaryOperator"
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
    "minsk/Util"

    "reflect"
)

type BoundBinaryExpression struct {
    *Util.ChildrenProvider

    Left BoundExpression
    Op *BoundBinaryOperator.BoundBinaryOperator
    Right BoundExpression
}

func NewBoundBinaryExpression(left BoundExpression, op *BoundBinaryOperator.BoundBinaryOperator, right BoundExpression) *BoundBinaryExpression {
    return &BoundBinaryExpression{
        ChildrenProvider: Util.NewChildrenProvider(left, right),

        Left: left,
        Op: op,
        Right: right,
    }
}

func (b *BoundBinaryExpression) Kind() BoundNodeKind.BoundNodeKind {
   return BoundNodeKind.BinaryExpression 
}

func (b *BoundBinaryExpression) GetType() reflect.Kind {
    return b.Op.ResultType
}
