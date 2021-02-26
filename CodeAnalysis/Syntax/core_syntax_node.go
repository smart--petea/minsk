package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
)

type CoreSyntaxNode interface {
    Kind() SyntaxKind.SyntaxKind
    Value()  interface{}
    GetChildren() <-chan CoreSyntaxNode 
}
