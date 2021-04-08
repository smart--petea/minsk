package Conversion

import (
    "minsk/CodeAnalysis/Symbols"
)

type Conversion struct {
    Exists bool 
    IsIdentity bool
    IsImplicit bool
    IsExplicit bool
}

func newConversion(exists, isIdentity, isImplicit bool) *Conversion {
    return &Conversion{
        Exists: exists ,
        IsIdentity: isIdentity,
        IsImplicit: isImplicit,
        IsExplicit: exists && !isImplicit,
    }
}

var (
    None *Conversion = newConversion(false, false, false)
    Identity *Conversion = newConversion(true, true, true)
    Implicit *Conversion = newConversion(true, false, true)
    Explicit *Conversion = newConversion(true, false, false)
)

func ConversionClassify(from, to *Symbols.TypeSymbol) *Conversion {
    if from == to {
        return Identity
    }

    if from == Symbols.TypeSymbolBool || from == Symbols.TypeSymbolInt {
        if to == Symbols.TypeSymbolString {
            return Explicit
        }
    }

    if from == Symbols.TypeSymbolString {
        if to == Symbols.TypeSymbolBool || to == Symbols.TypeSymbolInt {
            return Explicit
        }
    }

    return None
}
