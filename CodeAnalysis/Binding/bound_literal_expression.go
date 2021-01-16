package Binding

import (
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
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

func (b *BoundLiteralExpression) GetTypeCarrier() TypeCarrier {
    return TypeCarrier(b.Value)
}
