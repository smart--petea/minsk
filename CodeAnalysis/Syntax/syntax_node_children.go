package Syntax

type syntaxNodeChildren struct {
    children []CoreSyntaxNode
}

func newSyntaxNodeChildren(children ...CoreSyntaxNode) *syntaxNodeChildren {
    return &syntaxNodeChildren{
        children: children,
    }
}

func (s *syntaxNodeChildren) GetChildren() <-chan CoreSyntaxNode {
    chanChildren := make(chan CoreSyntaxNode)

    go func(){
        defer close(chanChildren)

        for _, child := range s.children {
            chanChildren <- child
        }

    }()

    return chanChildren
}
