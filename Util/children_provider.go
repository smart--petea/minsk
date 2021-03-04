package Util

type ChildrenProvider struct {
    children []interface{}
}

func NewChildrenProvider(children ...interface{}) *ChildrenProvider {
    return &ChildrenProvider{
        children: children,
    }
}

func (c *ChildrenProvider) GetChildren() <-chan interface{} {
    chanChildren := make(chan interface{})

    go func(){
        defer close(chanChildren)

        for _, child := range c.children {
            chanChildren <- child
        }

    }()

    return chanChildren
}
