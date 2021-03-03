package Binding

import (
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
    "reflect"
)

type BoundLiteralExpression struct {
    *boundNodeChildren

    Value interface{}
}

func NewBoundLiteralExpression(value interface{}) *BoundLiteralExpression {
    return &BoundLiteralExpression{
        boundNodeChildren: newBoundNodeChildren(),

        Value: value,
    }
}

func (b *BoundLiteralExpression) Kind() BoundNodeKind.BoundNodeKind {
    return BoundNodeKind.LiteralExpression
}

func (b *BoundLiteralExpression) GetType() reflect.Kind {
    return reflect.TypeOf(b.Value).Kind()
}
