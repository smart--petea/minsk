package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
)

type SyntaxNode interface {
    Kind() SyntaxKind.SyntaxKind
    Value()  interface{}
    GetChildren() []SyntaxNode 
}
