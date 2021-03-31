package Binding

import (
    "minsk/CodeAnalysis/Symbols"
)

type BoundScope struct {
    variables map[*Symbols.VariableSymbol]interface{}
    Parent *BoundScope
}

func NewBoundScope(parent *BoundScope) *BoundScope {
    return &BoundScope{
        variables: make(map[*Symbols.VariableSymbol]interface{}),
        Parent: parent,
    }
}

func (bs *BoundScope) TryLookup(name string, out **Symbols.VariableSymbol) bool {
    var variable *Symbols.VariableSymbol

    for variable, _ = range bs.variables {
        if variable.Name == name {
            *out = variable
            return true
        }
    }

    if bs.Parent == nil {
        return false
    }

    return bs.Parent.TryLookup(name, out)
}

func (bs *BoundScope) TryDeclare(variable *Symbols.VariableSymbol) bool {
    var tmp *Symbols.VariableSymbol

    for tmp, _ = range bs.variables {
        if variable.Name == tmp.Name {
            return false
        }
    }


    bs.variables[variable] = nil
    return true
}

func (bs *BoundScope) GetDeclaredVariables() []*Symbols.VariableSymbol {
    var d []*Symbols.VariableSymbol
    for variable, _ := range bs.variables {
        d = append(d, variable)
    }

    return d
}
