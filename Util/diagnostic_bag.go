package Util

import (
    "fmt"
)

/* todo
type Diagnostic struct {
    Diagnostics []string
}

func (d *Diagnostic) AddDiagnostic(format string, args ...interface{}) {
    d.Diagnostics = append(d.Diagnostics, fmt.Sprintf(format, args...))
}
*/

type DiagnosticBag struct {
    diagnostics []Diagnostic
}

func (db *DiagnosticBag) report(span *TextSpan, message string) {
    diagnostic := NewDiagnostic(span, message)

    db.diagnostics = append(db.diagnostics, diagnostic)
}

func (db *DiagnosticBag) ReportInvalidNumber(span *TextSpan, runes []rune, kind reflect.Kind) {
    message := fmt.Sprintf("The number %s isn't valid %s.", string(runes), kind)

    db.report(span, message)
}

func (db *DiagnosticBag) ReportBadCharacter(position int, character rune) {
    message := fmt.Sprintf("Bad character input: '%s'.", string(character))
    span := NewTextSpan(position, 1)

    db.report(span, message)
}

func (db *DiagnosticBag) GetDiagnostics() []Diagnostic {
    return db.diagnostics
}

func (db *DiagnosticBag) AddDiagnosticsRange(diagnostics []Diagnostic) {
    db.diagnostics = append(d.diagnostics, diagnostics...)
}

func (db *DiagnosticBag) ReportUnexpectedToken(span *TextSpan, actualKind SyntaxKind, expectedKind SyntaxKind) {
    message := fmt.Sprintf("Unexpected token <%s>, expected <%s>.", actualKind, expectedKind)

    db.report(span, message)
}

func (db *DiagnosticBag) ReportUndefinedUnaryOperator(span *TextSpan, operatorText []rune, operandType reflect.Kind) {
    message := fmt.Sprintf("Unary operator '%+v' is not defined for type %T", string(operatorText), operandType)

    db.report(span, message)
}

func (db *DiagnosticBag) ReportUndefinedBinaryOperator(span *TextSpan, operatorRunes []rune, leftType reflect.Kind, rightType reflect.Kind) {
    message := fmt.Sprintf("Binary operator '%+v' is not defined for types %T and %T", string(operatorRunes), leftType, rightType)

    db.report(span, message)
}

/*
func (db *diagnosticBag) GetEnumerator() []Diagnostic {
    return db.diagnostics
}
*/
