package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
)

type AssignmentExpressionSyntax struct {
    IdentifierToken *SyntaxToken
    EqualsToken *SyntaxToken
    Expression ExpressionSyntax
}

func (a *AssignmentExpressionSyntax) Value() interface{} {
    return nil
}

func NewAssignmentExpressionSyntax(identifierToken *SyntaxToken, equalsToken *SyntaxToken, expression ExpressionSyntax) *AssignmentExpressionSyntax {
    return &AssignmentExpressionSyntax{
        IdentifierToken: identifierToken,
        EqualsToken: equalsToken,
        Expression: expression,
    }
}

func (a *AssignmentExpressionSyntax) Kind() SyntaxKind.SyntaxKind {
    return SyntaxKind.AssignmentExpression 
}

func (a *AssignmentExpressionSyntax) GetChildren() []SyntaxNode {
    return []SyntaxNode{
        a.IdentifierToken,
        a.EqualsToken,
        a.Expression,
    }
}
