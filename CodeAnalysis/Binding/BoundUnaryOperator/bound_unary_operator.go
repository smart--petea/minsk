package BoundUnaryOperator

import (
    "minsk/CodeAnalysis/Binding/Kind/BoundUnaryOperatorKind"
    "minsk/CodeAnalysis/Binding/TypeCarrier"

    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
)

type BoundUnaryOperator struct {
        SyntaxKind SyntaxKind.SyntaxKind
        Kind BoundUnaryOperatorKind.BoundUnaryOperatorKind
        OperandTypeCarrier TypeCarrier.TypeCarrier
        ResultTypeCarrier TypeCarrier.TypeCarrier
}

func NewBoundUnaryOperator(
        syntaxKind SyntaxKind.SyntaxKind,
        kind BoundUnaryOperatorKind.BoundUnaryOperatorKind,
        operandTypeCarrier TypeCarrier.TypeCarrier,
        resultTypeCarrier TypeCarrier.TypeCarrier,
    ) *BoundUnaryOperator {
        return &BoundUnaryOperator{
            SyntaxKind: syntaxKind,
            Kind: kind,
            OperandTypeCarrier: operandTypeCarrier,
            ResultTypeCarrier: resultTypeCarrier,
        }
}

var _operators = []*BoundUnaryOperator{
    NewBoundUnaryOperator(SyntaxKind.BangToken, BoundUnaryOperatorKind.LogicalNegation, TypeCarrier.Bool(), TypeCarrier.Bool()),
    NewBoundUnaryOperator(SyntaxKind.PlusToken, BoundUnaryOperatorKind.Identity, TypeCarrier.Int(), TypeCarrier.Int()),
    NewBoundUnaryOperator(SyntaxKind.MinusToken, BoundUnaryOperatorKind.Negation, TypeCarrier.Int(), TypeCarrier.Int()),
}

func Bind(syntaxKind SyntaxKind.SyntaxKind, operandTypeCarrier TypeCarrier.TypeCarrier) *BoundUnaryOperator {
    for _, op := range _operators {
        if op.SyntaxKind == syntaxKind && TypeCarrier.Same(op.OperandTypeCarrier, operandTypeCarrier) {
            return op
        }
    }

    return nil
}
