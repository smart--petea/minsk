package Binding

import (
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
    "minsk/Util"
    "reflect"
)

type BoundVariableExpression struct {
    Variable *Util.VariableSymbol
}

func NewBoundVariableExpression(variable *Util.VariableSymbol) *BoundVariableExpression {
    return &BoundVariableExpression{
        Variable: variable,
    }
}

func (b *BoundVariableExpression) GetType() reflect.Kind {
    return b.Variable.Type
}

func (b *BoundVariableExpression) Kind() BoundNodeKind.BoundNodeKind {
    return BoundNodeKind.VariableExpression
}
