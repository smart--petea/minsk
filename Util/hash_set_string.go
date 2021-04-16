package Util

type HashSetString map[string]bool

func NewHashSetString() *HashSetString {
    h := make(HashSetString)
    return &h
}

func (h *HashSetString) Add(s string) bool {
    v := *h
    if _, ok := v[s]; ok {
        return false
    }

    v[s] = true
    return true
}
