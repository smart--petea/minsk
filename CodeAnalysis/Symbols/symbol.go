package Symbols

import (
    "minsk/CodeAnalysis/Symbols/SymbolKind"
)

type Symbol struct {
    name  string
    Kind func() SymbolKind.SymbolKind
}

func NewSymbol(name string) *Symbol {
    return &Symbol{
        name: name,
    }
}

func (s *Symbol) String() string {
    return s.name
}

func (s *Symbol) GetName() string {
    return s.name
}
