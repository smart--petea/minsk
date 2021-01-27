package CodeAnalysis

import (
    "minsk/CodeAnalysis/Syntax"
    "minsk/CodeAnalysis/Binding"
    "minsk/Util"
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
    if len(c.Syntax.GetDiagnostics()) > 0 {
        return NewEvaluationResult(c.Syntax.GetDiagnostics(), nil)
    }

    binder := Binding.NewBinder()
    boundExpression := binder.BindExpression(c.Syntax.Root)
    if len(binder.GetDiagnostics()) > 0 {
        return NewEvaluationResult(binder.GetDiagnostics(), nil)
    }

    evaluator := NewEvaluator(boundExpression)
    value := evaluator.Evaluate()
    return NewEvaluationResult([]*Util.Diagnostic{}, value)
}

