package BoundUnaryOperator

import (
    "minsk/CodeAnalysis/Binding/Kind/BoundUnaryOperatorKind"
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"

    "reflect"
)

type BoundUnaryOperator struct {
        SyntaxKind SyntaxKind.SyntaxKind
        Kind BoundUnaryOperatorKind.BoundUnaryOperatorKind
        OperandType reflect.Kind
        ResultType reflect.Kind
}

func NewBoundUnaryOperator(
        syntaxKind SyntaxKind.SyntaxKind,
        kind BoundUnaryOperatorKind.BoundUnaryOperatorKind,
        operandType reflect.Kind,
        resultType reflect.Kind,
    ) *BoundUnaryOperator {
        return &BoundUnaryOperator{
            SyntaxKind: syntaxKind,
            Kind: kind,
            OperandType: operandType,
            ResultType: resultType,
        }
}

var _operators = []*BoundUnaryOperator{
    NewBoundUnaryOperator(SyntaxKind.BangToken, BoundUnaryOperatorKind.LogicalNegation, reflect.Bool, reflect.Bool),
    NewBoundUnaryOperator(SyntaxKind.PlusToken, BoundUnaryOperatorKind.Identity, reflect.Int, reflect.Int),
    NewBoundUnaryOperator(SyntaxKind.MinusToken, BoundUnaryOperatorKind.Negation, reflect.Int, reflect.Int),
}

func Bind(syntaxKind SyntaxKind.SyntaxKind, operandType reflect.Kind) *BoundUnaryOperator {
    for _, op := range _operators {
        if op.SyntaxKind == syntaxKind && op.OperandType == operandType {
            return op
        }
    }

    return nil
}
