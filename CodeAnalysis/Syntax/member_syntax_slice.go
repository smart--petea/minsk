package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
)

type MemberSyntaxSlice []MemberSyntax

func (n MemberSyntaxSlice) OfType(kind SyntaxKind.SyntaxKind) MemberSyntaxSlice {
    var result MemberSyntaxSlice
    for _, node := range n {
        if n.Kind() == kind {
            result = append(result, node)
        }
    }

    return result
}

func (n MemberSyntaxSlice) Single() MemberSyntax {
    if n == nil {
        return n
    }

    return n[0]
}
