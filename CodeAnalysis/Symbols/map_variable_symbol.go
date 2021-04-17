package Symbols

type MapVariableSymbol map[*VariableSymbol]interface{}

func NewMapVariableSymbol() MapVariableSymbol {
    return make(MapVariableSymbol)
}
