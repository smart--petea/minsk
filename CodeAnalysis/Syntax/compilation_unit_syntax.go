package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/Util"
)

type CompilationUnitSyntax struct {
    *Util.ChildrenProvider

    Statement StatementSyntax
    EndOfFileToken *SyntaxToken
}

func NewCompilationUnitSyntax(statement StatementSyntax, endOfFileToken *SyntaxToken) *CompilationUnitSyntax {
    return &CompilationUnitSyntax{
        ChildrenProvider: Util.NewChildrenProvider(statement),

        Statement: statement,
        EndOfFileToken: endOfFileToken,
    }
}

func (c *CompilationUnitSyntax) Kind() SyntaxKind.SyntaxKind {
    return SyntaxKind.CompilationUnit
}

func (c *CompilationUnitSyntax) Value() interface{} {
    return c.Statement.Value()
}
