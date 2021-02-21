package Util

import (
    "fmt"
    "reflect"
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/CodeAnalysis/Text"
)

type DiagnosticBag struct {
    diagnostics []*Diagnostic
}

func (db *DiagnosticBag) report(span *Text.TextSpan, message string) {
    diagnostic := NewDiagnostic(span, message)

    db.diagnostics = append(db.diagnostics, diagnostic)
}

func (db *DiagnosticBag) ReportInvalidNumber(span *Text.TextSpan, runes []rune, kind reflect.Kind) {
    message := fmt.Sprintf("The number %s isn't valid %s.", string(runes), kind)

    db.report(span, message)
}

func (db *DiagnosticBag) ReportBadCharacter(position int, character rune) {
    message := fmt.Sprintf("Bad character input: '%s'.", string(character))
    span := Text.NewTextSpan(position, 1)

    db.report(span, message)
}

func (db *DiagnosticBag) GetDiagnostics() []*Diagnostic {
    cp := make([]*Diagnostic, len(db.diagnostics))
    copy(cp, db.diagnostics)
    return cp 
}

func (db *DiagnosticBag) AddDiagnosticsRange(diagnostics []*Diagnostic) {
    db.diagnostics = append(db.diagnostics, diagnostics...)
}

func (db *DiagnosticBag) ReportUnexpectedToken(span *Text.TextSpan, actualKind SyntaxKind.SyntaxKind, expectedKind SyntaxKind.SyntaxKind) {
    message := fmt.Sprintf("Unexpected token <%s>, expected <%s>.", actualKind, expectedKind)

    db.report(span, message)
}

func (db *DiagnosticBag) ReportUndefinedUnaryOperator(span *Text.TextSpan, operatorText []rune, operandType reflect.Kind) {
    message := fmt.Sprintf("Unary operator '%+v' is not defined for type %s", string(operatorText), operandType)

    db.report(span, message)
}

func (db *DiagnosticBag) ReportUndefinedBinaryOperator(span *Text.TextSpan, operatorRunes []rune, leftType reflect.Kind, rightType reflect.Kind) {
    message := fmt.Sprintf("Binary operator '%+v' is not defined for types %s and %s", string(operatorRunes), leftType, rightType)

    db.report(span, message)
}

func (db *DiagnosticBag) ReportUndefinedName(span *Text.TextSpan, name string) {
    message := fmt.Sprintf("Variable %s doesn't exist", name)

    db.report(span, message)
}

func (db *DiagnosticBag) ReportVariableAlreadyDeclared(span *Text.TextSpan, name string) {
    message := fmt.Sprintf("Variable %s is already declared", name)

    db.report(span, message)
}

func (db *DiagnosticBag) ReportCannotConvert(span *Text.TextSpan, fromType, toType reflect.Kind) {
    message := fmt.Sprintf("Cannot convert type '%s' to '%s'.", fromType, toType)

    db.report(span, message)
}

func (db *DiagnosticBag) ReportCannotAssign(span *Text.TextSpan, name string) {
    message := fmt.Sprintf("Variable %s is read-only and cannot be assigned to.", name)

    db.report(span, message)
}
