package CodeAnalysis

import (
    "minsk/CodeAnalysis/Syntax"
    "minsk/CodeAnalysis/Binding"
    "minsk/CodeAnalysis/Lowering"
    "minsk/Util"

    "sync"
    "io"
)

type Compilation struct {
   SyntaxTree *Syntax.SyntaxTree

   Previous *Compilation

   globalScope *Binding.BoundGlobalScope
   globalScopeOnce sync.Once
}

func NewCompilation(syntaxTree *Syntax.SyntaxTree) *Compilation {
    return newCompilation(nil, syntaxTree)
}

func newCompilation(previous *Compilation, syntaxTree *Syntax.SyntaxTree) *Compilation {
    return &Compilation{
        SyntaxTree: syntaxTree,
        Previous: previous,
    }
}

func (c *Compilation) GlobalScope() *Binding.BoundGlobalScope {
    if c.globalScope == nil {
        c.globalScopeOnce.Do(func() {
            var previousGlobalScope *Binding.BoundGlobalScope
            if c.Previous != nil {
                previousGlobalScope = c.Previous.GlobalScope()
            }

            c.globalScope = Binding.BoundGlobalScopeFromCompilationUnitSyntax(previousGlobalScope, c.SyntaxTree.Root)
        })
    }

    return c.globalScope
}

func (c *Compilation) ContinueWith(syntaxTree *Syntax.SyntaxTree) *Compilation {
    previous := c
    return newCompilation(previous, syntaxTree)
}

func (c *Compilation) Evaluate(variables map[*Util.VariableSymbol]interface{}) *EvaluationResult {
    if len(c.SyntaxTree.GetDiagnostics()) > 0 {
        return NewEvaluationResult(c.SyntaxTree.GetDiagnostics(), nil)
    }

    if len(c.GlobalScope().GetDiagnostics()) > 0 {
        return NewEvaluationResult(c.GlobalScope().GetDiagnostics(), nil)
    }

    statement := c.GetStatement()
    evaluator := NewEvaluator(statement, variables)
    value := evaluator.Evaluate()
    return NewEvaluationResult([]*Util.Diagnostic{}, value)
}

func (c *Compilation) EmitTree(writer io.StringWriter) {
    statement := c.GetStatement()
    Binding.WriteTo(writer, statement)
}

func (c *Compilation) GetStatement() Binding.BoundStatement {
    result := c.globalScope.Statement
    return Lowering.LowererLower(result)
}
