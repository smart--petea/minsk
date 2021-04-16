package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/CodeAnalysis/Text"
    "minsk/Util"
)

type ParameterSyntax struct {
    *Util.ChildrenProvider

    Identifier *SyntaxToken
    Ttype *TypeClauseSyntax
}

func NewParameterSyntax(identifier *SyntaxToken, ttype *TypeClauseSyntax) *ParameterSyntax {
    return &ParameterSyntax{
        ChildrenProvider: Util.NewChildrenProvider(identifier, ttype),

        Identifier: identifier,
        Ttype: ttype,
    }
}

func (g *ParameterSyntax) Kind() SyntaxKind.SyntaxKind {
    return SyntaxKind.Parameter
}

func (g *ParameterSyntax) Value() interface{} {
    return nil
}

func (g *ParameterSyntax) GetSpan() *Text.TextSpan {
    return SyntaxNodeToTextSpan(g)
}
