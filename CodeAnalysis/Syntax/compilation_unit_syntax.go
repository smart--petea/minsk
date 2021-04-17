package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/Util"
)

type CompilationUnitSyntax struct {
    *Util.ChildrenProvider

    Members MemberSyntaxSlice
    EndOfFileToken *SyntaxToken
}

func NewCompilationUnitSyntax(members MemberSyntaxSlice, endOfFileToken *SyntaxToken) *CompilationUnitSyntax {
    children := members.ToEmptyInterfaceSlice()
    children = append(children, endOfFileToken)

    return &CompilationUnitSyntax{
        ChildrenProvider: Util.NewChildrenProvider(children...),

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
