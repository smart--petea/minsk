package BoundBinaryOperator

import (
    "minsk/CodeAnalysis/Binding/Kind/BoundBinaryOperatorKind"
    "minsk/CodeAnalysis/Binding/TypeCarrier"

    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
)

type BoundBinaryOperator struct {
        SyntaxKind SyntaxKind.SyntaxKind
        Kind BoundBinaryOperatorKind.BoundBinaryOperatorKind
        LeftTypeCarrier TypeCarrier.TypeCarrier
        RightTypeCarrier TypeCarrier.TypeCarrier
        ResultTypeCarrier TypeCarrier.TypeCarrier
}

func NewBoundBinaryOperator(
        syntaxKind SyntaxKind.SyntaxKind,
        kind BoundBinaryOperatorKind.BoundBinaryOperatorKind,
        leftTypeCarrier TypeCarrier.TypeCarrier,
        rightTypeCarrier TypeCarrier.TypeCarrier,
        resultTypeCarrier TypeCarrier.TypeCarrier,
    ) *BoundBinaryOperator {
        return &BoundBinaryOperator{
            SyntaxKind: syntaxKind,
            Kind: kind,
            LeftTypeCarrier: leftTypeCarrier,
            RightTypeCarrier: rightTypeCarrier,
            ResultTypeCarrier: resultTypeCarrier,
        }
}

var _operators = []*BoundBinaryOperator{
    NewBoundBinaryOperator(SyntaxKind.PlusToken, BoundBinaryOperatorKind.Addition, TypeCarrier.Int(), TypeCarrier.Int(), TypeCarrier.Int()),
    NewBoundBinaryOperator(SyntaxKind.MinusToken, BoundBinaryOperatorKind.Subtraction, TypeCarrier.Int(), TypeCarrier.Int(), TypeCarrier.Int()),
    NewBoundBinaryOperator(SyntaxKind.StarToken, BoundBinaryOperatorKind.Multiplication, TypeCarrier.Int(), TypeCarrier.Int(), TypeCarrier.Int()),
    NewBoundBinaryOperator(SyntaxKind.SlashToken, BoundBinaryOperatorKind.Division, TypeCarrier.Int(), TypeCarrier.Int(), TypeCarrier.Int()),

    NewBoundBinaryOperator(SyntaxKind.EqualsEqualsToken, BoundBinaryOperatorKind.Equals, TypeCarrier.Int(), TypeCarrier.Int(), TypeCarrier.Bool()),
    NewBoundBinaryOperator(SyntaxKind.BangEqualsToken, BoundBinaryOperatorKind.NotEquals, TypeCarrier.Int(), TypeCarrier.Int(), TypeCarrier.Bool()),

    NewBoundBinaryOperator(SyntaxKind.AmpersandAmpersandToken, BoundBinaryOperatorKind.LogicalAnd, TypeCarrier.Bool(), TypeCarrier.Bool(), TypeCarrier.Bool()),
    NewBoundBinaryOperator(SyntaxKind.PipePipeToken, BoundBinaryOperatorKind.LogicalOr, TypeCarrier.Bool(), TypeCarrier.Bool(), TypeCarrier.Bool()),
    NewBoundBinaryOperator(SyntaxKind.EqualsEqualsToken, BoundBinaryOperatorKind.Equals, TypeCarrier.Bool(), TypeCarrier.Bool(), TypeCarrier.Bool()),
    NewBoundBinaryOperator(SyntaxKind.BangEqualsToken, BoundBinaryOperatorKind.NotEquals, TypeCarrier.Bool(), TypeCarrier.Bool(), TypeCarrier.Bool()),
}

func Bind(syntaxKind SyntaxKind.SyntaxKind, leftTypeCarrier TypeCarrier.TypeCarrier, rightTypeCarrier TypeCarrier.TypeCarrier) *BoundBinaryOperator {
    for _, op := range _operators {
        if op.SyntaxKind == syntaxKind &&
            TypeCarrier.Same(op.LeftTypeCarrier, leftTypeCarrier) &&
            TypeCarrier.Same(op.RightTypeCarrier, rightTypeCarrier) {
            return op
        }
    }

    return nil
}
