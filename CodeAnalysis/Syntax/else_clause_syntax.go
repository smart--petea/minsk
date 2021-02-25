package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
)

type ElseClauseSyntax struct {
    *syntaxNodeChildren

    ElseKeyword *SyntaxToken
    ElseStatement StatementSyntax
}

func NewElseClauseSyntax(elseKeyword *SyntaxToken, elseStatement StatementSyntax) *ElseClauseSyntax {
    return &ElseClauseSyntax{
        syntaxNodeChildren: newSyntaxNodeChildren(elseKeyword, elseStatement),

        ElseKeyword: elseKeyword,
        ElseStatement: elseStatement,
    }
}

func (e *ElseClauseSyntax) Kind() SyntaxKind.SyntaxKind {
    return SyntaxKind.ElseClause
}

func (e *ElseClauseSyntax) Value() interface{} {
    return nil
}
