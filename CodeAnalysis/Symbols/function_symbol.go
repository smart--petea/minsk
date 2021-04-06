package Symbols

import (
    "minsk/CodeAnalysis/Symbols/SymbolKind"
)

type FunctionSymbol struct {
    *Symbol

    Parameter []*ParameterSymbol
    Name string
    Type *TypeSymbol
}

func NewFunctionSymbol(name string, parameter []*ParameterSymbol, ttype *TypeSymbol) *FunctionSymbol {
    var functionSymbol FunctionSymbol

    functionSymbol.Symbol = NewSymbol(name)
    functionSymbol.Parameter = parameter
    functionSymbol.Type = ttype
    functionSymbol.Name = name

    return &functionSymbol
}

func (t *FunctionSymbol) Kind() SymbolKind.SymbolKind {
    return SymbolKind.Function
}
