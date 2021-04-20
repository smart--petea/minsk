package Binding

import (
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
    "minsk/CodeAnalysis/Symbols"
    "minsk/Util"
)

type BoundVariableDeclaration struct {
    *Util.ChildrenProvider

    Variable Symbols.IVariableSymbol
    Initializer BoundExpression
}

func NewBoundVariableDeclaration(variable Symbols.IVariableSymbol, initializer BoundExpression) *BoundVariableDeclaration {
    return &BoundVariableDeclaration{
        ChildrenProvider: Util.NewChildrenProvider(initializer),

        Variable: variable,
        Initializer: initializer,
    }
}

func (b *BoundVariableDeclaration) Kind() BoundNodeKind.BoundNodeKind {
    return BoundNodeKind.VariableDeclaration
}

func (b *BoundVariableDeclaration) GetProperties() []*BoundNodeProperty {
    return []*BoundNodeProperty{
        {
            Name: "variable",
            Value: b.Variable.GetName(),
        },
    }
}
