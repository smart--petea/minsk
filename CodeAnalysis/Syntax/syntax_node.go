package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/Util"
)

type SyntaxNode interface {
    Kind() SyntaxKind.SyntaxKind
    Value()  interface{}
    GetChildren() <-chan SyntaxNode 
}

func SyntaxNodeToTextSpan(sn SyntaxNode) *Util.TextSpan {
    if syntaxToken, ok := sn.(*SyntaxToken); ok {
        return Util.NewTextSpan(syntaxToken.Position, len(syntaxToken.Runes)) 
    }

    children := sn.GetChildren()
    first := <-children
    last := first
    for last = range children {}

    start := SyntaxNodeToTextSpan(first).Start
    end := SyntaxNodeToTextSpan(last).End()
    return Util.NewTextSpan(start, end) 
}
