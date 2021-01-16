package Binding

type BoundExpression interface {
    BoundNode

    GetTypeCarrier() TypeCarrier
}
