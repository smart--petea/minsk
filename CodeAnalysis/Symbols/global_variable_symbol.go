package Symbols

import (
    "minsk/CodeAnalysis/Symbols/SymbolKind"
)

type GlobalVariableSymbol struct {
    *VariableSymbol
}

func NewGlobalVariableSymbol(name string, isReadOnly bool, kind *TypeSymbol) *GlobalVariableSymbol {
    var v VariableSymbol
    v.VariableSymbol = NewVariableSymbol(name, isReadOnly, kind)

    return &v
}

func (v *GlobalVariableSymbol) Kind() SymbolKind.SymbolKind {
    return SymbolKind.GlobalVariable
}
