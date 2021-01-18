package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
)

type LiteralExpressionSyntax struct {
    LiteralToken *SyntaxToken
    value interface{}
}

func (n *LiteralExpressionSyntax) Value() interface{} {
    return n.value
}

func (n *LiteralExpressionSyntax) GetChildren() []SyntaxNode {
    return []SyntaxNode{n.LiteralToken}
}

func (n *LiteralExpressionSyntax) Kind() SyntaxKind.SyntaxKind {
    return SyntaxKind.LiteralExpression
}

func NewLiteralExpressionSyntax(literalToken *SyntaxToken, value interface{}) *LiteralExpressionSyntax {
    return &LiteralExpressionSyntax{
        LiteralToken: literalToken,
        value: value,
    }
}
