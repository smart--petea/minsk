package Util

import (
    "fmt"
)

/*
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
*/

type Diagnostic struct {
    Span *TextSpan
    Message string
}

func NewDiagnostic(span *TextSpan, message string) *Diagnostic {
    return &Diagnostic{
        Span: span, 
        Message: message,
    }
}

func (d *Diagnostic) String() string {
    return d.Message
}
