package Symbols

import (
    "minsk/CodeAnalysis/Symbols/SymbolKind"
)

type ISymbol interface {
    GetName() string
    Kind() SymbolKind.SymbolKind
}
