package Symbols

import (
    "minsk/CodeAnalysis/Symbols/SymbolKind"
)

type GlobalVariableSymbol struct {
    *VariableSymbol
}

func NewGlobalVariableSymbol(name string, isReadOnly bool, kind *TypeSymbol) *GlobalVariableSymbol {
    var g GlobalVariableSymbol
    g.VariableSymbol = NewVariableSymbol(name, isReadOnly, kind)

    return &g
}

func (g *GlobalVariableSymbol) Kind() SymbolKind.SymbolKind {
    return SymbolKind.GlobalVariable
}
