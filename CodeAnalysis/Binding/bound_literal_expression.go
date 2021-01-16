package Binding

import (
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
    "minsk/CodeAnalysis/Binding/TypeCarrier"
)

type BoundLiteralExpression struct {
    Value interface{}
}

func NewBoundLiteralExpression(value interface{}) *BoundLiteralExpression {
    return &BoundLiteralExpression{
        Value: value,
    }
}

func (b *BoundLiteralExpression) Kind() BoundNodeKind.BoundNodeKind {
    return BoundNodeKind.LiteralExpression
}

func (b *BoundLiteralExpression) GetTypeCarrier() TypeCarrier.TypeCarrier {
    return TypeCarrier.TypeCarrier(b.Value)
}
