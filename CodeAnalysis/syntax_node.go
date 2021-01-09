package CodeAnalysis

import (
    "minsk/CodeAnalysis/SyntaxKind"
)

type SyntaxNode interface {
    Kind() SyntaxKind.SyntaxKind
    Value()  interface{}
    GetChildren() []SyntaxNode 
}
