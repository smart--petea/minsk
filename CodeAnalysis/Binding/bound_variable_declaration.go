package Binding

import (
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
)

type BoundVariableDeclaration struct {
    Variable VariableSymbol
    Initializer BoundExpression
}

func NewBoundVariableDeclaration(variable VariableSymbol, initializer BoundExpression) *BoundVariableDeclaration {
    return &BoundVariableDeclaration{
        Variable: variable,
        Initializer: initializer,
    }
}

func (b *BoundVariableDeclaration) Kind() BoundNodeKind.BoundNodeKind {
    return BoundNodeKind.VariableDeclaration
}
