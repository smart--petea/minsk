package Syntax

import (
    "unicode"
    "strconv"

    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    SyntaxFacts "minsk/CodeAnalysis/Syntax/SyntaxFacts"
    "minsk/Util"
)

type Lexer struct {
    Util.Diagnostic

    Runes []rune
    Position int
}

func NewLexer(runes []rune) *Lexer {
    return &Lexer{
        Runes: runes,
    }
}

func (l *Lexer) Current() rune {
    if l.Position >= len(l.Runes) {
        return '\x00'
    }

    return l.Runes[l.Position]
}

func (l *Lexer) Next() {
    l.Position = l.Position + 1
}

func (l *Lexer) Lex() *SyntaxToken {
    if l.Position >= len(l.Runes) {
        return NewSyntaxToken(SyntaxKind.EndOfFileToken, l.Position, []rune{'\x00'}, nil)
    } 

    if unicode.IsDigit(l.Current()) {
        start := l.Position

        for unicode.IsDigit(l.Current()) {
            l.Next()
        }

        length := l.Position - start 
        runes := l.Runes[start: start + length]
        val, err := strconv.Atoi(string(runes))
        if err != nil {
            l.AddDiagnostic("The number %s isn't valid int32.", string(runes))
        }

        return NewSyntaxToken(SyntaxKind.NumberToken, start, runes, val)
    }

    if unicode.IsSpace(l.Current()) {
        start := l.Position

        for unicode.IsSpace(l.Current()) {
            l.Next()
        }

        length := l.Position - start 
        runes := l.Runes[start: start + length]

        return NewSyntaxToken(SyntaxKind.WhitespaceToken, start, runes, nil)
    }

    if unicode.IsLetter(l.Current()) {
        start := l.Position

        for unicode.IsLetter(l.Current()) {
            l.Next()
        }

        length := l.Position - start 
        runes := l.Runes[start: start + length]
        kind := SyntaxFacts.GetKeywordKind(string(runes))

        return NewSyntaxToken(kind, start, runes, nil)
    }

    position := l.Position
    current := l.Current()
    l.Next()
    switch current {
    case '+':
        return NewSyntaxToken(SyntaxKind.PlusToken, position, []rune{current}, nil)

    case '-':
        return NewSyntaxToken(SyntaxKind.MinusToken, position, []rune{current}, nil)

    case '*':
        return NewSyntaxToken(SyntaxKind.StarToken, position, []rune{current}, nil)

    case '/':
        return NewSyntaxToken(SyntaxKind.SlashToken, position, []rune{current}, nil)

    case '(':
        return NewSyntaxToken(SyntaxKind.OpenParenthesisToken, position, []rune{current}, nil)

    case ')':
        return NewSyntaxToken(SyntaxKind.CloseParenthesisToken, position, []rune{current}, nil)
    }

    l.AddDiagnostic("ERROR: bad character input: '%s'", string(current))
    return NewSyntaxToken(SyntaxKind.BadToken, position, []rune{current}, nil)
}
