package Symbols

type MapVariableSymbol map[IVariableSymbol]interface{}

func NewMapVariableSymbol() MapVariableSymbol {
    return make(MapVariableSymbol)
}

func (n MapVariableSymbol) Add(symbol IVariableSymbol, value interface{}) {
    n[symbol] = value
}
