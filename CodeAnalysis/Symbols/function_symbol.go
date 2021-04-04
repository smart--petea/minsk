package Symbols

import (
    "minsk/CodeAnalysis/Symbols/SymbolKind"
)

type FunctionSymbol struct {
    *Symbol

    Parameter []*ParameterSymbol
    Type *TypeSymbol
}

func newFunctionSymbol(name string, parameter []*ParameterSymbol, ttype *TypeSymbol) *FunctionSymbol {
    var functionSymbol functionSymbol

    functionSymbol.Symbol = NewSymbol(name)
    functionSymbol.Parameter = Parameter
    functionSymbol.Type = ttype

    return &functionSymbol
}

func (t *FunctionSymbol) Kind() SymbolKind.SymbolKind {
    return SymbolKind.Function
}
