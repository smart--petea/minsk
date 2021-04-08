package Util

import (
    "math/rand"
)

type Random struct {}

func NewRandom() *Random {
    return &Random{}
}

func (r *Random) Next(max int) int {
    return rand.Intn(max)
}
