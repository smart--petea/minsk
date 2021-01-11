package CodeAnalysis

import (
    "testing"

    "minsk/CodeAnalysis/SyntaxKind"
)

func TestEvaluateUnary(t *testing.T) {
    unaryTests := []struct{
        OperatorToken *SyntaxToken
        OperandToken *SyntaxToken
        Expected int
    }{
        {
            OperatorToken: NewSyntaxToken(SyntaxKind.PlusToken, 0, []rune{'+'}, nil),
            OperandToken: NewSyntaxToken(SyntaxKind.NumberToken, 0, []rune{'5'}, 5),
            Expected: 5,
        },
        {
            OperatorToken: NewSyntaxToken(SyntaxKind.MinusToken, 0, []rune{'+'}, nil),
            OperandToken: NewSyntaxToken(SyntaxKind.NumberToken, 0, []rune{'5'}, 5),
            Expected: -5,
        },
    }

    for _, unaryTest := range unaryTests {
        unaryToken := NewUnaryExpressionSyntax(unaryTest.OperatorToken, unaryTest.OperandToken)
        e := NewEvaluator(unaryToken)
        output := e.Evaluate()

        if output != unaryTest.Expected {
            t.Errorf("%s %s->%d expected %d",
                string(unaryTest.OperatorToken.Runes),
                string(unaryTest.OperandToken.Runes),
                output,
                unaryTest.Expected,
            )
        }
    }
}

func TestEvaluateBinary(t *testing.T) {
    numberToken1 := NewSyntaxToken(SyntaxKind.NumberToken, 0, []rune{'5'}, 5)
    n := NewLiteralExpressionSyntax(numberToken1)
    e := NewEvaluator(n)
    result := e.Evaluate()
    if result != 5 {
        t.Errorf("Result should be 5, got %+v", result)
    }

    left := NewSyntaxToken(SyntaxKind.NumberToken, 0, []rune{'5'}, 5)
    right := NewSyntaxToken(SyntaxKind.NumberToken, 0, []rune{'6'}, 6)
    binaryTests := []struct{
            Operator *SyntaxToken
            Expected int
        } {
            {
                Operator: NewSyntaxToken(SyntaxKind.PlusToken, 0, []rune{'+'}, nil),
                Expected: 11,
            },
            {
                Operator: NewSyntaxToken(SyntaxKind.MinusToken, 0, []rune{'-'}, nil),
                Expected: -1,
            },
            {
                Operator: NewSyntaxToken(SyntaxKind.StarToken, 0, []rune{'-'}, nil),
                Expected: 30,
            },
            {
                Operator: NewSyntaxToken(SyntaxKind.SlashToken, 0, []rune{'-'}, nil),
                Expected: 5/6,
            },
        }
    for _, binaryTest := range binaryTests {
        b := NewBinaryExpressionSyntax(left, binaryTest.Operator, right)
        e := NewEvaluator(b)
        result := e.Evaluate()
        if result != binaryTest.Expected {
            t.Errorf("(%d %s %d) should be %d, got %+v", left.Value(), string(binaryTest.Operator.Runes), right.Value(), binaryTest.Expected, result)
        }
    }
}
