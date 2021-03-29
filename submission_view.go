package minsk

import (
    "minsk/Util"
    "minsk/Util/Console"
    "log"
    "fmt"
    "strings"
)

type SubmissionView struct {
    _currentLine int
    _currentCharacter int
    SubmissionDocument *Util.ObservableCollection
    lineRenderer func(string)

    cursorTop int
    renderedLineCount int
}

func NewSubmissionView(lineRenderer func(string), submissionDocument *Util.ObservableCollection) *SubmissionView {
    top, _ := Console.GetCursorPos()
    s := &SubmissionView{
        SubmissionDocument: submissionDocument,
        cursorTop: top,
        lineRenderer: lineRenderer,
    }

    submissionDocument.CollectionChanged(s.SubmissionDocumentChanged)
    s.Render()

    return s
}

func (s *SubmissionView) SubmissionDocumentChanged(sender interface{}) {
    s.Render()
}

func (s *SubmissionView) Render() {
    left := 1
    top := s.cursorTop
    log.Printf("SubmissionView.Render left=%d top=%d ", left, top)

    Console.SetCursorVisibile(false)

    var lineCount int
    for _, line := range s.SubmissionDocument.Collection {
        Console.SetCursorPos(left, s.cursorTop + lineCount)
        Console.ForegroundColour(Console.COLOUR_GREEN)
        if lineCount == 0 {
            fmt.Print("»")
        } else {
            fmt.Print("·")
        }

        Console.ResetColour()
        s.lineRenderer(line)

        lineCount = lineCount + 1
    }

    numberOfBlankLines := s.renderedLineCount - lineCount
    if numberOfBlankLines > 0 {
        blankLine := strings.Repeat(" ", Console.WindowWidth())
        for i := 0; i < numberOfBlankLines; i = i+1 {
            Console.SetCursorPos(0, s.cursorTop + lineCount + i)
            fmt.Println(blankLine)
        }
    }

    s.renderedLineCount = lineCount
    Console.SetCursorVisibile(true)
    s.UpdateCursorPosition() 
}

func (s *SubmissionView) UpdateCursorPosition() {
    top := s.cursorTop + s.GetCurrentLine()
    left := 2 + s._currentCharacter

    //log.Printf("UpdateCursorPosition cursorTop=%d lineIndex=%d currentCharacter=%d left=%d top=%d", s.cursorTop, s.GetCurrentLine(), s._currentCharacter, left, top)

    Console.SetCursorPos(left, top)
}

func (s *SubmissionView) GetCurrentLine() int {
    return s._currentLine
}

func (s *SubmissionView) SetCurrentLine(value int) {
    //log.Printf("SetCurrentLine old=%d new=%d", s._currentLine, value)
    if value != s._currentLine {
        s._currentLine  = value

        lineLen := len(s.SubmissionDocument.Collection[s._currentLine])
        if s._currentCharacter > lineLen {
            s._currentCharacter  = lineLen
        }

        s.UpdateCursorPosition()
    }
}

func (s *SubmissionView) GetCurrentCharacter() int {
    return s._currentCharacter 
}

func (s *SubmissionView) SetCurrentCharacter(value int) {
    //log.Printf("SetCurrentCharacter old=%d new=%d", s._currentCharacter, value)
    if value != s._currentCharacter {
        s._currentCharacter  = value
        s.UpdateCursorPosition()
    }
}
