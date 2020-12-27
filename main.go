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

        parser := NewParser(line)
        expression := parser.Parse()

        fmt.Print("\033[90m")
        PrettyPrint(expression, "")
        fmt.Print("\033[37m")
    }
}

func PrettyPrint(node SyntaxNode, indent string) {
    fmt.Print(node.Kind())

    if node.Value() != nil {
        fmt.Print(" ")
        fmt.Print(node.Value())
    }

    indent = indent + "   "
    for _, child := range node.GetChildren() {
        PrettyPrint(child, indent)
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
    BinaryExpression SyntaxKind = "BinaryExpression" 
    NumberExpression SyntaxKind = "NumberExpression" 
)

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

type Parser struct {
    Tokens []SyntaxToken
    Position int
}

func (p *Parser) Peek(offset int) *SyntaxToken {
    index := p.Position + offset
    if index > len(p.Tokens) {
        return &p.Tokens[len(p.Tokens) - 1]
    }

    return &p.Tokens[index]
}

func (p *Parser) Current() *SyntaxToken {
    return p.Peek(0)
}

func NewParser(text string) *Parser {
    lexer := NewLexer([]rune(text))
    var token *SyntaxToken
    var tokens []SyntaxToken

    for {
        token = lexer.NextToken()

        if token.Kind() != EndOfFileToken {
            break
        }

        if token.Kind() != WhitespaceToken && token.Kind() != BadToken {
            tokens = append(tokens, *token)
        }
    }

    return &Parser{
        Tokens: tokens,
    }
}

type SyntaxNode interface {
    Kind() SyntaxKind
    Value()  interface{}
    GetChildren() []SyntaxNode //todo IEnumerable
}

type ExpressionSyntax interface {
    SyntaxNode
}

type NumberExpressionSyntax struct {
    NumberToken *SyntaxToken
}

func (n *NumberExpressionSyntax) Value() interface{} {
    return nil
}

func (n *NumberExpressionSyntax) GetChildren() []SyntaxNode {
    return []SyntaxNode{n.NumberToken}
}

func (n *NumberExpressionSyntax) Kind() SyntaxKind {
    return NumberExpression
}

func NewNumberExpressionSyntax(numberToken *SyntaxToken) *NumberExpressionSyntax {
    return &NumberExpressionSyntax{
        NumberToken: numberToken,
    }
}

type BinaryExpressionSyntax struct {
    Left ExpressionSyntax
    Right ExpressionSyntax
    OperatorNode SyntaxNode
}

func (b *BinaryExpressionSyntax) Value() interface{} {
    return nil
}

func NewBinaryExpressionSyntax(left ExpressionSyntax, operatorNode SyntaxNode, right ExpressionSyntax) *BinaryExpressionSyntax {
    return &BinaryExpressionSyntax{
        Left: left,
        Right: right,
        OperatorNode: operatorNode,
    }
}

func (b *BinaryExpressionSyntax) Kind() SyntaxKind {
    return BinaryExpression
}

func (b *BinaryExpressionSyntax) GetChildren() []SyntaxNode {
    //todo yield operator in go
    return []SyntaxNode{
        b.Left,
        b.OperatorNode,
        b.Right,
    }
}

func (p *Parser) Parse() ExpressionSyntax {
    var left = p.ParsePrimaryExpression()
    var right ExpressionSyntax

    for p.Current().Kind() == PlusToken || p.Current().Kind() == MinusToken {
            operatorToken := p.NextToken()
            right = p.ParsePrimaryExpression()
            left = NewBinaryExpressionSyntax(left, operatorToken, right)
    }

    return left
}

func (p *Parser) NextToken() *SyntaxToken {
    current := p.Current()
    p.Position = p.Position + 1

    return current
}

func (p *Parser) Match(kind SyntaxKind) *SyntaxToken {
    if p.Current().Kind() == kind {
        return p.NextToken()
    }

    return NewSyntaxToken(kind, p.Current().Position, nil, nil)
}

func (p *Parser) ParsePrimaryExpression() ExpressionSyntax {
    numberToken := p.Match(NumberToken)

    return NewNumberExpressionSyntax(numberToken)
}
