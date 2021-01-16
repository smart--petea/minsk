package Binding

import (
    "minsk/CodeAnalysis/Binding/TypeCarrier"
)

type BoundExpression interface {
    BoundNode

    GetTypeCarrier() TypeCarrier.TypeCarrier
}
