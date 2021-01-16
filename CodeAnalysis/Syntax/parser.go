package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/Util"
)

type Parser struct {
    Util.Diagnostic

    Tokens []SyntaxToken
    Position int
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
        token = lexer.Lex()

        if token.Kind() == SyntaxKind.EndOfFileToken {
            tokens = append(tokens, *token)
            break
        }

        if token.Kind() == SyntaxKind.WhitespaceToken {
            continue
        }

        if token.Kind() == SyntaxKind.BadToken {
            continue
        }

        tokens = append(tokens, *token)
    }

    parser := &Parser{
        Tokens: tokens,
    }

    parser.LoadDiagnostics(lexer.GetDiagnostics())
    return parser
}

func (p *Parser) Parse() (ExpressionSyntax, *SyntaxToken) {
    var rootExpression ExpressionSyntax = p.ParseExpression(0)

    endOfFileToken := p.MatchToken(SyntaxKind.EndOfFileToken)

    return rootExpression, endOfFileToken
}

func (p *Parser) NextToken() *SyntaxToken {
    current := p.Current()
    if current != nil {
        p.Position = p.Position + 1
    }

    return current
}

func (p *Parser) MatchToken(kind SyntaxKind.SyntaxKind) *SyntaxToken {
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

func (p *Parser) ParsePrimaryExpression() ExpressionSyntax {
    if p.Current().Kind() == SyntaxKind.OpenParenthesisToken {
        left := p.NextToken()
        expression := p.ParseExpression(0)
        right := p.MatchToken(SyntaxKind.CloseParenthesisToken)

        return NewParenthesizedExpressionSyntax(left, expression, right)
    }

    numberToken := p.MatchToken(SyntaxKind.NumberToken)

    if numberToken == nil {
        return nil
    }

    return NewLiteralExpressionSyntax(numberToken)
}

func (p *Parser) ParseExpression(parentPrecedence int) ExpressionSyntax {
    var left ExpressionSyntax 

    unaryOperatorPrecedence := p.Current().Kind().GetUnaryOperatorPrecedence()
    if  unaryOperatorPrecedence != 0 && unaryOperatorPrecedence >= parentPrecedence {
        operatorToken := p.NextToken()
        operand := p.ParseExpression(unaryOperatorPrecedence)

        left = NewUnaryExpressionSyntax(operatorToken, operand)
    } else {
        left = p.ParsePrimaryExpression()
    }

    for {
        precedence := p.Current().Kind().GetBinaryOperatorPrecedence()
        if precedence == 0 || precedence <= parentPrecedence {
            break
        }

        operatorToken := p.NextToken()
        right := p.ParseExpression(precedence)
        left = NewBinaryExpressionSyntax(left, operatorToken, right)
    }

    return left
}
