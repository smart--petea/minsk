package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    SyntaxFacts "minsk/CodeAnalysis/Syntax/SyntaxFacts"
    "minsk/CodeAnalysis/Text"
    "minsk/Util"
)

type Parser struct {
    *Util.DiagnosticBag 

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
        DiagnosticBag: Util.NewDiagnosticBag(),

        Text: text,
        Tokens: tokens,
    }

    parser.DiagnosticBag.AddRange(lexer.DiagnosticBag)
    return parser
}

func (p *Parser) ParseCompilationUnit() *CompilationUnitSyntax {
    members  := p.ParseMembers()
    endOfFileToken := p.MatchToken(SyntaxKind.EndOfFileToken)

    return NewCompilationUnitSyntax(members, endOfFileToken)
}


func (p *Parser) ParseMembers() []MemberSyntax {
    var members []MemberSyntax

    for p.Current().Kind() != SyntaxKind.EndOfFileToken {
        startToken := p.Current()

        member := p.ParseMember()
        members = append(members, member)

        //If ParseStatement() did not consume any tokens,
        //we need to skip the current token and continue.
        //In order to avoid an infinite loop.
        //
        // We do not need to report and error, because we'll
        //already tried to parse an expression statement 
        //and reported one.
        if p.Current() == startToken {
            p.NextToken()
        }
    }
    //closeBraceToken := p.MatchToken(SyntaxKind.CloseBraceToken)

    return members
}

func (p *Parser) ParseMember() MemberSyntax {
    if p.Current().Kind() == SyntaxKind.FunctionKeyword {
        return p.ParseFunctionDeclaration()
    }

    return p.ParseGlobalStatement()
}

func (p *Parser) ParseFunctionDeclaration() MemberSyntax {
    functionKeyword := p.MatchToken(SyntaxKind.FunctionKeyword)
    identifier := p.MatchToken(SyntaxKind.IdentifierToken)
    openParenthesisToken := p.MatchToken(SyntaxKind.OpenParenthesisToken)
    parameters := p.ParseParameterList()
    closeParenthesisToken := p.MatchToken(SyntaxKind.CloseParenthesisToken)
    ttype := p.ParseOptionalTypeClause();
    body := p.ParseBlockStatement()
    return NewFunctionDeclarationSyntax(functionKeyword, identifier, openParenthesisToken, parameters, closeParenthesisToken, ttype, body)
}

func (p *Parser) ParseParameterList() *SeparatedSyntaxList {
    var nodeAndSeparators []SyntaxNode

    for p.Current().Kind() != SyntaxKind.CloseParenthesisToken && p.Current().Kind() != SyntaxKind.EndOfFileToken {
        parameter := p.ParseParameter()
        nodeAndSeparators = append(nodeAndSeparators, parameter)

        if p.Current().Kind() != SyntaxKind.CloseParenthesisToken {
            comma := p.MatchToken(SyntaxKind.CommaToken)
            nodeAndSeparators = append(nodeAndSeparators, comma)
        }
    }

    return NewSeparatedSyntaxList(nodeAndSeparators)
}

func (p *Parser) ParseParameter() *ParameterSyntax {
    identifier := p.MatchToken(SyntaxKind.IdentifierToken)
    ttype := p.ParseTypeClause()
    return NewParameterSyntax(identifier, ttype)
}

func (p *Parser) ParseGlobalStatement() MemberSyntax {
    statement := p.ParseStatement()
    return NewGlobalStatementSyntax(statement)
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

    p.ReportUnexpectedToken(current.GetSpan(), current.Kind(), kind)

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

    case SyntaxKind.StringToken:
        return p.ParseStringLiteral()

    case SyntaxKind.IdentifierToken:
        return p.ParseNameOrCallExpression()

    default:
        return p.ParseNameOrCallExpression()
    }
}

