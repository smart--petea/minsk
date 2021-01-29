package Binding

import (
    "reflect"
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
)

type BoundVariableExpression struct {
    Name string
    Type reflect.Kind
}

func NewBoundVariableExpression(name string, variableType reflect.Kind) *BoundVariableExpression {
    return &BoundVariableExpression{
        Name: name,
        Type: variableType,
    }
}

func (b *BoundVariableExpression) GetType() reflect.Kind {
    return b.Type
}

func (b *BoundVariableExpression) Kind() BoundNodeKind.BoundNodeKind {
    return BoundNodeKind.VariableExpression
}
