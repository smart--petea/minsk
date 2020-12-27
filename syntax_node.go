package minsk

type SyntaxNode interface {
    Kind() SyntaxKind
    Value()  interface{}
    GetChildren() []SyntaxNode 
}
