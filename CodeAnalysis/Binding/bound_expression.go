package Binding

import (
    "reflect"
)

type BoundExpression interface {
    BoundNode

    GetType() reflect.Kind
}
