package CodeAnalysis

import (
    "minsk/CodeAnalysis/Syntax"
    "minsk/CodeAnalysis/Binding"
    "minsk/Util"
)

type Compilation struct {
   SyntaxTree *Syntax.SyntaxTree
}

func NewCompilation(syntaxTree *Syntax.SyntaxTree) *Compilation {
    return &Compilation{
        SyntaxTree: syntaxTree,
    }
}

func (c *Compilation) Evaluate(variables map[*Util.VariableSymbol]interface{}) *EvaluationResult {
    if len(c.SyntaxTree.GetDiagnostics()) > 0 {
        return NewEvaluationResult(c.SyntaxTree.GetDiagnostics(), nil)
    }

    globalScope := Binding.BoundGlobalScopeFromCompilationUnitSyntax(c.SyntaxTree.Root)
    if len(globalScope.GetDiagnostics()) > 0 {
        return NewEvaluationResult(globalScope.GetDiagnostics(), nil)
    }

    evaluator := NewEvaluator(globalScope.Expression, variables)
    value := evaluator.Evaluate()
    return NewEvaluationResult([]*Util.Diagnostic{}, value)
}

