package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/CodeAnalysis/Text"
    "minsk/Util"
)

type ParenthesizedExpressionSyntax struct {
    *Util.ChildrenProvider

    OpenParenthesisToken *SyntaxToken
    Expression ExpressionSyntax
    CloseParenthesisToken *SyntaxToken
}

func NewParenthesizedExpressionSyntax(openParenthesisToken *SyntaxToken, expression ExpressionSyntax, closeParenthesisToken *SyntaxToken)  *ParenthesizedExpressionSyntax {
    return &ParenthesizedExpressionSyntax{
        ChildrenProvider: Util.NewChildrenProvider(openParenthesisToken, expression, closeParenthesisToken),

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

func (p *ParenthesizedExpressionSyntax) GetSpan() *Text.TextSpan {
    return SyntaxNodeToTextSpan(p)
}
