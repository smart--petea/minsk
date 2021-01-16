package Binding

import (
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
)

type BoundNode interface {
    Kind() BoundNodeKind.BoundNodeKind
}
