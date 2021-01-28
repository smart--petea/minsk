package Syntax

import (
    "unicode"
    "strconv"
    "reflect"

    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    SyntaxFacts "minsk/CodeAnalysis/Syntax/SyntaxFacts"
    "minsk/Util"
)

type Lexer struct {
    Util.DiagnosticBag 

    Runes []rune
    Position int
}

func NewLexer(runes []rune) *Lexer {
    return &Lexer{
        Runes: runes,
    }
}

func (l *Lexer) Current() rune {
    return l.Peek(0)
}

func (l *Lexer) Lookahead() rune {
    return l.Peek(1)
}

func (l *Lexer) Peek(offset int) rune {
    index := l.Position + offset

    if index >= len(l.Runes) {
        return '\x00'
    }

    return l.Runes[index]
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
            l.ReportInvalidNumber(Util.NewTextSpan(start, length), runes, reflect.Int)
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
    switch current {
    case '+':
        l.Next()
        return NewSyntaxToken(SyntaxKind.PlusToken, position, l.Runes[position:l.Position], nil)

    case '-':
        l.Next()
        return NewSyntaxToken(SyntaxKind.MinusToken, position, l.Runes[position:l.Position], nil)

    case '*':
        l.Next()
        return NewSyntaxToken(SyntaxKind.StarToken, position, l.Runes[position:l.Position], nil)

    case '/':
        l.Next()
        return NewSyntaxToken(SyntaxKind.SlashToken, position, l.Runes[position:l.Position], nil)

    case '(':
        l.Next()
        return NewSyntaxToken(SyntaxKind.OpenParenthesisToken, position, l.Runes[position:l.Position], nil)

    case ')':
        l.Next()
        return NewSyntaxToken(SyntaxKind.CloseParenthesisToken, position, l.Runes[position:l.Position], nil)

    case '&':
        if l.Lookahead() == '&' {
            l.Next()
            l.Next()
            return NewSyntaxToken(SyntaxKind.AmpersandAmpersandToken, position, l.Runes[position:l.Position], nil)
        }

    case '|':
        if l.Lookahead() == '|' {
            l.Next()
            l.Next()
            return NewSyntaxToken(SyntaxKind.PipePipeToken, position, l.Runes[position:l.Position], nil)
        }

    case '=':
        if l.Lookahead() == '=' {
            l.Next()
            l.Next()
            return NewSyntaxToken(SyntaxKind.EqualsEqualsToken, position, l.Runes[position:l.Position], nil)
        } else {
            l.Next()
            return NewSyntaxToken(SyntaxKind.EqualsToken, position, l.Runes[position:l.Position], nil)
        }

    case '!':
        if l.Lookahead() == '=' {
            l.Next()
            l.Next()
            return NewSyntaxToken(SyntaxKind.BangEqualsToken, position, l.Runes[position:l.Position], nil)
        }

        l.Next()
        return NewSyntaxToken(SyntaxKind.BangToken, position, l.Runes[position:l.Position], nil)

    }

    l.Next()

    l.ReportBadCharacter(position, current)
    return NewSyntaxToken(SyntaxKind.BadToken, position, l.Runes[position:l.Position], nil)
}
