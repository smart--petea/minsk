package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    SyntaxFacts "minsk/CodeAnalysis/Syntax/SyntaxFacts"
    "minsk/Util"
)

type Parser struct {
    Util.DiagnosticBag 

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

    parser.AddDiagnosticsRange(lexer.GetDiagnostics())
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

    p.ReportUnexpectedToken(current.Span(), current.Kind(), kind)

    return NewSyntaxToken(kind, current.Position, nil, nil)
}

func (p *Parser) ParsePrimaryExpression() ExpressionSyntax {
    switch p.Current().Kind() {
    case SyntaxKind.OpenParenthesisToken:
        left := p.NextToken()
        expression := p.ParseExpression(0)
        right := p.MatchToken(SyntaxKind.CloseParenthesisToken)

        return NewParenthesizedExpressionSyntax(left, expression, right)

    case SyntaxKind.FalseKeyword,  SyntaxKind.TrueKeyword:
        value := p.Current().Kind() == SyntaxKind.TrueKeyword
        return NewLiteralExpressionSyntax(p.NextToken(), value)

    case SyntaxKind.IdentifierToken:
        identifierToken := p.NextToken()

        return NewNameExpressionSyntax(identifierToken)

    default:
        numberToken := p.MatchToken(SyntaxKind.NumberToken)
        if numberToken == nil {
            return nil
        }

        return NewLiteralExpressionSyntax(numberToken, numberToken.Value())
    }
}

func (p *Parser) ParseExpression() ExpressionSyntax
{
    return p.ParseAssignmentExpression()
}

func (p *Parser) ParseAssignmentExpression() ExpressionSyntax {
    if p.Peek(0).Kind() == SyntaxKind.IdentifierToken && p.Peek(1).Kind() == SyntaxKind.EqualsToken {
        identifierToken = p.NextToken()
        operatorToken = p.NextToken()
        right := p.ParseAssignmentExpression()

        return NewParseAssignmentExpression(identifierToken, operatorToken, right)
    }

    return NewParseBinaryExpression(0)
}

func (p *Parser) ParseBinaryExpression(parentPrecedence int) ExpressionSyntax {
    var left ExpressionSyntax 

    unaryOperatorPrecedence := SyntaxFacts.GetUnaryOperatorPrecedence(p.Current().Kind())
    if  unaryOperatorPrecedence != 0 && unaryOperatorPrecedence >= parentPrecedence {
        operatorToken := p.NextToken()
        operand := p.ParseBinaryExpression(unaryOperatorPrecedence)

        left = NewUnaryExpressionSyntax(operatorToken, operand)
    } else {
        left = p.ParsePrimaryExpression()
    }

    for {
        precedence := SyntaxFacts.GetBinaryOperatorPrecedence(p.Current().Kind())
        if precedence == 0 || precedence <= parentPrecedence {
            break
        }

        operatorToken := p.NextToken()
        right := p.ParseBinaryExpression(precedence)
        left = NewBinaryExpressionSyntax(left, operatorToken, right)
    }

    return left
}
