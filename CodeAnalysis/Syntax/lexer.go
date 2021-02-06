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
    start int
    kind SyntaxKind.SyntaxKind
    value interface{}
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
    l.start = l.Position
    l.kind = SyntaxKind.BadToken
    l.value = nil

    current := l.Current()
    switch current {
    case '\x00':
        l.kind = SyntaxKind.EndOfFileToken

    case '+':
        l.Next()
        l.kind = SyntaxKind.PlusToken

    case '-':
        l.Next()
        l.kind = SyntaxKind.MinusToken

    case '*':
        l.Next()
        l.kind = SyntaxKind.StarToken

    case '/':
        l.Next()
        l.kind = SyntaxKind.SlashToken

    case '(':
        l.Next()
        l.kind = SyntaxKind.OpenParenthesisToken

    case ')':
        l.Next()
        l.kind = SyntaxKind.CloseParenthesisToken

    case '&':
        if l.Lookahead() == '&' {
            l.Next()
            l.Next()
            l.kind = SyntaxKind.AmpersandAmpersandToken
        }

    case '|':
        if l.Lookahead() == '|' {
            l.Next()
            l.Next()
            l.kind = SyntaxKind.PipePipeToken
        }

    case '=':
        if l.Lookahead() == '=' {
            l.Next()
            l.Next()
            l.kind = SyntaxKind.EqualsEqualsToken
        } else {
            l.Next()
            l.kind = SyntaxKind.EqualsToken
        }

    case '!':
        if l.Lookahead() == '=' {
            l.Next()
            l.Next()
            l.kind = SyntaxKind.BangEqualsToken
        } else {
            l.Next()
            l.kind = SyntaxKind.BangToken
        }
    case '0','1','2','3','4','5','6','7','8','9':
            l.ReadNumberToken() 
    case ' ','\t','\n','\r':
            l.ReadWhiteSpace()

    default:
        if unicode.IsLetter(current) {
            l.ReadIdentifierOrKeyword()
        } else if unicode.IsSpace(current) {
            l.ReadWhiteSpace()
        } else {
            l.Next()
            l.ReportBadCharacter(l.start, current)
        }
    }

    var runes []rune
    text := SyntaxFacts.GetText(l.kind)
    if text == "" {
        runes = l.Runes[l.start:l.Position]
    } else {
        runes = []rune(text)
    }

    return NewSyntaxToken(l.kind, l.start, runes, l.value)
}

func (l *Lexer) ReadNumberToken() {
    for unicode.IsDigit(l.Current()) {
        l.Next()
    }

    length := l.Position - l.start 
    runes := l.Runes[l.start: l.start + length]
    value, err := strconv.Atoi(string(runes))
    if err != nil {
        l.ReportInvalidNumber(Util.NewTextSpan(l.start, length), runes, reflect.Int)
    }

    l.value = value
    l.kind = SyntaxKind.NumberToken
}

func (l *Lexer) ReadWhiteSpace() {
    for unicode.IsSpace(l.Current()) {
        l.Next()
    }

    l.kind = SyntaxKind.WhitespaceToken
}

func (l *Lexer) ReadIdentifierOrKeyword() {
    for unicode.IsLetter(l.Current()) {
        l.Next()
    }

    runes := l.Runes[l.start: l.Position]
    l.kind = SyntaxFacts.GetKeywordKind(string(runes))
}
