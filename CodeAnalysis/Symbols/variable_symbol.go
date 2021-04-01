package Symbols

import (
    "minsk/CodeAnalysis/Symbols/SymbolKind"
)

type VariableSymbol struct {
    *Symbol

    Type *TypeSymbol
    IsReadOnly bool
}

func NewVariableSymbol(name string, isReadOnly bool, kind *TypeSymbol) *VariableSymbol {
    var v VariableSymbol
    v.Symbol = NewSymbol(name)
    v.Type = kind
    v.IsReadOnly = isReadOnly

    return &v
}

func (v *VariableSymbol) Kind() SymbolKind.SymbolKind {
    return SymbolKind.Variable
}
