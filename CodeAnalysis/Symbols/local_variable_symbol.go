package Symbols

import (
    "minsk/CodeAnalysis/Symbols/SymbolKind"
)

type LocalVariableSymbol struct {
    *VariableSymbol
}

func NewLocalVariableSymbol(name string, isReadOnly bool, kind *TypeSymbol) *LocalVariableSymbol {
    var v VariableSymbol
    v.VariableSymbol = NewVariableSymbol(name, isReadOnly, kind)

    return &v
}

func (v *LocalVariableSymbol) Kind() SymbolKind.SymbolKind {
    return SymbolKind.LocalVariable
}
