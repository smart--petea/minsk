package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/CodeAnalysis/Text"
    "minsk/Util"
)

type ElseClauseSyntax struct {
    *Util.ChildrenProvider

    ElseKeyword *SyntaxToken
    ElseStatement StatementSyntax
}

func NewElseClauseSyntax(elseKeyword *SyntaxToken, elseStatement StatementSyntax) *ElseClauseSyntax {
    return &ElseClauseSyntax{
        ChildrenProvider: Util.NewChildrenProvider(elseKeyword, elseStatement),

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
