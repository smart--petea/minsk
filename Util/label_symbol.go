package Util

type LabelSymbol struct {
    Name string
}

func NewLabelSymbol(name string) *LabelSymbol {
    return &VariableSymbol{
        Name: name,
    }
}
