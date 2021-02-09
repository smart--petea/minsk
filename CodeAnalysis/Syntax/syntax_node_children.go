package Syntax

type syntaxNodeChildren struct {
    children []SyntaxNode
}

func newSyntaxNodeChildren(children ...SyntaxNode) *syntaxNodeChildren {
    return &syntaxNodeChildren{
        children: children,
    }
}

func (s *syntaxNodeChildren) GetChildren() []SyntaxNode {
    return s.children
}
