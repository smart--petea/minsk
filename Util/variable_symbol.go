package Util

import (
    "reflect"
)

type VariableSymbol struct {
    Name string
    Type reflect.Kind
    IsReadOnly bool
}

func NewVariableSymbol(name string, isReadOnly bool, kind reflect.Kind) *VariableSymbol {
    return &VariableSymbol{
        Name: name,
        Type: kind,
        IsReadOnly: isReadOnly,
    }
}
