package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
)

type CompilationUnitSyntax struct {
    *syntaxNodeChildren

    Statement StatementSyntax
    EndOfFileToken *SyntaxToken
}

func NewCompilationUnitSyntax(statement StatementSyntax, endOfFileToken *SyntaxToken) *CompilationUnitSyntax {
    return &CompilationUnitSyntax{
        syntaxNodeChildren: newSyntaxNodeChildren(statement.(SyntaxNode)),

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
