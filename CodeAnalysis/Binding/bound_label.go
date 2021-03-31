package Binding

type BoundLabel struct {
    Name string
}

func NewBoundLabel(name string) *BoundLabel {
    return &BoundLabel{
        Name: name,
    }
}
