package Binding

import (
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
    "minsk/CodeAnalysis/Symbols"
    "minsk/Util"
    "reflect"
)

type BoundVariableExpression struct {
    *Util.ChildrenProvider

    Variable *Symbols.VariableSymbol
}

func NewBoundVariableExpression(variable *Symbols.VariableSymbol) *BoundVariableExpression {
    return &BoundVariableExpression{
        ChildrenProvider: Util.NewChildrenProvider(),

        Variable: variable,
    }
}

func (b *BoundVariableExpression) GetType() reflect.Kind {
    return b.Variable.Type
}

func (b *BoundVariableExpression) Kind() BoundNodeKind.BoundNodeKind {
    return BoundNodeKind.VariableExpression
}

func (b *BoundVariableExpression) GetProperties() []*BoundNodeProperty {
    return []*BoundNodeProperty{
        {
            Name: "type",
            Value: b.GetType(),
        },
        {
            Name: "variable",
            Value: b.Variable.Name,
        },
    }
}
