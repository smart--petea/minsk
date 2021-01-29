package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
)

type NameExpressionSyntax struct {
    IdentifierToken *SyntaxToken
}

func (b *NameExpressionSyntax) Value() interface{} {
    return nil
}

func NewNameExpressionSyntax(identifierToken *SyntaxToken) *NameExpressionSyntax {
    return &NameExpressionSyntax{
        IdentifierToken: identifierToken,
    }
}

func (n *NameExpressionSyntax) Kind() SyntaxKind.SyntaxKind {
    return SyntaxKind.NameExpression 
}

func (n *NameExpressionSyntax) GetChildren() []SyntaxNode {
    return []SyntaxNode{
        n.IdentifierToken,
    }
}
