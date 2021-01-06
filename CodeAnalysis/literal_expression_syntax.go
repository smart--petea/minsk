package CodeAnalysis

type LiteralExpressionSyntax struct {
    LiteralToken *SyntaxToken
}

func (n *LiteralExpressionSyntax) Value() interface{} {
    return nil
}

func (n *LiteralExpressionSyntax) GetChildren() []SyntaxNode {
    return []SyntaxNode{n.LiteralToken}
}

func (n *LiteralExpressionSyntax) Kind() SyntaxKind {
    return NumberExpression
}

func NewLiteralExpressionSyntax(literalToken *SyntaxToken) *LiteralExpressionSyntax {
    return &LiteralExpressionSyntax{
        LiteralToken: literalToken,
    }
}
