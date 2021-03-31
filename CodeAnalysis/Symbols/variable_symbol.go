package Symbols

import (
    "minsk/CodeAnalysis/Symbols/SymbolKind"

    "reflect"
)

type VariableSymbol struct {
    *Symbol

    Type reflect.Kind
    IsReadOnly bool
}

func NewVariableSymbol(name string, isReadOnly bool, kind reflect.Kind) *VariableSymbol {
    var v VariableSymbol
    v.Symbol = NewSymbol(name)
    v.Type = kind
    v.IsReadOnly = isReadOnly

    return &v
}

func (v *VariableSymbol) Kind() SymbolKind.SymbolKind {
    return SymbolKind.Variable
}
