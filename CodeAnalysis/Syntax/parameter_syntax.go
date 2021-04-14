package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/Util"
)

type ParamterSyntax struct {
    *Util.ChildrenProvider

    Identifier *SyntaxToken
    Ttype *TypeClauseSyntax
}

func NewParameterSyntax(identifier *SyntaxToken, ttype *TypeClauseSyntax) *ParamterSyntax {
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
