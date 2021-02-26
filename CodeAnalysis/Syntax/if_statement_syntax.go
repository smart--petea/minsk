package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/CodeAnalysis/Text"
)

type IfStatementSyntax struct {
    *syntaxNodeChildren

    IfKeyword *SyntaxToken
    Condition ExpressionSyntax
    ThenStatement StatementSyntax
    ElseClause *ElseClauseSyntax
}

func NewIfStatementSyntax(ifKeyword *SyntaxToken, condition ExpressionSyntax, thenStatement StatementSyntax, elseClause *ElseClauseSyntax) *IfStatementSyntax {
    return &IfStatementSyntax{
        syntaxNodeChildren: newSyntaxNodeChildren(ifKeyword, condition, thenStatement, elseClause),

        IfKeyword: ifKeyword,
        Condition: condition,
        ThenStatement: thenStatement,
        ElseClause: elseClause,
    }
}

func (i *IfStatementSyntax) Kind() SyntaxKind.SyntaxKind {
    return SyntaxKind.IfStatement
}

func (i *IfStatementSyntax) Value() interface{} {
    return nil
}

func (i *IfStatementSyntax) GetSpan() *Text.TextSpan {
    return SyntaxNodeToTextSpan(i)
}
