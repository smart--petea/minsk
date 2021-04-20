package Binding

import (
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
    "minsk/CodeAnalysis/Symbols"
    "minsk/Util"
)

type BoundVariableExpression struct {
    *Util.ChildrenProvider

    Variable Symbols.IVariableSymbol
}

func NewBoundVariableExpression(variable Symbols.IVariableSymbol) *BoundVariableExpression {
    return &BoundVariableExpression{
        ChildrenProvider: Util.NewChildrenProvider(),

        Variable: variable,
    }
}

func (b *BoundVariableExpression) GetType() *Symbols.TypeSymbol {
    return b.Variable.GetType()
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
            Value: b.Variable.GetName(),
        },
    }
}
