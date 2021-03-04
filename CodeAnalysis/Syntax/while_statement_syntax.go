package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/CodeAnalysis/Text"
    "minsk/Util"
)

type WhileStatementSyntax struct {
    *Util.ChildrenProvider

    WhileKeyword *SyntaxToken
    Condition ExpressionSyntax
    Body StatementSyntax
}

func NewWhileStatementSyntax(whileKeyword *SyntaxToken, condition ExpressionSyntax, body StatementSyntax) *WhileStatementSyntax {
    return &WhileStatementSyntax{
        ChildrenProvider: Util.NewChildrenProvider(),

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
