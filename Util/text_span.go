package Util

type TextSpan struct {
    Start int
    Length int
}

func NewTextSpan(start, length int) *TextSpan {
    return &TextSpan{
        Start: start,
        Length: length,
    }
}

func (ts *TextSpan) End() int {
    return ts.Start + ts.Length
}
