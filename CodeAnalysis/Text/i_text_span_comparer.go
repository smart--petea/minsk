package Text

type ITextSpanComparer interface {
    Compare(*TextSpan, *TextSpan) int
}
