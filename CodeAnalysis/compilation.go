package CodeAnalysis

import (
    "minsk/CodeAnalysis/Syntax"
    "minsk/CodeAnalysis/Binding"
    "minsk/Util"
    "sync/atomic"
)

type Compilation struct {
   SyntaxTree *Syntax.SyntaxTree

   globalScope *Binding.BoundGlobalScope
   globalScopeOnce sync.Once
}

func NewCompilation(syntaxTree *Syntax.SyntaxTree) *Compilation {
    return &Compilation{
        SyntaxTree: syntaxTree,
    }
}

func (c *Compilation) GlobalScope() *Binding.BoundGlobalScope {
    if c.globalScope == nil {
        c.globalScopeOnce.Do(func() {
            c.globalScope = Binder.BoundGlobalScopeFromCompilationUnitSyntax(c.SyntaxTree.Root)
        })
    }

    return c.globalScope
}

func (c *Compilation) Evaluate(variables map[*Util.VariableSymbol]interface{}) *EvaluationResult {
    if len(c.SyntaxTree.GetDiagnostics()) > 0 {
        return NewEvaluationResult(c.SyntaxTree.GetDiagnostics(), nil)
    }

    if len(c.GlobalScope().GetDiagnostics()) > 0 {
        return NewEvaluationResult(c.GlobalScope().GetDiagnostics(), nil)
    }

    evaluator := NewEvaluator(c.GlobalScope().Expression, variables)
    value := evaluator.Evaluate()
    return NewEvaluationResult([]*Util.Diagnostic{}, value)
}

