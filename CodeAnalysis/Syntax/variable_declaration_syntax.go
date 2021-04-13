package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/CodeAnalysis/Text"
    "minsk/Util"
)

type VariableDeclarationSyntax struct {
    *Util.ChildrenProvider

    Keyword *SyntaxToken
    Identifier *SyntaxToken
    TypeClause *TypeClauseSyntax
    EqualsToken *SyntaxToken
    Initializer ExpressionSyntax
}

func (v *VariableDeclarationSyntax) Kind() SyntaxKind.SyntaxKind {
    return SyntaxKind.VariableDeclaration
}

func NewVariableDeclarationSyntax(keyword, identifier *SyntaxToken, typeClause *TypeClauseSyntax, equalsToken  *SyntaxToken, initializer ExpressionSyntax) *VariableDeclarationSyntax {
    return &VariableDeclarationSyntax{
        ChildrenProvider: Util.NewChildrenProvider(keyword, identifier, typeClause, equalsToken, initializer),

        Keyword: keyword,
        Identifier: identifier,
        TypeClause: typeClause,
        EqualsToken: equalsToken,
        Initializer: initializer,
    }
}

func (v *VariableDeclarationSyntax) Value() interface{} {
    return nil
}

func (v *VariableDeclarationSyntax) GetSpan() *Text.TextSpan {
    return SyntaxNodeToTextSpan(v)
}
