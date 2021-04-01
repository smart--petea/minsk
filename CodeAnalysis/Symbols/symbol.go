package Symbols

import (
    "minsk/CodeAnalysis/Symbols/SymbolKind"
)

type Symbol struct {
    Name  string
    Kind func() SymbolKind.SymbolKind
}

func NewSymbol(name string) *Symbol {
    return &Symbol{
        Name: name,
    }
}

func (s *Symbol) String() string {
    return s.Name
}
