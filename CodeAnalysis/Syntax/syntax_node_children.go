package Syntax

type syntaxNodeChildren struct {
    children []SyntaxNode
}

func newSyntaxNodeChildren(children ...SyntaxNode) *syntaxNodeChildren {
    return &syntaxNodeChildren{
        children: children,
    }
}

func (s *syntaxNodeChildren) GetChildren() <-chan SyntaxNode {
    chanChildren := make(chan SyntaxNode)

    go func(){
        defer close(chanChildren)

        for _, child := range s.children {
            chanChildren <- child
        }

    }()

    return chanChildren
}
