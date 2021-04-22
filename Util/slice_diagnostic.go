package Util

import (
    "sort"
    "minsk/CodeAnalysis/Text"
)

type SliceDiagnostic []*Diagnostic 

func (s SliceDiagnostic) OrderBy(pluck func(*Diagnostic) *Text.TextSpan, comparer Text.ITextSpanComparer) SliceDiagnostic {
    sorted := make(SliceDiagnostic, 0, len(s))
    for _, d := range s {
        sorted = append(sorted, d)
    }

    sort.Slice(sorted, func(i, j int) bool { return comparer.Compare(pluck(sorted[i]), pluck(sorted[j])) > 0 })

    return sorted
}
