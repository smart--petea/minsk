package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    SyntaxFacts "minsk/CodeAnalysis/Syntax/SyntaxFacts"
    "minsk/CodeAnalysis/Text"
    "minsk/Util"
)

type Parser struct {
    Util.DiagnosticBag 

    Tokens []SyntaxToken
    Position int
    Text *Text.SourceText
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

func NewParser(text *Text.SourceText) *Parser {
    lexer := NewLexer(text)
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
        Text: text,
        Tokens: tokens,
    }

    parser.AddDiagnosticsRange(lexer.GetDiagnostics())
    return parser
}

func (p *Parser) ParseCompilationUnit() *CompilationUnitSyntax {
    statement  := p.ParseStatement()
    endOfFileToken := p.MatchToken(SyntaxKind.EndOfFileToken)

    return NewCompilationUnitSyntax(statement, endOfFileToken)
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

    p.ReportUnexpectedToken(SyntaxNodeToTextSpan(current), current.Kind(), kind)

    return NewSyntaxToken(kind, current.Position, nil, nil)
}

func (p *Parser) ParsePrimaryExpression() ExpressionSyntax {
    switch p.Current().Kind() {
    case SyntaxKind.OpenParenthesisToken:
        return p.ParseParenthesizedExpression()
    case SyntaxKind.FalseKeyword,  SyntaxKind.TrueKeyword:
        return p.ParseBooleanLiteral()
    case SyntaxKind.NumberToken:
        return p.ParseNumberLiteral()
    case SyntaxKind.IdentifierToken:
        return p.ParseNameExpression()
    default:
        return p.ParseNameExpression()
    }
}

func (p *Parser) ParseNumberLiteral() ExpressionSyntax {
    numberToken := p.MatchToken(SyntaxKind.NumberToken)
    if numberToken == nil {
        return nil
    }

    return NewLiteralExpressionSyntax(numberToken, numberToken.Value())
}

func (p *Parser) ParseParenthesizedExpression() ExpressionSyntax {
    left := p.MatchToken(SyntaxKind.OpenParenthesisToken)
    expression := p.ParseExpression()
    right := p.MatchToken(SyntaxKind.CloseParenthesisToken)

    return NewParenthesizedExpressionSyntax(left, expression, right)
}

func (p *Parser) ParseBooleanLiteral() ExpressionSyntax {
    isTrue := p.Current().Kind() == SyntaxKind.TrueKeyword

    var keywordToken *SyntaxToken
    if isTrue {
        keywordToken = p.MatchToken(SyntaxKind.TrueKeyword)
    } else {
        keywordToken = p.MatchToken(SyntaxKind.FalseKeyword)
    }

    return NewLiteralExpressionSyntax(keywordToken, isTrue)
}

func (p *Parser) ParseNameExpression() ExpressionSyntax {
    identifierToken := p.MatchToken(SyntaxKind.IdentifierToken)
    return NewNameExpressionSyntax(identifierToken)
}

func (p *Parser) ParseExpression() ExpressionSyntax {
    return p.ParseAssignmentExpression()
}

func (p *Parser) ParseAssignmentExpression() ExpressionSyntax {
    if p.Peek(0).Kind() == SyntaxKind.IdentifierToken && p.Peek(1).Kind() == SyntaxKind.EqualsToken {
        identifierToken := p.NextToken()
        operatorToken := p.NextToken()
        right := p.ParseAssignmentExpression()

        return NewAssignmentExpressionSyntax(identifierToken, operatorToken, right)
    }

    return p.ParseBinaryExpression(0)
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

func (p *Parser) ParseStatement() StatementSyntax {
    if p.Current().Kind() == SyntaxKind.OpenBraceToken {
        return p.ParseBlockStatement()
    }

    return p.ParseExpressionStatement()
}

func (p *Parser) ParseBlockStatement() BlockStatementSyntax {
    var statements []StatementSyntax

    openBraceToken := p.MatchToken(SyntaxKind.OpenBraceToken)
    for p.Current().Kind() != SyntaxKind.EndOfFileToken && p.Current().Kind() != SyntaxKind.CloseBraceToken {
        statement := p.ParseStatement()
        statements = append(statements, statement)
    }
    closeBraceToken := p.MatchToken(SyntaxKind.CloseBraceToken)

    return NewBlockStatementSyntax(openBraceToken, statements, closeBraceToken)
}

func (p *Parser) ParseExpressionStatement() StatementSyntax {
    expression := p.ParseExpression()
    return NewExpressionStatementSyntax(expression)
}
