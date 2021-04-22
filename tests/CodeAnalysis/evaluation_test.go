package CodeAnalysisTest

import (
    "testing"
    "minsk/CodeAnalysis/Syntax"
    CA "minsk/CodeAnalysis"
    "minsk/CodeAnalysis/Syntax/Symbols"
    "fmt"
)

func TestEvaluations(t *testing.T) {
    tests := []struct{
        text string
        expectedValue interface{}
    } {
        {
            text: "1",
            expectedValue: 1,
        },
        {
            text: "+1",
            expectedValue: 1,
        },
        {
            text: "-1",
            expectedValue: -1,
        },
        {
            text: "~1",
            expectedValue: -2,
        },
        {
            text: "14 + 12",
            expectedValue: 26,
        },
        {
            text: "12 - 3",
            expectedValue: 9,
        },
        {
            text: "4 * 2",
            expectedValue: 8,
        },
        {
            text: "9 / 3",
            expectedValue: 3,
        },
        {
            text: "(10)",
            expectedValue: 10,
        },
        {
            text: "12 == 3",
            expectedValue: false,
        },
        {
            text: "3 == 3",
            expectedValue: true,
        },
        {
            text: "12 != 3",
            expectedValue: true,
        },
        {
            text: "3 != 3",
            expectedValue: false,
        },
        {
            text: "3 < 4",
            expectedValue: true,
        },
        {
            text: "5 < 4",
            expectedValue: false,
        },
        {
            text: "4 <= 4",
            expectedValue: true,
        },
        {
            text: "4 <= 5",
            expectedValue: true,
        },
        {
            text: "5 <= 4",
            expectedValue: false,
        },
        {
            text: "4 > 3",
            expectedValue: true,
        },
        {
            text: "4 > 5",
            expectedValue: false,
        },
        {
            text: "4 >= 4",
            expectedValue: true,
        },
        {
            text: "5 >= 4",
            expectedValue: true,
        },
        {
            text: "4 >= 5",
            expectedValue: false,
        },

        {
            text: "1 | 2",
            expectedValue: 3,
        },
        {
            text: "1 | 0",
            expectedValue: 1,
        },
        {
            text: "1 & 3",
            expectedValue: 1,
        },
        {
            text: "1 & 0",
            expectedValue: 0,
        },
        {
            text: "1 ^ 0",
            expectedValue: 1,
        },
        {
            text: "0 ^ 1",
            expectedValue: 1,
        },
        {
            text: "1 ^ 3",
            expectedValue: 2,
        },
        {
            text: "false == false",
            expectedValue: true,
        },
        {
            text: "true == false",
            expectedValue: false,
        },
        {
            text: "false != false",
            expectedValue: false,
        },
        {
            text: "true != false",
            expectedValue: true,
        },
        {
            text: "false | false",
            expectedValue: false,
        },
        {
            text: "false | true",
            expectedValue: true,
        },
        {
            text: "true | false",
            expectedValue: true,
        },
        {
            text: "true | true",
            expectedValue: true,
        },

        {
            text: "false & false",
            expectedValue: false,
        },
        {
            text: "false & true",
            expectedValue: false,
        },
        {
            text: "true & false",
            expectedValue: false,
        },
        {
            text: "true & true",
            expectedValue: true,
        },
        {
            text: "false ^ false",
            expectedValue: false,
        },
        {
            text: "false ^ true",
            expectedValue: true,
        },
        {
            text: "true ^ false",
            expectedValue: true,
        },
        {
            text: "true ^ true",
            expectedValue: false,
        },
        {
            text: "~true",
            expectedValue: false,
        },
        {
            text: "~false",
            expectedValue: true,
        },
        {
            text: "true",
            expectedValue: true,
        },
        {
            text: "false",
            expectedValue: false,
        },
        {
            text: "!true",
            expectedValue: false,
        },
        {
            text: "!false",
            expectedValue: true,
        },
        {
            text: "\"test\"",
            expectedValue: "test",
        },
        {
            text: "\"te\"\"st\"",
            expectedValue: "te\"st",
        },
        {
            text: "\"test\"==\"test\"",
            expectedValue: true,
        },
        {
            text: "\"test\"!=\"test\"",
            expectedValue: false,
        },
        {
            text: "\"test\"==\"abc\"",
            expectedValue: false,
        },
        {
            text: "\"test\"!=\"abc\"",
            expectedValue: true,
        },
        {
            text: "{ var a = 0 (a = 10) * a }",
            expectedValue: 100,
        },
        {
            text: "{ var a = 0 if a == 0 a = 10 a }",
            expectedValue: 10,
        },
        {
            text: "{ var a = 0 if a == 4 a = 10 a }",
            expectedValue: 0,
        },

        {
            text: "{ var a = 0 if a == 0 a = 10 else a = 5 a }",
            expectedValue: 10,
        },
        {
            text: "{ var a = 0 if a == 4 a = 10 else a = 5 a }",
            expectedValue: 5,
        },
        {
            text: "{ var i = 10 var result = 0 while i > 0 { result = result + i i = i - 1 } result }",
            expectedValue: 55,
        },
        {
            text: "{ var result = 0 for i = 1 to 10 { result = result + i} result}",
            expectedValue: 55,
        },
    }

    for _, test := range tests {
        syntaxTree := Syntax.SyntaxTreeParse(test.text)
        compilation := CA.NewCompilation(syntaxTree)
        variables := make(map[*Symbols.VariableSymbol]interface{})
        result := compilation.Evaluate(variables)

        if len(result.Diagnostics) > 0 {
            t.Errorf("Diagnostics not empty %+v", result.Diagnostics)
        }

        if test.expectedValue != result.Value {
            t.Errorf("(%s)=%+v, expected=%+v", test.text, result.Value, test.expectedValue)
        }
    }
}

