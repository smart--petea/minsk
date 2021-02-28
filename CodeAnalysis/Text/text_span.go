package Text

type TextSpan struct {
    Start int
    Length int
}

func NewTextSpan(start, length int) *TextSpan {
    if length < 0 {
        length = 0
    }

    return &TextSpan{
        Start: start,
        Length: length,
    }
}

func (ts *TextSpan) End() int {
    return ts.Start + ts.Length
}
