package BoundBinaryOperator

import (
    "minsk/CodeAnalysis/Binding/Kind/BoundBinaryOperatorKind"
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"

    "reflect"
)

type BoundBinaryOperator struct {
        SyntaxKind SyntaxKind.SyntaxKind
        Kind BoundBinaryOperatorKind.BoundBinaryOperatorKind
        LeftType reflect.Kind
        RightType reflect.Kind
        ResultType reflect.Kind
}

func NewBoundBinaryOperator(
        syntaxKind SyntaxKind.SyntaxKind,
        kind BoundBinaryOperatorKind.BoundBinaryOperatorKind,
        leftType reflect.Kind,
        rightType reflect.Kind,
        resultType reflect.Kind,
    ) *BoundBinaryOperator {
        return &BoundBinaryOperator{
            SyntaxKind: syntaxKind,
            Kind: kind,
            LeftType: leftType,
            RightType: rightType,
            ResultType: resultType,
        }
}

var _operators = []*BoundBinaryOperator{
    NewBoundBinaryOperator(SyntaxKind.PlusToken, BoundBinaryOperatorKind.Addition, reflect.Int, reflect.Int, reflect.Int),
    NewBoundBinaryOperator(SyntaxKind.MinusToken, BoundBinaryOperatorKind.Subtraction, reflect.Int, reflect.Int, reflect.Int),
    NewBoundBinaryOperator(SyntaxKind.StarToken, BoundBinaryOperatorKind.Multiplication, reflect.Int, reflect.Int, reflect.Int),
    NewBoundBinaryOperator(SyntaxKind.SlashToken, BoundBinaryOperatorKind.Division, reflect.Int, reflect.Int, reflect.Int),
    NewBoundBinaryOperator(SyntaxKind.EqualsEqualsToken, BoundBinaryOperatorKind.Equals, reflect.Int, reflect.Int, reflect.Bool),
    NewBoundBinaryOperator(SyntaxKind.BangEqualsToken, BoundBinaryOperatorKind.NotEquals, reflect.Int, reflect.Int, reflect.Bool),
    NewBoundBinaryOperator(SyntaxKind.LessToken, BoundBinaryOperatorKind.Less, reflect.Int, reflect.Int, reflect.Bool),
    NewBoundBinaryOperator(SyntaxKind.LessOrEqualsToken, BoundBinaryOperatorKind.LessOrEquals, reflect.Int, reflect.Int, reflect.Bool),
    NewBoundBinaryOperator(SyntaxKind.GreaterToken, BoundBinaryOperatorKind.Greater, reflect.Int, reflect.Int, reflect.Bool),
    NewBoundBinaryOperator(SyntaxKind.GreaterOrEqualsToken, BoundBinaryOperatorKind.GreaterOrEquals, reflect.Int, reflect.Int, reflect.Bool),
    NewBoundBinaryOperator(SyntaxKind.AmpersandToken, BoundBinaryOperatorKind.BitwiseAnd, reflect.Int, reflect.Int, reflect.Int),
    NewBoundBinaryOperator(SyntaxKind.PipeToken, BoundBinaryOperatorKind.BitwiseOr, reflect.Int, reflect.Int, reflect.Int),
    NewBoundBinaryOperator(SyntaxKind.HatToken, BoundBinaryOperatorKind.BitwiseXor, reflect.Int, reflect.Int, reflect.Int),

    NewBoundBinaryOperator(SyntaxKind.AmpersandToken, BoundBinaryOperatorKind.BitwiseAnd, reflect.Bool, reflect.Bool, reflect.Bool),
    NewBoundBinaryOperator(SyntaxKind.AmpersandAmpersandToken, BoundBinaryOperatorKind.LogicalAnd, reflect.Bool, reflect.Bool, reflect.Bool),
    NewBoundBinaryOperator(SyntaxKind.PipeToken, BoundBinaryOperatorKind.BitwiseOr, reflect.Bool, reflect.Bool, reflect.Bool),
    NewBoundBinaryOperator(SyntaxKind.PipePipeToken, BoundBinaryOperatorKind.LogicalOr, reflect.Bool, reflect.Bool, reflect.Bool),
    NewBoundBinaryOperator(SyntaxKind.HatToken, BoundBinaryOperatorKind.BitwiseXor, reflect.Bool, reflect.Bool, reflect.Bool),
    NewBoundBinaryOperator(SyntaxKind.EqualsEqualsToken, BoundBinaryOperatorKind.Equals, reflect.Bool, reflect.Bool, reflect.Bool),
    NewBoundBinaryOperator(SyntaxKind.BangEqualsToken, BoundBinaryOperatorKind.NotEquals, reflect.Bool, reflect.Bool, reflect.Bool),
}

func Bind(syntaxKind SyntaxKind.SyntaxKind, leftType reflect.Kind, rightType reflect.Kind) *BoundBinaryOperator {
    for _, op := range _operators {
        if op.SyntaxKind == syntaxKind && leftType == op.LeftType && rightType == op.RightType {
            return op
        }
    }

    return nil
}
