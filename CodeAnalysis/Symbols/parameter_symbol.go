package Symbols

import (
    "minsk/CodeAnalysis/Symbols/SymbolKind"
)

type ParameterSymbol struct {
    *LocalVariableSymbol
}

func NewParameterSymbol(name string, ttype *TypeSymbol) *ParameterSymbol {
    var parameterSymbol ParameterSymbol

    isReadOnly := true
    parameterSymbol.LocalVariableSymbol = NewLocalVariableSymbol(name, isReadOnly, ttype)

    return &parameterSymbol
}

func (t *ParameterSymbol) Kind() SymbolKind.SymbolKind {
    return SymbolKind.Parameter
}
