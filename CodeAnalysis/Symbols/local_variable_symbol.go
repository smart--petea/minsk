package Symbols

import (
    "minsk/CodeAnalysis/Symbols/SymbolKind"
)

type LocalVariableSymbol struct {
    *VariableSymbol
}

func NewLocalVariableSymbol(name string, isReadOnly bool, kind *TypeSymbol) *LocalVariableSymbol {
    var l LocalVariableSymbol
    l.VariableSymbol = NewVariableSymbol(name, isReadOnly, kind)

    return &l
}

func (l *LocalVariableSymbol) Kind() SymbolKind.SymbolKind {
    return SymbolKind.LocalVariable
}
