package CodeAnalysis

type SyntaxNode interface {
    Kind() SyntaxKind
    Value()  interface{}
    GetChildren() []SyntaxNode 
}
