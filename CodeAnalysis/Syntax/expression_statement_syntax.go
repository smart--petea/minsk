package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/CodeAnalysis/Text"
)

type ExpressionStatementSyntax struct {
    *syntaxNodeChildren

    Expression ExpressionSyntax
}

func NewExpressionStatementSyntax(expression ExpressionSyntax) *ExpressionStatementSyntax {
    return &ExpressionStatementSyntax{
        syntaxNodeChildren: newSyntaxNodeChildren(expression),

        Expression: expression,
    }
}

func (e *ExpressionStatementSyntax) Kind() SyntaxKind.SyntaxKind {
    return SyntaxKind.ExpressionStatement
}

func (e *ExpressionStatementSyntax) Value() interface{} {
    return e.Expression.Value()
}

func (e *ExpressionStatementSyntax) GetSpan() *Text.TextSpan {
    return SyntaxNodeToTextSpan(e)
}
