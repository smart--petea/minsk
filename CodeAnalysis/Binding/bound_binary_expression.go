package Binding

import (
    "minsk/CodeAnalysis/Binding/Kind/BoundBinaryOperatorKind"
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
)

type BoundBinaryExpression struct {
    Left BoundExpression
    OperatorKind BoundBinaryOperatorKind.BoundBinaryOperatorKind
    Right BoundExpression
}

func NewBoundBinaryExpression(left BoundExpression, operatorKind BoundBinaryOperatorKind.BoundBinaryOperatorKind, right BoundExpression) *BoundBinaryExpression {
    return &BoundBinaryExpression{
        Left: left,
        OperatorKind: operatorKind,
        Right: right,
    }
}

func (b *BoundBinaryExpression) Kind() BoundNodeKind.BoundNodeKind {
   return BoundNodeKind.BinaryExpression 
}

func (b *BoundBinaryExpression) GetTypeCarrier() TypeCarrier {
    return b.Left.GetTypeCarrier()
}
