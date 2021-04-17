package Symbols

import (
    "minsk/CodeAnalysis/Symbols/SymbolKind"
    "minsk/CodeAnalysis/Syntax"
)

type FunctionSymbol struct {
    *Symbol

    Parameter []*ParameterSymbol
    Name string
    Type *TypeSymbol
    Declaration *Syntax.FunctionDeclarationSyntax
}

//declaration is optional. Should be provided as nil if not required
func NewFunctionSymbol(name string, parameter []*ParameterSymbol, ttype *TypeSymbol, declaration *Syntax.FunctionDeclarationSyntax) *FunctionSymbol {
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
