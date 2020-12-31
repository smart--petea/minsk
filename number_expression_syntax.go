package minsk

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
