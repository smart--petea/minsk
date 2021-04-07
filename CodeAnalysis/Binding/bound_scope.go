package Binding

import (
    "minsk/CodeAnalysis/Symbols"
)

type BoundScope struct {
    variables map[*Symbols.VariableSymbol]interface{}
    functions map[*Symbols.FunctionSymbol]interface{}

    Parent *BoundScope
}

func NewBoundScope(parent *BoundScope) *BoundScope {
    return &BoundScope{
        Parent: parent,
    }
}

func (bs *BoundScope) TryLookupVariable(name string, out **Symbols.VariableSymbol) bool {
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

    return bs.Parent.TryLookupVariable(name, out)
}

func (bs *BoundScope) TryDeclareVariable(variable *Symbols.VariableSymbol) bool {
    if bs.variables == nil {
        bs.variables = make(map[*Symbols.VariableSymbol]interface{})
    }

    var tmp *Symbols.VariableSymbol
    for tmp, _ = range bs.variables {
        if variable.Name == tmp.Name {
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

func (bs *BoundScope) GetDeclaredVariables() []*Symbols.VariableSymbol {
    var d []*Symbols.VariableSymbol
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
