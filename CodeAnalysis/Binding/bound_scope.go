package Binding

import (
    "minsk/CodeAnalysis/Symbols"
)

type BoundScope struct {
    variables map[Symbols.IVariableSymbol]interface{}
    functions map[*Symbols.FunctionSymbol]interface{}

    Parent *BoundScope
}

func NewBoundScope(parent *BoundScope) *BoundScope {
    return &BoundScope{
        Parent: parent,
    }
}

func (bs *BoundScope) TryLookupVariable(name string, out *Symbols.IVariableSymbol) bool {
    var variable Symbols.IVariableSymbol

    for variable, _ = range bs.variables {
        if variable.GetName() == name {
            *out = variable
            return true
        }
    }

    if bs.Parent == nil {
        return false
    }

    return bs.Parent.TryLookupVariable(name, out)
}

func (bs *BoundScope) TryDeclareVariable(variable Symbols.IVariableSymbol) bool {
    if bs.variables == nil {
        bs.variables = make(map[Symbols.IVariableSymbol]interface{})
    }

    var tmp Symbols.IVariableSymbol
    for tmp, _ = range bs.variables {
        if variable.GetName() == tmp.GetName() {
            return false
        }
    }


    bs.variables[variable] = nil
    return true
}

func (bs *BoundScope) TryLookupFunction(name string, out **Symbols.FunctionSymbol) bool {
    var function *Symbols.FunctionSymbol

    for function, _ = range bs.functions {
        if function.Name == name {
            *out = function
            return true
        }
    }

    if bs.Parent == nil {
        return false
    }

    return bs.Parent.TryLookupFunction(name, out)
}

func (bs *BoundScope) TryDeclareFunction(function *Symbols.FunctionSymbol) bool {
    if bs.functions == nil {
        bs.functions = make(map[*Symbols.FunctionSymbol]interface{})
    }

    var tmp *Symbols.FunctionSymbol
    for tmp, _ = range bs.functions {
        if function.Name == tmp.Name {
            return false
        }
    }


    bs.functions[function] = nil
    return true
}

func (bs *BoundScope) GetDeclaredVariables() []Symbols.IVariableSymbol {
    var d []Symbols.IVariableSymbol
    for variable, _ := range bs.variables {
        d = append(d, variable)
    }

    return d
}

func (bs *BoundScope) GetDeclaredFunctions() []*Symbols.FunctionSymbol {
    var d []*Symbols.FunctionSymbol
    for function, _ := range bs.functions {
        d = append(d, function)
    }

    return d
}
