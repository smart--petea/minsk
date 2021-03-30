package Binding

import (
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
    "minsk/Util"
    "reflect"

    "log"
)

type BoundLiteralExpression struct {
    *Util.ChildrenProvider

    Value interface{}
}

func NewBoundLiteralExpression(value interface{}) *BoundLiteralExpression {
    log.Printf("NewBoundLiteralExpression %+v", value)
    return &BoundLiteralExpression{
        ChildrenProvider: Util.NewChildrenProvider(),

        Value: value,
    }
}

func (b *BoundLiteralExpression) Kind() BoundNodeKind.BoundNodeKind {
    return BoundNodeKind.LiteralExpression
}

func (b *BoundLiteralExpression) GetType() reflect.Kind {
    return reflect.TypeOf(b.Value).Kind()
}

func (b *BoundLiteralExpression) GetProperties() []*BoundNodeProperty {
    return []*BoundNodeProperty{
        {
            Name: "type",
            Value: b.GetType(),
        },
        {
            Name: "value",
            Value: b.Value,
        },
    }
}
