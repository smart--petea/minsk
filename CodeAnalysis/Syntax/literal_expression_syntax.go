package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/CodeAnalysis/Text"
)

type LiteralExpressionSyntax struct {
    *syntaxNodeChildren

    LiteralToken *SyntaxToken
    value interface{}
}

func (l *LiteralExpressionSyntax) Value() interface{} {
    return l.value
}

func (l *LiteralExpressionSyntax) Kind() SyntaxKind.SyntaxKind {
    return SyntaxKind.LiteralExpression
}

func NewLiteralExpressionSyntax(literalToken *SyntaxToken, value interface{}) *LiteralExpressionSyntax {
    return &LiteralExpressionSyntax{
        syntaxNodeChildren: newSyntaxNodeChildren(literalToken),

        LiteralToken: literalToken,
        value: value,
    }
}

func (l *LiteralExpressionSyntax) GetSpan() *Text.TextSpan {
    return SyntaxNodeToTextSpan(l)
}
