package Symbols

import (
    "minsk/CodeAnalysis/Symbols/SymbolKind"
)

type VariableSymbol struct {
    *Symbol

    ttype *TypeSymbol
    isReadOnly bool
}

func NewVariableSymbol(name string, isReadOnly bool, ttype *TypeSymbol) *VariableSymbol {
    var v VariableSymbol
    v.Symbol = NewSymbol(name)
    v.ttype = ttype
    v.isReadOnly = isReadOnly

    return &v
}

func (v *VariableSymbol) Kind() SymbolKind.SymbolKind {
    return SymbolKind.Variable
}

func (v *VariableSymbol) GetType() *TypeSymbol {
    return v.ttype
}

func (v *VariableSymbol) IsReadOnly() bool {
    return v.isReadOnly
}
