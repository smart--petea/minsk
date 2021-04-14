package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/Util"
)

type CompilationUnitSyntax struct {
    *Util.ChildrenProvider

    Members []MemberSyntax
    EndOfFileToken *SyntaxToken
}

func NewCompilationUnitSyntax(members []MemberSyntax, endOfFileToken *SyntaxToken) *CompilationUnitSyntax {
    return &CompilationUnitSyntax{
        ChildrenProvider: Util.NewChildrenProvider(members...),

        Members: members,
        EndOfFileToken: endOfFileToken,
    }
}

func (c *CompilationUnitSyntax) Kind() SyntaxKind.SyntaxKind {
    return SyntaxKind.CompilationUnit
}

func (c *CompilationUnitSyntax) Value() interface{} {
    return nil
}
