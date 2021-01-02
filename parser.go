package minsk

import (
    "fmt"
)

type Parser struct {
    Tokens []SyntaxToken
    Position int
    Diagnostics []string
}

func (p *Parser) Peek(offset int) *SyntaxToken {
    index := p.Position + offset
    if index >= len(p.Tokens) {
        return nil
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

        if token.Kind() == EndOfFileToken {
            break
        }

        if token.Kind() == WhitespaceToken {
            continue
        }

        if token.Kind() == BadToken {
            continue
        }

        tokens = append(tokens, *token)
    }

    return &Parser{
        Tokens: tokens,
        Diagnostics: lexer.Diagnostics,
    }
}

func (p *Parser) AddDiagnostic(format string, args ...interface{}) {
    p.Diagnostics = append(p.Diagnostics, fmt.Sprintf(format, args...))
}

func (p *Parser) Parse() *SyntaxTree {
    rootExpression := p.ParseTerm()
    endOfFileToken := p.Match(EndOfFileToken)

    return NewSyntaxTree(p.Diagnostics, rootExpression, endOfFileToken)
}

func (p *Parser) ParseTerm() ExpressionSyntax {
    var left = p.ParseFactor()

    for p.Current() != nil && (p.Current().Kind() == PlusToken || p.Current().Kind() == MinusToken) {
            operatorToken := p.NextToken()
            right := p.ParseFactor()
            left = NewBinaryExpressionSyntax(left, operatorToken, right)
    }

    return left
}

func (p *Parser) ParseFactor() ExpressionSyntax {
    var left = p.ParsePrimaryExpression()
    var right ExpressionSyntax

    for p.Current() != nil && (p.Current().Kind() == StarToken || p.Current().Kind() == SlashToken) {
            operatorToken := p.NextToken()
            right = p.ParsePrimaryExpression()
            left = NewBinaryExpressionSyntax(left, operatorToken, right)
    }

    return left
}

func (p *Parser) NextToken() *SyntaxToken {
    current := p.Current()
    if current != nil {
        p.Position = p.Position + 1
    }

    return current
}

func (p *Parser) Match(kind SyntaxKind) *SyntaxToken {
    current := p.Current()
    if current == nil {
        return nil
    }

    if current.Kind() == kind {
        return p.NextToken()
    }

    p.AddDiagnostic("ERROR: Unexpected token <%s>, expected <%s>", current.Kind(), kind)

    return NewSyntaxToken(kind, current.Position, nil, nil)
}

func (p *Parser) ParseExpression() ExpressionSyntax {
    return p.ParseTerm()
}

func (p *Parser) ParsePrimaryExpression() ExpressionSyntax {
    current := p.Current()
    if current.Kind() == OpenParenthesisToken {
        left := p.NextToken()
        expression := p.ParseExpression()
        right := p.Match(CloseParenthesisToken)

        return NewParenthesizedExpressionSyntax(left, expression, right)
    }

    numberToken := p.Match(NumberToken)

    if numberToken == nil {
        return nil
    }

    return NewNumberExpressionSyntax(numberToken)
}
