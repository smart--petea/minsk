package minsk

type ParenthesizedExpressionSyntax struct {
    OpenParenthesisToken *SyntaxToken
    Expression ExpressionSyntax
    CloseParenthesisToken *SyntaxToken
}

func NewParenthesizedExpressionSyntax(openParenthesisToken *SyntaxToken, expression ExpressionSyntax, closeParenthesisToken *SyntaxToken)  *ParenthesizedExpressionSyntax {
    return &ParenthesizedExpressionSyntax{
        OpenParenthesisToken: openParenthesisToken,
        Expression: expression,
        CloseParenthesisToken: closeParenthesisToken,
    }
}

func (p *ParenthesizedExpressionSyntax) Kind() SyntaxKind {
    return ParenthesizedExpression
}

func (p *ParenthesizedExpressionSyntax) GetChildren() []SyntaxNode {
    return []SyntaxNode{
        p.OpenParenthesisToken,
        p.Expression,
        p.CloseParenthesisToken,
    }
}

func (p *ParenthesizedExpressionSyntax) Value() interface{} {
    return nil
}
