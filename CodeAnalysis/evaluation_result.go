package CodeAnalysis

import (
    "minsk/Util"
)

type EvaluationResult struct {
    Diagnostics []*Util.Diagnostic
    Value interface{}
}

func NewEvaluationResult(diagnostics []*Util.Diagnostic, value interface{}) *EvaluationResult {
    return &EvaluationResult{
        Diagnostics: diagnostics,
        Value: value,
    }
}
