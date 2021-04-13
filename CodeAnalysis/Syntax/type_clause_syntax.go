package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/CodeAnalysis/Text"
    "minsk/Util"
)

type TypeClauseSyntax struct {
    *Util.ChildrenProvider

    ColonToken *SyntaxToken
    Identifier *SyntaxToken
}

func NewTypeClauseSyntax(colonToken *SyntaxToken, identifier *SyntaxToken) *TypeClauseSyntax {
    return &TypeClauseSyntax{
        ChildrenProvider: Util.NewChildrenProvider(colonToken, identifier),

        ColonToken: colonToken,
        Identifier: identifier,
    }
}

func (e *TypeClauseSyntax) Kind() SyntaxKind.SyntaxKind {
    return SyntaxKind.TypeClause
}

func (e *TypeClauseSyntax) Value() interface{} {
    return nil
}

func (e *TypeClauseSyntax) GetSpan() *Text.TextSpan {
    return SyntaxNodeToTextSpan(e)
}
