package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
)

type NameExpressionSyntax struct {
    *syntaxNodeChildren
    IdentifierToken *SyntaxToken
}

func (b *NameExpressionSyntax) Value() interface{} {
    return nil
}

func NewNameExpressionSyntax(identifierToken *SyntaxToken) *NameExpressionSyntax {
    return &NameExpressionSyntax{
        syntaxNodeChildren: newSyntaxNodeChildren(identifierToken),
        IdentifierToken: identifierToken,
    }
}

func (n *NameExpressionSyntax) Kind() SyntaxKind.SyntaxKind {
    return SyntaxKind.NameExpression 
}
