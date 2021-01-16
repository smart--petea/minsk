package Util

import (
    "fmt"
)

type Diagnostic struct {
    Diagnostics []string
}

func (d *Diagnostic) AddDiagnostic(format string, args ...interface{}) {
    d.Diagnostics = append(d.Diagnostics, fmt.Sprintf(format, args...))
}

func (d *Diagnostic) LoadDiagnostics(diagnostics []string) {
    d.Diagnostics = append(d.Diagnostics, diagnostics...)
}

func (d *Diagnostic) GetDiagnostics() []string {
    return d.Diagnostics
}
