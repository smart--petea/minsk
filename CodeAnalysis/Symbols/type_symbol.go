package Symbols

import (
    "minsk/CodeAnalysis/Symbols/SymbolKind"
)

type TypeSymbol struct {
    *Symbol
}

func newTypeSymbol(name string) *TypeSymbol {
    var typeSymbol TypeSymbol

    typeSymbol.Symbol = NewSymbol(name)
    return &typeSymbol
}

func (t *TypeSymbol) Kind() SymbolKind.SymbolKind {
    return SymbolKind.Type
}

var (
    TypeSymbolError *TypeSymbol = newTypeSymbol("?")
    TypeSymbolInt *TypeSymbol = newTypeSymbol("int")
    TypeSymbolBool *TypeSymbol = newTypeSymbol("bool")
    TypeSymbolString *TypeSymbol = newTypeSymbol("string")
    TypeSymbolVoid *TypeSymbol = newTypeSymbol("void")
)
