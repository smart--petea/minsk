package Text

type TextSpanComparer struct {}

func NewTextSpanComparer() *TextSpanComparer {
    return &TextSpanComparer{}
}

func (t *TextSpanComparer) Compare(x *TextSpan, y *TextSpan) int {
    cmp := x.Start - y.Start
    if cmp == 0 {
        cmp = x.Length - y.Length
    }

    return cmp
}
