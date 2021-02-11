package Text

type SourceText struct{
    Lines []*TextLine
}

func newSourceText(text string) *SourceText {
    return &SourceText{
        lines: ParseLines(text)
    }
}

func From(text string) *SourceText {
    return newSourceText(text)
}

func ParseLines(sourceText *SourceText, text string) []*TextLine {
    var result []*TextLine

    var position, lineStart int
    for position < len(text) {
        lineBreakWidth := GetLineBreakWidth(text, position)
        if lineBreakWidth == 0 {
            position += 1
        } else {
            lineLength = position - lineStart
            lineLengthIncludingLineBreak = lineLength + lineBreakWidth
            line := NewTextLine(sourceText, lineStart, lineLength, lineLengthIncludingLineBreak)

            result = append(result, line)

            position += lineBreakWidth
            lineStart = position
        }
    }

    return result
}

func GetLineBreakWidth(text string, i int) int {
    c := text[i]

    l := '\0'
    if (i + 1) > len(text) {
        l := text[i+1]
    }

    if c == '\r' && l == '\n' {
        return 2
    }

    if c == '\r' || c == '\n' {
        return 1
    }

    return 0
}
