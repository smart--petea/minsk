package CodeAnalysis

type EvaluationResult struct {
    Diagnostics []string
    Value interface{}
}

func NewEvaluationResult(diagnostics []string, value interface{}) *EvaluationResult {
    return &EvaluationResult{
        Diagnostics: diagnostics,
        Value: value,
    }
}
