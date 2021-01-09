package CodeAnalysis

import (
    "minsk/CodeAnalysis/SyntaxKind"
)

type SyntaxToken struct {
    kind SyntaxKind.SyntaxKind
    Position int
    Runes []rune
    value interface{}
}

func (s *SyntaxToken) Kind() SyntaxKind.SyntaxKind {
    return s.kind
}

func (s *SyntaxToken) Value() interface{} {
    return s.value
}

func (s *SyntaxToken) GetChildren() []SyntaxNode {
    return nil
}

func NewSyntaxToken(kind SyntaxKind.SyntaxKind, position int, runes []rune, value interface{}) *SyntaxToken {
    return &SyntaxToken{
        kind: kind,
        Position: position,
        Runes: runes,
        value: value,
    }
}
