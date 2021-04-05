package Symbols

import (
    "minsk/CodeAnalysis/Symbols/SymbolKind"
)

type ParameterSymbol struct {
    *VariableSymbol
}

func newParameterSymbol(name string, ttype *TypeSymbol) *ParameterSymbol {
    var parameterSymbol ParameterSymbol

    isReadOnly := true
    parameterSymbol.VariableSymbol = NewVariableSymbol(name, isReadOnly, ttype)

    return &parameterSymbol
}

func (t *ParameterSymbol) Kind() SymbolKind.SymbolKind {
    return SymbolKind.Parameter
}
