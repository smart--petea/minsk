package Util

import (
    "reflect"
)

type VariableSymbol struct {
    Name string
    Type reflect.Kind
}

func NewVariableSymbol(name string, kind reflect.Kind) *VariableSymbol {
    return &VariableSymbol{
        Name: name,
        Type: kind,
    }
}
