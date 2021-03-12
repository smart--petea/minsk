package minsk

import (
    "fmt"
    "bufio"
    "os"
    "log"
    "strings"

    "minsk/Util/Console"
)

type Repl struct {
    reader *bufio.Reader

    EvaluateSubmission func(text string) 
    IsCompleteSubmission func(text string) bool
}

func (r *Repl) getReader() *bufio.Reader {
    if r.reader == nil {
        r.reader = bufio.NewReader(os.Stdin)
    }

    return r.reader
}

func (r *Repl) GetInput() (string, error) {
    return r.getReader().ReadString('\n')
}

func (r *Repl) Run(ir IRepl) {
    for {
        text := r.editSubmission(ir)
        if text == "" {
            return
        }
        ir.EvaluateSubmission(text)
    }
}

func (r *Repl) editSubmission(ir IRepl) string {
    document := Util.NewObservableCollection()
    view := NewSubmissionView(document)

    for {
        key := Console.ReadKey(true) //todo c#
        r.handleKey(key, document, view)
    }
}

func (r *Repl) handleKey(key *ConsoleKeyInfo, document *Util.ObservableCollection, view *SubmissionView) {
    if key.Modifiers == default(ConsoleModifiers) { //todo c#
        switch key {
        case ConsoleKey.Enter:
            r.HandleEnter(document, view)
        case ConsoleKey.LeftArrow:
            r.HandleLefttArrow(document, view)
        case ConsoleKey.RightArrow:
            r.HandleRightArrow(document, view)
        case ConsoleKey.UpArrow:
            r.HandleUpArrow(document, view)
        case ConsoleKey.DownArrow:
            r.HandleDownArrow(document, view)
        default:
            if key.KeyChar > ' ' {//todo c#
                r.HandleTyping(document, view, key.KeyChar)
            }

        }
    }
}

func (r *Repl) HandleEnter(document *Util.ObservableCollection, view *SubmissionView) {
}
func (r *Repl) HandleLefttArrow(document *Util.ObservableCollection, view *SubmissionView) {
}
func (r *Repl) HandleRightArrow(document *Util.ObservableCollection, view *SubmissionView){
}
func (r *Repl) HandleUpArrow(document *Util.ObservableCollection, view *SubmissionView) {
}
func (r *Repl) HandleDownArrow(document *Util.ObservableCollection, view *SubmissionView) {
}
func (r *Repl) HandleTyping(document *Util.ObservableCollection, view *SubmissionView, text string) {
    lineIndex := view.CurrentLineIndex
    start := view.CurrentCharacter
    text := kkk

    document[lineIndex] = document[lineIndex].Insert(start, text)
    view.CurrentCharacter = view.CurrentCharacter + 1
}

func (r *Repl) editSubmissionOld(ir IRepl) string {
    var textBuilder strings.Builder

    for {
        Console.ForegroundColour(Console.COLOUR_GREEN)
        if textBuilder.Len() == 0 {
            fmt.Print("» ")
        } else {
            fmt.Print("· ")
        }
        Console.ResetColour()

        input, err := r.GetInput()
        if err != nil {
            log.Fatal(err)
        }
        isBlank := len(strings.TrimSpace(input)) == 0

        if textBuilder.Len() == 0 {
            if isBlank {
                return ""
            } 
            
            if (strings.HasPrefix(input, "#")) {
                ir.EvaluateMetaCommand(input)
                continue
            }
        }

        textBuilder.WriteString(input)
        text := textBuilder.String()

        if !ir.IsCompleteSubmission(text) {
            continue
        }

        return text
    }
}

func (r *Repl) EvaluateMetaCommand(input string) {
    Console.ForegroundColour(Console.COLOUR_RED)
    fmt.Printf("Invalid command %s.", input)
    Console.ResetColour()
}

type SubmissionView struct {
    _currentLineIndex int
    _currentCharacter int
    SubmissionDocument *Util.ObservableCollection

    cursorTop int
    renderedLineCount int
}

func NewSubmissionView(submissionDocument *Util.ObservableCollection) *SubmissionView {
    s := &SubmissionView{
        SubmissionDocument: submissionDocument,
        cursorTop: Util.Console.CursorTop,
    }

    submissionDocument.CollectionChanged(s.SubmissionDocumentChanged)

    return s
}

func (s *SubmissionView) SubmissionDocumentChanged(sender interface{}, e *Util.NotifyCollectionChangedEventArgs) {
    s.Render()
}

func (s *SubmissionView) Render() {
    Console.SetCursorPosition(0, s.cursorTop) //todo C#
    Console.CursorVisibile = false //todo C#

    var lineCount int
    for _, line := range s.submissionDocument {
        Console.ForegroundColour(Console.COLOUR_GREEN)
        if lineCount == 0 {
            fmt.Print("» ")
        } else {
            fmt.Print("· ")
        }

        Console.ResetColour()

        fmt.Println(line)
        lineCount = lineCount + 1
    }

    numberOfBlankLines := s.renderedLineCount - lineCount
    if numberOfBlankLines > 0 {
        blankLine := new string( ' ', Console.WindowWidth)
        for numberOfBlankLines > 0 {
            fmt.Println(blankLine)
        }
    }

    s.renderedLineCount = lineCount
    Console.CursorVisibile = true //todo C#
    s.UpdateCursorPosition() 
}

func (s *SubmissionView) UpdateCursorPosition() {
    Console.CursorTop = s.cursorTop + s.GetCurrentLineIndex()
    Console.CursorLeft = 2 + s.CurrentCharacter
}

func (s *SubmissionView) GetCurrentLineIndex() int {
    return s._currentLineIndex 
}

func (s *SubmissionView) SetCurrentLineIndex(value int) int {
    if value != s._currentLineIndex {
        s._currentLineIndex  = value
        s.UpdateCursorPosition()
    }
}

func (s *SubmissionView) GetCurrentCharacter() int {
    return s._currentCharacter 
}

func (s *SubmissionView) SetCurrentCharacter(value int) int {
    if value != s._currentCharacter {
        s._currentCharacter  = value
        s.UpdateCursorPosition()
    }
}
