package TypeCarrier

type TypeCarrier interface{}

func Bool() TypeCarrier {
    var x bool
    return TypeCarrier(x)
}

func Int() TypeCarrier {
    var x int
    return TypeCarrier(x)
}

func Same(leftTypeCarrier TypeCarrier, rightTypeCarrier TypeCarrier) bool {
    return (IsBool(leftTypeCarrier) && IsBool(rightTypeCarrier)) ||
        (IsInt(leftTypeCarrier) && IsInt(rightTypeCarrier))
}

func IsInt(val TypeCarrier) bool {
    switch val.(type) {
    case int, int32, int64:
        return true
    default:
        return false
    }
}

func IsBool(val TypeCarrier) bool {
    switch val.(type) {
    case bool:
        return true
    default:
        return false
    }
}
