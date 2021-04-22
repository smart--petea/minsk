package CodeAnalysis

import (
    "minsk/Util"
)

type EvaluationResult struct {
    Diagnostics Util.SliceDiagnostic
    Value interface{}
}

func NewEvaluationResult(diagnostics Util.SliceDiagnostic, value interface{}) *EvaluationResult {
    return &EvaluationResult{
        Diagnostics: diagnostics,
        Value: value,
    }
}
