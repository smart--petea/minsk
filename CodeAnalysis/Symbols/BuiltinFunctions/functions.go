package BuiltinFunctions

import (
    "minsk/CodeAnalysis/Symbols"
)

var Print *Symbols.FunctionSymbol = Symbols.NewFunctionSymbol("print", []*Symbols.ParameterSymbol{Symbols.NewParameterSymbol("text", Symbols.TypeSymbolString)}, Symbols.TypeSymbolVoid, nil)
var Input *Symbols.FunctionSymbol = Symbols.NewFunctionSymbol("input", nil, Symbols.TypeSymbolString, nil)
var Rnd *Symbols.FunctionSymbol = Symbols.NewFunctionSymbol("rnd", []*Symbols.ParameterSymbol{Symbols.NewParameterSymbol("text", Symbols.TypeSymbolInt)}, Symbols.TypeSymbolInt, nil)

func GetAll() <-chan *Symbols.FunctionSymbol {
    c := make(chan *Symbols.FunctionSymbol)

    go func () {
        defer close(c)

        c<-Print
        c<-Input
        c<-Rnd
    }()

    return c
}
