package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/CodeAnalysis/Text"
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

func (e *ElseClauseSyntax) GetSpan() *Text.TextSpan {
    return SyntaxNodeToTextSpan(e)
}
