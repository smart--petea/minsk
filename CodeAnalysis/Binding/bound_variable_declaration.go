package Binding

import (
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
    "minsk/Util"
)

type BoundVariableDeclaration struct {
    Variable *Util.VariableSymbol
    Initializer BoundExpression
}

func NewBoundVariableDeclaration(variable *Util.VariableSymbol, initializer BoundExpression) *BoundVariableDeclaration {
    return &BoundVariableDeclaration{
        Variable: variable,
        Initializer: initializer,
    }
}

func (b *BoundVariableDeclaration) Kind() BoundNodeKind.BoundNodeKind {
    return BoundNodeKind.VariableDeclaration
}
