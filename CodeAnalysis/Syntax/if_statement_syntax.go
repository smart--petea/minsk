package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
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
        syntaxNodeChildren: newSyntaxNodeChildren(ifKeyword, condition thenStatement, elseClauseSyntax),

        IfKeyword: ifKeyword,
        Condition: condition,
        ThenStatement: thenStatement,
        ElseClause: elseClause,
    }
}

func (e *IfStatementSyntax) Kind() SyntaxKind.SyntaxKind {
    return SyntaxKind.IfStatement
}

func (e *IfStatementSyntax) Value() interface{} {
    return nil
}
