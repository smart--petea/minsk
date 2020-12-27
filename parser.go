package minsk

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
