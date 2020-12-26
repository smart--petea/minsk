package main

import (
    "fmt"
    "bufio"
    "os"
    "log"
    "strings"
    "unicode"
    "strconv"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("> ")
        line, err := reader.ReadString('\n')
        if err != nil {
            log.Fatal(err)
        }

        line = strings.TrimSpace(line)
        if len(line) == 0 {
            os.Exit(0)
        }

        lexer := NewLexer([]rune(line))
        for {
            token := lexer.NextToken()
            if token.Kind == EndOfFileToken {
                break
            }

            fmt.Printf("%s: '%s'", token.Kind, string(token.Runes))
            if token.Value != nil {
                fmt.Printf(" value: %v", token.Value)
            }
            fmt.Println()
        }
    }
}

type SyntaxKind string

const (
    NumberToken SyntaxKind = "NumberToken"
    WhitespaceToken SyntaxKind = "WhitespaceToken"
    PlusToken SyntaxKind = "PlusToken"
    EndOfFileToken SyntaxKind = "EndOfFileToken"
    MinusToken SyntaxKind = "MinusToken"
    StartToken SyntaxKind = "StartToken"
    SlashToken SyntaxKind = "SlashToken"
    OpenParenthesisToken SyntaxKind = "OpenParenthisToken"
    CloseParenthesisToken SyntaxKind = "CloseParenthisToken"
    BadToken SyntaxKind = "BadToken"
)

type SyntaxToken struct {
    Kind SyntaxKind
    Position int
    Runes []rune
    Value interface{}
}

func NewSyntaxToken(kind SyntaxKind, position int, runes []rune, value interface{}) *SyntaxToken {
    return &SyntaxToken{
        Kind: kind,
        Position: position,
        Runes: runes,
        Value: value,
    }
}

type Lexer struct {
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

func (l *Lexer) NextToken() *SyntaxToken {
    if l.Position >= len(l.Runes) {
        return NewSyntaxToken(EndOfFileToken, l.Position, []rune{'\x00'}, nil)
    } 

    if unicode.IsDigit(l.Current()) {
        start := l.Position

        for unicode.IsDigit(l.Current()) {
            l.Next()
        }

        length := l.Position - start 
        runes := l.Runes[start: start + length]
        val, _ := strconv.Atoi(string(runes))

        return NewSyntaxToken(NumberToken, start, runes, val)
    }

    if unicode.IsSpace(l.Current()) {
        start := l.Position

        for unicode.IsSpace(l.Current()) {
            l.Next()
        }

        length := l.Position - start 
        runes := l.Runes[start: start + length]

        return NewSyntaxToken(WhitespaceToken, start, runes, nil)
    }

    position := l.Position
    current := l.Current()
    l.Next()
    if current == '+' {
        return NewSyntaxToken(PlusToken, position, []rune{current}, nil)
    }

    if current == '-' {
        return NewSyntaxToken(MinusToken, position, []rune{current}, nil)
    }

    if current == '*' {
        return NewSyntaxToken(StartToken, position, []rune{current}, nil)
    }

    if current == '/' {
        return NewSyntaxToken(SlashToken, position, []rune{current}, nil)
    }

    if current == '(' {
        return NewSyntaxToken(OpenParenthesisToken, position, []rune{current}, nil)
    }

    if current == ')' {
        return NewSyntaxToken(CloseParenthesisToken, position, []rune{current}, nil)
    }

    return NewSyntaxToken(BadToken, position, []rune{current}, nil)
}
