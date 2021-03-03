package Binding

import (
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
    "minsk/Util"
)

type BoundVariableDeclaration struct {
    *boundNodeChildren

    Variable *Util.VariableSymbol
    Initializer BoundExpression
}

func NewBoundVariableDeclaration(variable *Util.VariableSymbol, initializer BoundExpression) *BoundVariableDeclaration {
    return &BoundVariableDeclaration{
        boundNodeChildren: newBoundNodeChildren(initializer),

        Variable: variable,
        Initializer: initializer,
    }
}

func (b *BoundVariableDeclaration) Kind() BoundNodeKind.BoundNodeKind {
    return BoundNodeKind.VariableDeclaration
}
