package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
)

type CompilationUnitSyntax struct {
    *syntaxNodeChildren

    statement StatementSyntax
    EndOfFileToken *SyntaxToken
}

func NewCompilationUnitSyntax(statement StatementSyntax, endOfFileToken *SyntaxToken) *CompilationUnitSyntax {
    return &CompilationUnitSyntax{
        syntaxNodeChildren: newSyntaxNodeChildren(expression),

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
