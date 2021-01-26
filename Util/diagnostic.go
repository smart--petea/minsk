package Util

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
