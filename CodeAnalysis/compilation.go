package CodeAnalysis

import (
    "minsk/CodeAnalysis/Syntax"
    "minsk/CodeAnalysis/Binding"
)

type Compilation struct {
   Syntax *Syntax.SyntaxTree
}

func NewCompilation(syntax *Syntax.SyntaxTree) *Compilation {
    return &Compilation{
        Syntax: syntax,
    }
}

func (c *Compilation) Evaluate() *EvaluationResult {
    binder := Binding.NewBinder()
    boundExpression := binder.BindExpression(c.Syntax.Root)


    diagnostics := append(c.Syntax.GetDiagnostics(), binder.GetDiagnostics()...)
    if len(diagnostics) > 0 {
        return NewEvaluationResult(diagnostics, nil)
    }

    evaluator := NewEvaluator(boundExpression)
    value := evaluator.Evaluate()
    return NewEvaluationResult(diagnostics, value)
}

