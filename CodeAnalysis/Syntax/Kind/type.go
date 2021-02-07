package Kind

type SyntaxKind string
type SyntaxKindSlice []SyntaxKind

func (sk SyntaxKindSlice) Unique() SyntaxKindSlice {
    uniqueMap := make(map[SyntaxKind]bool)
    var result SyntaxKindSlice

    for _, kind := range sk {
        if _, ok := uniqueMap[kind]; !ok {
            uniqueMap[kind] = true
            result = append(result, kind) 
        }
    }

    return result
}

func (sk SyntaxKindSlice) ExceptWith(exceptWith SyntaxKindSlice) SyntaxKindSlice {
    skMap := make(map[SyntaxKind]bool)
    for _, kind := range sk {
        skMap[kind] = true
    }

    for _, kind := range exceptWith {
        if _, ok := skMap[kind]; ok { 
            delete(skMap, kind)
        }
    }

    var result SyntaxKindSlice
    for kind, _ := range skMap {
        result = append(result, kind)
    }

    return result
}
