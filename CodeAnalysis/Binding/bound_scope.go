package Binding

import (
    "minsk/Util"
)

type BoundScope struct {
    variables map[*Util.VariableSymbol]interface{}
    Parent *BoundScope
}

func NewBoundScope(parent *BoundScope) *BoundScope {
    return &BoundScope{
        variables: make(map[*Util.VariableSymbol]interface{}),
        Parent: parent,
    }
}

func (bs *BoundScope) TryLookup(name string, out **Util.VariableSymbol) bool {
    var variable *Util.VariableSymbol

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

func (bs *BoundScope) TryDeclare(variable *Util.VariableSymbol) bool {
    var tmp *Util.VariableSymbol

    for tmp, _ = range bs.variables {
        if variable.Name == tmp.Name {
            return false
        }
    }


    bs.variables[variable] = nil
    return true
}

func (bs *BoundScope) GetDeclaredVariables() []*Util.VariableSymbol {
    var d []*Util.VariableSymbol
    for variable, _ := range bs.variables {
        d = append(d, variable)
    }

    return d
}
