package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/CodeAnalysis/Text"
)

type AssignmentExpressionSyntax struct {
    *syntaxNodeChildren

    IdentifierToken *SyntaxToken
    EqualsToken *SyntaxToken
    Expression ExpressionSyntax
}

func (a *AssignmentExpressionSyntax) Value() interface{} {
    return nil
}

func NewAssignmentExpressionSyntax(identifierToken *SyntaxToken, equalsToken *SyntaxToken, expression ExpressionSyntax) *AssignmentExpressionSyntax {
    return &AssignmentExpressionSyntax{
        syntaxNodeChildren: newSyntaxNodeChildren(identifierToken, equalsToken, expression),

        IdentifierToken: identifierToken,
        EqualsToken: equalsToken,
        Expression: expression,
    }
}

func (a *AssignmentExpressionSyntax) Kind() SyntaxKind.SyntaxKind {
    return SyntaxKind.AssignmentExpression 
}

func (a *AssignmentExpressionSyntax) GetSpan() *Text.TextSpan {
    return SyntaxNodeToTextSpan(a)
}