func assertDiagnostics(text, diagnosticText string, t *testing.T) {
    annotatedText := AnnotatedTextParse(text)
    syntaxTree := Syntax.SyntaxTreeParse(annotatedText.Text)
    fmt.Printf("%s %d\n", annotatedText.Text, len(annotatedText.Text))
    compilation := CA.NewCompilation(syntaxTree)
    variables := make(map[*Symbols.VariableSymbol]interface{})
    result := compilation.Evaluate(variables)

    expectedDiagnostics := AnnotatedTextUnindentLines(diagnosticText)
    if len(annotatedText.Spans) != len(expectedDiagnostics) {
        message := fmt.Sprintf("Must mark as many spans as there are expected diagnostics")
        panic(message)
    }

    if len(expectedDiagnostics) != len(result.Diagnostics) {
        t.Errorf("len(expectedDiagnostics) != len(result.Diagnostics). actual=%d expected=%d", len(expectedDiagnostics), len(result.Diagnostics))
    }

    for i, _ := range expectedDiagnostics {
        expectedMessage := expectedDiagnostics[i]
        actualMessage := result.Diagnostics[i].Message
        if expectedMessage != actualMessage {
            t.Errorf("actualMessage != expectedMessage. actual=%s expected=%s", expectedMessage, actualMessage)
        }

        expectedSpan := annotatedText.Spans[i]
        actualSpan := result.Diagnostics[i].Span
        if *expectedSpan != *actualSpan {
            t.Errorf("actualSpan != expectedSpan. actual=%+v expected=%+v", expectedSpan, actualSpan)
        }
    }
}

func TestEvaluatorVariableDeclarationsReportsRedeclaration(t *testing.T) {
    text := `
        {
            var x = 10
            var y = 100
            {
                var x = 10
            }
            var [x] = 5
        }
    `

    diagnostics := `
        Variable x is already declared
    `

    assertDiagnostics(text, diagnostics, t)
}

func TestEvaluatorBlockStatementNoInfiniteLoop(t *testing.T) {
    text := `
    {
        [)]
    }
    `

    diagnostics := `
        Unexpected token <CloseParenthisToken>, expected <IdentifierToken>.
    `

    assertDiagnostics(text, diagnostics, t)
}

func TestEvaluatorIfStatementsReportsCannotConvert(t *testing.T) {
    text := `
    {
        var x = 0
        if [10]
            x = 10
    }
    `

    diagnostics := `
        Cannot convert type 'int' to 'bool'.
    `

    assertDiagnostics(text, diagnostics, t)
}

func TestEvaluatorWhileStatementsReportsCannotConvert(t *testing.T) {
    text := `
    {
        var x = 0
        while [10]
            x = 10
    }
    `

    diagnostics := `
        Cannot convert type 'int' to 'bool'.
    `

    assertDiagnostics(text, diagnostics, t)
}

func TestEvaluatorForStatementsReportsCannotConvertLowerBound(t *testing.T) {
    text := `
    {
        var result = 0
        for i = [false] to 10
            result = result + i
    }
    `

    diagnostics := `
        Cannot convert type 'bool' to 'int'.
    `

    assertDiagnostics(text, diagnostics, t)
}

func TestEvaluatorForStatementsReportsCannotConvertUpperBound(t *testing.T) {
    text := `
    {
        var result = 0
        for i = 1 to [true]
            result = result + i
    }
    `

    diagnostics := `
        Cannot convert type 'bool' to 'int'.
    `

    assertDiagnostics(text, diagnostics, t)
}

func TestEvaluatorUnaryExpressionReportsUndefined(t *testing.T) {
    text := `[+]true`

    diagnostics := `
        Unary operator '+' is not defined for type 'bool'.
    `

    assertDiagnostics(text, diagnostics, t)
}

func TestEvaluatorNameExpressionReportsUndefined(t *testing.T) {
    text := `[x] * 10 `

    diagnostics := `
        Variable x doesn't exist
    `

    assertDiagnostics(text, diagnostics, t)
}

func TestEvaluatorNameExpressionReportsNoErrorForInsertedToken(t *testing.T) {
    text := `[]`

    diagnostics := `
        Unexpected token <EndOfFileToken>, expected <IdentifierToken>.
    `

    assertDiagnostics(text, diagnostics, t)
}

func TestEvaluatorAssignmentExpressionReportsUndefined(t *testing.T) {
    text := `[x] = 10 `

    diagnostics := `
        Variable x doesn't exist
    `

    assertDiagnostics(text, diagnostics, t)
}

func TestEvaluatorAssignmentExpressionReportsCannotAssign(t *testing.T) {
    text := `
    {
        let x = 10
        x [=] 0
    }
    `

    diagnostics := `
        Variable x is read-only and cannot be assigned to.
    `

    assertDiagnostics(text, diagnostics, t)
}

func TestEvaluatorAssignmentExpressionReportsCannotConvert(t *testing.T) {
    text := `
    {
        var x = 10
        [x] = true
    }
    `

    diagnostics := `
        Cannot convert type 'bool' to 'int'.
    `

    assertDiagnostics(text, diagnostics, t)
}

func TestEvaluatorBinaryExpressionReportsUndefined(t *testing.T) {
    text := `10 [+] false`

    diagnostics := `
        Binary operator '+' is not defined for types 'int' and 'bool'.
    `

    assertDiagnostics(text, diagnostics, t)
}
