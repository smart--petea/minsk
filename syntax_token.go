package minsk

type SyntaxToken struct {
    kind SyntaxKind
    Position int
    Runes []rune
    value interface{}
}

func (s *SyntaxToken) Kind() SyntaxKind {
    return s.kind
}

func (s *SyntaxToken) Value() interface{} {
    return s.value
}

func (s *SyntaxToken) GetChildren() []SyntaxNode {
    return nil
}

func NewSyntaxToken(kind SyntaxKind, position int, runes []rune, value interface{}) *SyntaxToken {
    return &SyntaxToken{
        kind: kind,
        Position: position,
        Runes: runes,
        value: value,
    }
}
