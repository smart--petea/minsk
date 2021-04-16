package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/CodeAnalysis/Text"
    "minsk/Util"
)

type GlobalStatementSyntax struct {
    *Util.ChildrenProvider

    Statement StatementSyntax
}

func NewGlobalStatementSyntax(statement StatementSyntax) *GlobalStatementSyntax {
    return &GlobalStatementSyntax{
        ChildrenProvider: Util.NewChildrenProvider(statement),

        Statement: statement,
    }
}

func (g *GlobalStatementSyntax) Kind() SyntaxKind.SyntaxKind {
    return SyntaxKind.GlobalStatement
}

func (g *GlobalStatementSyntax) Value() interface{} {
    return g.Statement.Value()
}

func (g *GlobalStatementSyntax) GetSpan() *Text.TextSpan {
    return SyntaxNodeToTextSpan(g)
}
