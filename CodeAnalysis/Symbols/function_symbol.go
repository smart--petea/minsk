package Symbols

import (
    "minsk/CodeAnalysis/Symbols/SymbolKind"
)

type FunctionSymbol struct {
    *Symbol

    Parameter []*ParameterSymbol
    Name string
    Type *TypeSymbol
    Declaration interface{} //*Syntax.FunctionDeclarationSyntax it could solved through interfaces
}

//declaration is optional. Should be provided as nil if not required
func NewFunctionSymbol(name string, parameter []*ParameterSymbol, ttype *TypeSymbol, declaration interface{}) *FunctionSymbol {
    var functionSymbol FunctionSymbol

    functionSymbol.Symbol = NewSymbol(name)
    functionSymbol.Parameter = parameter
    functionSymbol.Type = ttype
    functionSymbol.Name = name
    functionSymbol.Declaration = declaration

    return &functionSymbol
}

func (t *FunctionSymbol) Kind() SymbolKind.SymbolKind {
    return SymbolKind.Function
}
