package Binding

import (
    "minsk/CodeAnalysis/Binding/BoundBinaryOperator"
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
    "minsk/CodeAnalysis/Binding/TypeCarrier"
)

type BoundBinaryExpression struct {
    Left BoundExpression
    Op *BoundBinaryOperator.BoundBinaryOperator
    Right BoundExpression
}

func NewBoundBinaryExpression(left BoundExpression, op *BoundBinaryOperator.BoundBinaryOperator, right BoundExpression) *BoundBinaryExpression {
    return &BoundBinaryExpression{
        Left: left,
        Op: op,
        Right: right,
    }
}

func (b *BoundBinaryExpression) Kind() BoundNodeKind.BoundNodeKind {
   return BoundNodeKind.BinaryExpression 
}

func (b *BoundBinaryExpression) GetTypeCarrier() TypeCarrier.TypeCarrier {
    return b.Op.ResultTypeCarrier
}
