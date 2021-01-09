package CodeAnalysis

import (
    "minsk/CodeAnalysis/SyntaxKind"
)

type LiteralExpressionSyntax struct {
    LiteralToken *SyntaxToken
}

func (n *LiteralExpressionSyntax) Value() interface{} {
    return nil
}

func (n *LiteralExpressionSyntax) GetChildren() []SyntaxNode {
    return []SyntaxNode{n.LiteralToken}
}

func (n *LiteralExpressionSyntax) Kind() SyntaxKind.SyntaxKind {
    return SyntaxKind.LiteralExpression
}

func NewLiteralExpressionSyntax(literalToken *SyntaxToken) *LiteralExpressionSyntax {
    return &LiteralExpressionSyntax{
        LiteralToken: literalToken,
    }
}
