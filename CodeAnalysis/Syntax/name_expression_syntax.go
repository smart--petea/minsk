package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/CodeAnalysis/Text"
    "minsk/Util"
)

type NameExpressionSyntax struct {
    *Util.ChildrenProvider

    IdentifierToken *SyntaxToken
}

func (b *NameExpressionSyntax) Value() interface{} {
    return nil
}

func NewNameExpressionSyntax(identifierToken *SyntaxToken) *NameExpressionSyntax {
    return &NameExpressionSyntax{
        ChildrenProvider: Util.NewChildrenProvider(identifierToken),
        IdentifierToken: identifierToken,
    }
}

func (n *NameExpressionSyntax) Kind() SyntaxKind.SyntaxKind {
    return SyntaxKind.NameExpression 
}

func (n *NameExpressionSyntax) GetSpan() *Text.TextSpan {
    return SyntaxNodeToTextSpan(n)
}
