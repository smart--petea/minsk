package CodeAnalysisTest

import (
    "testing"
    "minsk/CodeAnalysis/Syntax"
    CA "minsk/CodeAnalysis"
    "minsk/Util"
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
            text: "{ var a = 0 (a = 10) * a }",
            expectedValue: 100,
        },
    }

    for _, test := range tests {
        syntaxTree := Syntax.ParseSyntaxTree(test.text)
        compilation := CA.NewCompilation(syntaxTree)
        variables := make(map[*Util.VariableSymbol]interface{})
        result := compilation.Evaluate(variables)

        if len(result.Diagnostics) > 0 {
            t.Errorf("Diagnostics not empty %+v", result.Diagnostics)
        }

        if test.expectedValue != result.Value {
            t.Errorf("(%s)=%+v, expected=%+v", test.text, result.Value, test.expectedValue)
        }
    }
}
