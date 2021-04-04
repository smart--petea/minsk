package BuiltinFunctions

import (
    "minsk/CodeAnalysis/Symbols"
)

var Print *FunctionSymbol = newFunctionSymbol("print", []*ParameterSymbol{newParameterSymbol("text", TypeSymbol.String)}, Symbols.TypeSymbolVoid)
var Input *FunctionSymbol = newFunctionSymbol("input", nil, Symbols.TypeSymbolString)
