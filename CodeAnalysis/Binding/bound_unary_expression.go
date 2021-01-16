package Binding

import (
    "minsk/CodeAnalysis/Binding/Kind/BoundUnaryOperatorKind"
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
)

type BoundUnaryExpression struct {
    Operand BoundExpression
    OperatorKind BoundUnaryOperatorKind.BoundUnaryOperatorKind
}

func NewBoundUnaryExpression(operatorKind BoundUnaryOperatorKind.BoundUnaryOperatorKind, operand BoundExpression) *BoundUnaryExpression {
    return &BoundUnaryExpression{
        Operand: operand,
        OperatorKind: operatorKind,
    }
}

func (b *BoundUnaryExpression) GetTypeCarrier() TypeCarrier {
    return b.Operand.GetTypeCarrier()
}

func (b *BoundUnaryExpression) Kind() BoundNodeKind.BoundNodeKind {
    return BoundNodeKind.UnaryExpression
}
