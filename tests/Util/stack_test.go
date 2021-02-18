package UtilTest

import (
    "testing"
    "minsk/Util"
)

func TestStack(t *testing.T) {
    var a1, a2, a3 interface{}
    stack := Util.NewStack()

    stack.Push(a1)
    stack.Push(a2)
    stack.Push(a3)
    if stack.Count() != 3 {
        t.Errorf("a1,a2,a3 in stack. count should be 3")
    }

    b3 := stack.Pop()
    if a3 != b3 {
        t.Errorf("a3 != b3")
    }
    if stack.Count() != 2 {
        t.Errorf("a1,a2 in stack. count should be 2")
    }

    b2 := stack.Pop()
    if a2 != b2 {
        t.Errorf("a2 != b2")
    }
    if stack.Count() != 1 {
        t.Errorf("a1 in stack. count should be 1")
    }

    b1 := stack.Pop()
    if a1 != b1 {
        t.Errorf("a1 != b1")
    }
    if stack.Count() != 0 {
        t.Errorf("nothing in stack. count should be 0")
    }
}
