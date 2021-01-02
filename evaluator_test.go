package minsk

import (
    "testing"
)

func TestEvaluate(t *testing.T) {
    numberToken1 := NewSyntaxToken(NumberToken, 0, []rune{'5'}, 5)
    n := NewNumberExpressionSyntax(numberToken1)
    e := NewEvaluator(n)
    result := e.Evaluate()
    if result != 5 {
        t.Errorf("Result should be 5, got %+v", result)
    }

    left := NewSyntaxToken(NumberToken, 0, []rune{'5'}, 5)
    right := NewSyntaxToken(NumberToken, 0, []rune{'6'}, 6)
    binaryTests := []struct{
            Operator *SyntaxToken
            Expected int
        } {
            {
                Operator: NewSyntaxToken(PlusToken, 0, []rune{'+'}, nil),
                Expected: 11,
            },
            {
                Operator: NewSyntaxToken(MinusToken, 0, []rune{'-'}, nil),
                Expected: -1,
            },
            {
                Operator: NewSyntaxToken(StarToken, 0, []rune{'-'}, nil),
                Expected: 30,
            },
            {
                Operator: NewSyntaxToken(SlashToken, 0, []rune{'-'}, nil),
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
