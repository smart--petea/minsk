package Util

import (
    "minsk/CodeAnalysis/Text"
)

type Diagnostic struct {
    Span *Text.TextSpan
    Message string
}

func NewDiagnostic(span *Text.TextSpan, message string) *Diagnostic {
    return &Diagnostic{
        Span: span, 
        Message: message,
    }
}

func (d *Diagnostic) String() string {
    return d.Message
}
