package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
)

type CompilationUnitSyntax struct {
    *syntaxNodeChildren
    Expression ExpressionSyntax
    EndOfFileToken *SyntaxToken
}

func NewCompilationUnitSyntax(expression ExpressionSyntax, endOfFileToken *SyntaxToken) *CompilationUnitSyntax {
    return &CompilationUnitSyntax{
        syntaxNodeChildren: newSyntaxNodeChildren(expression),
        Expression: expression,
        EndOfFileToken: endOfFileToken,
    }
}

func (c *CompilationUnitSyntax) Kind() SyntaxKind.SyntaxKind {
    return SyntaxKind.CompilationUnit
}

func (c *CompilationUnitSyntax) Value() interface{} {
    return c.Expression.Value()
}
