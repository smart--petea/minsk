package BuiltinFunctions

import (
    "minsk/CodeAnalysis/Symbols"
)

var Print *Symbols.FunctionSymbol = Symbols.NewFunctionSymbol("print", []*Symbols.ParameterSymbol{Symbols.NewParameterSymbol("text", Symbols.TypeSymbolString)}, Symbols.TypeSymbolVoid)
var Input *Symbols.FunctionSymbol = Symbols.NewFunctionSymbol("input", nil, Symbols.TypeSymbolString)

func GetAll() <-chan *Symbols.FunctionSymbol {
    c := make(chan *Symbols.FunctionSymbol)

    go func () {
        defer close(c)

        c<-Print
        c<-Input
    }()

    return c
}
