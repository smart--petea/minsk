package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/CodeAnalysis/Text"
    "minsk/Util"
)

type SyntaxToken struct {
    *Util.ChildrenProvider

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

func NewSyntaxToken(kind SyntaxKind.SyntaxKind, position int, runes []rune, value interface{}) *SyntaxToken {
    return &SyntaxToken{
        ChildrenProvider: Util.NewChildrenProvider(),

        kind: kind,
        Position: position,
        Runes: runes,
        value: value,
    }
}

func (s *SyntaxToken) GetSpan() *Text.TextSpan {
    return SyntaxNodeToTextSpan(s)
}
