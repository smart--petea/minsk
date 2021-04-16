package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
)

type MemberSyntaxSlice []MemberSyntax

func (n MemberSyntaxSlice) OfType(kind SyntaxKind.SyntaxKind) MemberSyntaxSlice {
    var result MemberSyntaxSlice
    for _, node := range n {
        if node.Kind() == kind {
            result = append(result, node)
        }
    }

    return result
}

func (n MemberSyntaxSlice) Single() MemberSyntax {
    if n == nil {
        var p MemberSyntax
        return p
    }

    return n[0]
}

func (n MemberSyntaxSlice) ToEmptyInterfaceSlice() []interface{} {
    var c []interface{}
    for _, i := range n {
        c = append(c, i)
    }

    return c
}
