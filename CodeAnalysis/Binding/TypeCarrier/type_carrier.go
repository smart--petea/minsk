package TypeCarrier

type TypeCarrier interface{}

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
