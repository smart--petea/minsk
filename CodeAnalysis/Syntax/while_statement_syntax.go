package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/CodeAnalysis/Text"
)

type WhileStatementSyntax struct {
    *syntaxNodeChildren

    WhileKeyword *SyntaxToken
    Condition ExpressionSyntax
    Body StatementSyntax
}

func NewWhileStatementSyntax(whileKeyword *SyntaxToken, condition ExpressionSyntax, body StatementSyntax) *WhileStatementSyntax {
    return &WhileStatementSyntax{
        syntaxNodeChildren: newSyntaxNodeChildren(),

        WhileKeyword: whileKeyword,
        Condition: condition,
        Body: body,
    }
}

func (w *WhileStatementSyntax) Kind() SyntaxKind.SyntaxKind {
    return SyntaxKind.WhileStatement
}

func (w *WhileStatementSyntax) Value() interface{} {
    return nil
}

func (w *WhileStatementSyntax) GetSpan() *Text.TextSpan {
    return SyntaxNodeToTextSpan(w)
}
