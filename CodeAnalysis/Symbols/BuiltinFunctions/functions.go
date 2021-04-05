package BuiltinFunctions

import (
    "minsk/CodeAnalysis/Symbols"
)

var Print *FunctionSymbol = newFunctionSymbol("print", []*ParameterSymbol{newParameterSymbol("text", TypeSymbol.String)}, Symbols.TypeSymbolVoid)
var Input *FunctionSymbol = newFunctionSymbol("input", nil, Symbols.TypeSymbolString)

func GetAll() <-chan *FunctionSymbol {
    c := make(chan *FunctionSymbol)

    go func () {
        defer close(c)

        c<-Print
        c<-Input
    }()

    return c
}