func (p *Parser) ParseStringLiteral() ExpressionSyntax {
    stringToken := p.MatchToken(SyntaxKind.StringToken)
    if stringToken == nil {
        return nil
    }

    return NewLiteralExpressionSyntax(stringToken, stringToken.Value())
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

func (p *Parser) ParseNameOrCallExpression() ExpressionSyntax {
    if p.Peek(0).Kind() == SyntaxKind.IdentifierToken && p.Peek(1).Kind() == SyntaxKind.OpenParenthesisToken {
        return p.ParseCallExpression()
    }

    return p.ParseNameExpression()
}

func (p *Parser) ParseExpression() ExpressionSyntax {
    return p.ParseAssignmentExpression()
}

func (p *Parser) ParseCallExpression() ExpressionSyntax {
    identifier := p.MatchToken(SyntaxKind.IdentifierToken)
    openParenthesisToken := p.MatchToken(SyntaxKind.OpenParenthesisToken)
    arguments := p.ParseArguments()
    closeParenthesisToken := p.MatchToken(SyntaxKind.CloseParenthesisToken)

    return NewCallExpressionSyntax(identifier, openParenthesisToken, arguments, closeParenthesisToken)
}

func (p *Parser) ParseArguments() *SeparatedSyntaxList {
    var nodeAndSeparators []SyntaxNode

    for p.Current().Kind() != SyntaxKind.CloseParenthesisToken && p.Current().Kind() != SyntaxKind.EndOfFileToken {
        expression := p.ParseExpression()
        nodeAndSeparators = append(nodeAndSeparators, expression)

        if p.Current().Kind() != SyntaxKind.CloseParenthesisToken {
            comma := p.MatchToken(SyntaxKind.CommaToken)
            nodeAndSeparators = append(nodeAndSeparators, comma)
        }
    }

    return NewSeparatedSyntaxList(nodeAndSeparators)
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
    switch p.Current().Kind() {
    case SyntaxKind.OpenBraceToken:
        return p.ParseBlockStatement()
    case SyntaxKind.LetKeyword, SyntaxKind.VarKeyword:
        return p.ParseVariableDeclaration()
    case SyntaxKind.IfKeyword:
        return p.ParseIfStatement()
    case SyntaxKind.WhileKeyword:
        return p.ParseWhileStatement()
    case SyntaxKind.ForKeyword:
        return p.ParseForStatement()
    default:
        return p.ParseExpressionStatement()
    }
}

func (p *Parser) ParseVariableDeclaration() *VariableDeclarationSyntax {
    var expected SyntaxKind.SyntaxKind
    if p.Current().Kind() == SyntaxKind.LetKeyword {
        expected = SyntaxKind.LetKeyword
    } else {
        expected = SyntaxKind.VarKeyword
    }

    keyword := p.MatchToken(expected)
    identifier := p.MatchToken(SyntaxKind.IdentifierToken)
    typeClause := p.ParseOptionalTypeClause()
    equals := p.MatchToken(SyntaxKind.EqualsToken)
    initializer := p.ParseExpression()
    return NewVariableDeclarationSyntax(keyword, identifier, typeClause, equals, initializer)
}

func (p *Parser) ParseOptionalTypeClause() *TypeClauseSyntax {
    if p.Current().Kind() != SyntaxKind.ColonToken {
        return nil
    }

    return p.ParseTypeClause()
}

func (p *Parser) ParseTypeClause() *TypeClauseSyntax {
    colonToken := p.MatchToken(SyntaxKind.ColonToken) 
    identifier := p.MatchToken(SyntaxKind.IdentifierToken) 

    return NewTypeClauseSyntax(colonToken, identifier)
}

func (p *Parser) ParseIfStatement() *IfStatementSyntax {
    keyword := p.MatchToken(SyntaxKind.IfKeyword)
    condition := p.ParseExpression()
    statement := p.ParseStatement()
    elseClause := p.ParseElseClause()

    return NewIfStatementSyntax(keyword, condition, statement, elseClause) 
}

func (p *Parser) ParseForStatement() *ForStatementSyntax {
    keyword := p.MatchToken(SyntaxKind.ForKeyword)
    identifier := p.MatchToken(SyntaxKind.IdentifierToken)
    equalsToken := p.MatchToken(SyntaxKind.EqualsToken)
    lowerBound := p.ParseExpression()
    toKeyword := p.MatchToken(SyntaxKind.ToKeyword)
    upperBound := p.ParseExpression()
    body := p.ParseStatement()

    return NewForStatementSyntax(keyword, identifier, equalsToken, lowerBound, toKeyword, upperBound, body)
}

func (p *Parser) ParseWhileStatement() *WhileStatementSyntax {
    keyword := p.MatchToken(SyntaxKind.WhileKeyword)
    condition := p.ParseExpression()
    body := p.ParseStatement()

    return NewWhileStatementSyntax(keyword, condition, body)
}

func (p *Parser) ParseElseClause() *ElseClauseSyntax {
    if p.Current().Kind() != SyntaxKind.ElseKeyword {
        return nil
    }

    elseKeyword := p.NextToken()
    elseStatement := p.ParseStatement()

    return NewElseClauseSyntax(elseKeyword, elseStatement)
}

func (p *Parser) ParseBlockStatement() *BlockStatementSyntax {
    var statements []StatementSyntax

    openBraceToken := p.MatchToken(SyntaxKind.OpenBraceToken)
    for p.Current().Kind() != SyntaxKind.EndOfFileToken && p.Current().Kind() != SyntaxKind.CloseBraceToken {
        startToken := p.Current()

        statement := p.ParseStatement()
        statements = append(statements, statement)

        //If ParseStatement() did not consume any tokens,
        //we need to skip the current token and continue.
        //In order to avoid an infinite loop.
        //
        // We do not need to report and error, because we'll
        //already tried to parse an expression statement 
        //and reported one.
        if p.Current() == startToken {
            p.NextToken()
        }
    }
    closeBraceToken := p.MatchToken(SyntaxKind.CloseBraceToken)

    return NewBlockStatementSyntax(openBraceToken, statements, closeBraceToken)
}

func (p *Parser) ParseExpressionStatement() StatementSyntax {
    expression := p.ParseExpression()
    return NewExpressionStatementSyntax(expression)
}
