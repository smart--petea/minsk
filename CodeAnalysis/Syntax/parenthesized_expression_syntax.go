package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
)

type ParenthesizedExpressionSyntax struct {
    *syntaxNodeChildren
    OpenParenthesisToken *SyntaxToken
    Expression ExpressionSyntax
    CloseParenthesisToken *SyntaxToken
}

func NewParenthesizedExpressionSyntax(openParenthesisToken *SyntaxToken, expression ExpressionSyntax, closeParenthesisToken *SyntaxToken)  *ParenthesizedExpressionSyntax {
    return &ParenthesizedExpressionSyntax{
        syntaxNodeChildren: newSyntaxNodeChildren(openParenthesisToken, expression, closeParenthesisToken),
        OpenParenthesisToken: openParenthesisToken,
        Expression: expression,
        CloseParenthesisToken: closeParenthesisToken,
    }
}

func (p *ParenthesizedExpressionSyntax) Kind() SyntaxKind.SyntaxKind {
    return SyntaxKind.ParenthesizedExpression
}

func (p *ParenthesizedExpressionSyntax) Value() interface{} {
    return nil
}
