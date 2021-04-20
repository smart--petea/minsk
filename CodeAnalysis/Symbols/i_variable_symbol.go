package Symbols

import (
    "minsk/CodeAnalysis/Symbols/SymbolKind"
)

type IVariableSymbol interface {
    ISymbol

    GetType() *TypeSymbol
    IsReadOnly() bool
    Kind() SymbolKind.SymbolKind
}
