package Text

type TextLine struct {
        Text *SourceText 
        Start int 
        Length int 
        LengthIncludingLineBreak int
}

func NewTextLine(text *SourceText, start, length, lengthIncludingLineBreak) *TextLine {
    return &TextLine{
        Text: text, 
        Start: start, 
        Length: length, 
        LengthIncludingLineBreak: lengthIncludingLineBreak,
    }
}

func (tl *TextLine) String() string {
    start := tl.Start
    end := tl.Start + tl.Length
    return tl.Text.String()[start:end]
}

func (tl *TextLine) End() int {
    return tl.Start + tl.Length
}

func (tl *TextLine) Span() *TextSpan {
    return NewTextSpan(tl.Start, tl.Length)
}

func (tl *TextLine) SpanIncludingLineBreak() *TextSpan {
    return NewTextSpan(tl.Start, tl.LengthIncludingLineBreak)
}
