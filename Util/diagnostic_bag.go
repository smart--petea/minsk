package Util

import (
    "fmt"
    "reflect"
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
)

type DiagnosticBag struct {
    diagnostics []*Diagnostic
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

func (db *DiagnosticBag) GetDiagnostics() []*Diagnostic {
    return db.diagnostics
}

func (db *DiagnosticBag) AddDiagnosticsRange(diagnostics []*Diagnostic) {
    db.diagnostics = append(db.diagnostics, diagnostics...)
}

func (db *DiagnosticBag) ReportUnexpectedToken(span *TextSpan, actualKind SyntaxKind.SyntaxKind, expectedKind SyntaxKind.SyntaxKind) {
    message := fmt.Sprintf("Unexpected token <%s>, expected <%s>.", actualKind, expectedKind)

    db.report(span, message)
}

func (db *DiagnosticBag) ReportUndefinedUnaryOperator(span *TextSpan, operatorText []rune, operandType reflect.Kind) {
    message := fmt.Sprintf("Unary operator '%+v' is not defined for type %s", string(operatorText), operandType)

    db.report(span, message)
}

func (db *DiagnosticBag) ReportUndefinedBinaryOperator(span *TextSpan, operatorRunes []rune, leftType reflect.Kind, rightType reflect.Kind) {
    message := fmt.Sprintf("Binary operator '%+v' is not defined for types %s and %s", string(operatorRunes), leftType, rightType)

    db.report(span, message)
}

func (db *DiagnosticBag) ReportUndefinedName(span *TextSpan, name string) {
    message := fmt.Sprintf("Variable %s doesn't exist", name)

    db.report(span, message)
}
