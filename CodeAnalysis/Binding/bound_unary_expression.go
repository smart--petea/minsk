package Binding

import (
    "minsk/CodeAnalysis/Binding/BoundUnaryOperator"
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
    "minsk/CodeAnalysis/Binding/TypeCarrier"
)

type BoundUnaryExpression struct {
    Operand BoundExpression
    Op *BoundUnaryOperator.BoundUnaryOperator
}

func NewBoundUnaryExpression(op *BoundUnaryOperator.BoundUnaryOperator, operand BoundExpression) *BoundUnaryExpression {
    return &BoundUnaryExpression{
        Operand: operand,
        Op: op,
    }
}

func (b *BoundUnaryExpression) GetTypeCarrier() TypeCarrier.TypeCarrier {
    return b.Operand.GetTypeCarrier()
}

func (b *BoundUnaryExpression) Kind() BoundNodeKind.BoundNodeKind {
    return BoundNodeKind.UnaryExpression
}
