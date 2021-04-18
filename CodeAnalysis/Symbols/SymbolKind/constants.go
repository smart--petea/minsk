package SymbolKind

type SymbolKind string

const (
    LocalVariable SymbolKind = "LocalVariable"
    GlobalVariable SymbolKind = "GlobalVariable"
    Variable SymbolKind = "Variable"
    Type SymbolKind = "Type"
    Function SymbolKind = "Function"
    Parameter SymbolKind = "Parameter"
)
