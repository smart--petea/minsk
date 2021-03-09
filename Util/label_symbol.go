package Util

type LabelSymbol struct {
    Name string
}

func NewLabelSymbol(name string) *LabelSymbol {
    return &LabelSymbol{
        Name: name,
    }
}
