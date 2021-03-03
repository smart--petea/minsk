package Binding

type boundNodeChildren struct {
    children []BoundNode
}

func newBoundNodeChildren(children ...BoundNode) *boundNodeChildren {
    return &boundNodeChildren{
        children: children,
    }
}

func (b *boundNodeChildren) GetChildren() <-chan BoundNode {
    chanChildren := make(chan BoundNode)

    go func(){
        defer close(chanChildren)

        for _, child := range b.children {
            chanChildren <- child
        }

    }()

    return chanChildren
}
