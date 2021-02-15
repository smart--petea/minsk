package Text

import (
    "testing"
    "minsk/CodeAnalysis/Text"
)


func TestSourceTextIncludesLastLine(t *testing.T) {
    tests := []struct{
        Text string
        ExpectedLines int
    }{
        {
            Text: ".",
            ExpectedLines: 1,
        },
        {
            Text: ".\r\n",
            ExpectedLines: 2,
        },
        {
            Text: ".\r\n\r\n",
            ExpectedLines: 3,
        },
    }

    for _, test := range tests {
        sourceText := Text.SourceTextFrom(test.Text)
        lenLines := len(sourceText.Lines)
        if lenLines != test.ExpectedLines {
            t.Errorf("Lines('%s')=%d, expected=%d", test.Text, lenLines, test.ExpectedLines)
        }
    }
}
