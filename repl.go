package minsk

import (
    "fmt"
    "bufio"
    "os"
    "os/exec"
    "log"
    "strings"

    "minsk/Util/Console"
)

type Repl struct {
    reader *bufio.Reader
    submissionText string

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
    r.submissionText = ""
    document := NewObservableCollection("")
    view := NewSubmissionView(document)

    for r.submissionText == "" {
        key := ConsoleReadKey()
        r.handleKey(key, document, view)
    }

    return r.submissionText
}

func ConsoleInit() {
    //disable input buffering
    exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
    //do not display entered characters on the screen
    exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
}

func ConsoleReadKey() *ConsoleKeyInfo {
    b := make([]byte, 8)
    size, _ := os.Stdin.Read(b)

    return NewConsoleKeyInfo(b[:size])
}

type ConsoleKeyInfo struct {
    Bytes []byte
}

func NewConsoleKeyInfo(bytes []byte) *ConsoleKeyInfo {
    return &ConsoleKeyInfo{
        Bytes: bytes,
    }
}

func (c *ConsoleKeyInfo) Kind() ConsoleKey {
    if len(c.Bytes) == 3 && c.Bytes[0] == 27 && c.Bytes[1] == 91 {
        switch c.Bytes[3] {
        case 10:
            return Enter
        case 67:
            return LeftArrow
        case 68:
            return RightArrow
        case 65:
            return UpArrow
        case 66:
            return DownArrow
        default:
            panic(fmt.Sprintf("Unknown console command %+v", c.Bytes[:3]))
        }
    }

    return Symbol
}

type ConsoleKey string

const (
        Enter ConsoleKey = "Enter"
        LeftArrow ConsoleKey = "LeftArrow"
        RightArrow ConsoleKey = "RightArrow"
        UpArrow ConsoleKey = "UpArrow"
        DownArrow ConsoleKey = "DownArrow"
        Symbol ConsoleKey = "Symbol"
)

type ObservableCollection struct {
    Elements []string
}

func (o *ObservableCollection) Add(el string) {
    o.Elements = append(o.Elements, el)
    //todo callback
}

func (o *ObservableCollection) GetElement(index int) string {
    return o.Elements[index]
}

func (o *ObservableCollection) SetElement(index int, val string) {
    o.Elements[index] = val
}

func (o *ObservableCollection) Count() int {
    return len(o.Elements)
}

func NewObservableCollection(elements... string) *ObservableCollection {
    return &ObservableCollection{
        Elements: elements,
    }
}

func (r *Repl) handleKey(key *ConsoleKeyInfo, document *ObservableCollection, view *SubmissionView) {
        switch key.Kind() {
        case Enter:
            r.HandleEnter(document, view)
        case LeftArrow:
            r.HandleLefttArrow(document, view)
        case RightArrow:
            r.HandleRightArrow(document, view)
        case UpArrow:
            r.HandleUpArrow(document, view)
        case DownArrow:
            r.HandleDownArrow(document, view)
        default:
            r.HandleTyping(document, view, string(key.Bytes))
        }
}

func (r *Repl) HandleEnter(document *ObservableCollection, view *SubmissionView) {
    lines := document.Elements
    submissionText :=  strings.Join(lines, "\n")
    if r.IsCompleteSubmission(submissionText) {
        r.submissionText = submissionText
        return
    }

    document.Add("")
    view.SetCurrentCharacter(0)
    view.SetCurrentLineIndex(document.Count() - 1)
}

func (r *Repl) HandleLefttArrow(document *ObservableCollection, view *SubmissionView) {
    currentCharacter := view.GetCurrentCharacter()
    if currentCharacter > 0 {
        view.SetCurrentCharacter(currentCharacter - 1)
    }
}

func (r *Repl) HandleRightArrow(document *ObservableCollection, view *SubmissionView){
    line := document.GetElement(view.GetCurrentLineIndex())

     currentCharacter := view.GetCurrentCharacter()
    if currentCharacter < len(line) - 1 {
        view.SetCurrentCharacter(currentCharacter + 1)
    }
}

func (r *Repl) HandleUpArrow(document *ObservableCollection, view *SubmissionView) {
    currentLineIndex := view.GetCurrentLineIndex()
    if currentLineIndex > 0 {
        view.SetCurrentLineIndex(currentLineIndex - 1)
    }
}

func (r *Repl) HandleDownArrow(document *ObservableCollection, view *SubmissionView) {
    currentLineIndex := view.GetCurrentLineIndex()
    if currentLineIndex < document.Count() - 1 {
        view.SetCurrentLineIndex(currentLineIndex + 1)
    }
}

func (r *Repl) HandleTyping(document *ObservableCollection, view *SubmissionView, text string) {
    lineIndex := view.GetCurrentLineIndex()
    start := view.GetCurrentCharacter()

    line := document.GetElement(lineIndex)
    line =  line[:start] + text + line[start:]

    document.SetElement(lineIndex, line)
    currentCharacter := view.GetCurrentCharacter() 
    currentCharacter = currentCharacter + len(text)
    view.SetCurrentCharacter(currentCharacter)
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
    SubmissionDocument *ObservableCollection

    cursorTop int
    renderedLineCount int
}

func NewSubmissionView(submissionDocument *ObservableCollection) *SubmissionView {
    s := &SubmissionView{
        SubmissionDocument: submissionDocument,
        cursorTop: Util.Console.CursorTop,
    }

    submissionDocument.CollectionChanged(s.SubmissionDocumentChanged)
    s.Render()

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
        blankLine := strings.Repeat(" ", ConsoleWindowWidth())
        for numberOfBlankLines > 0 {
            fmt.Println(blankLine)
        }
    }

    s.renderedLineCount = lineCount
    Console.CursorVisibile = true //todo C#
    s.UpdateCursorPosition() 
}

func ConsoleWindowWidth() int {
    cmd := exec.Command("stty", "size")
    cmd.Stdin = os.Stdin
    out, err := cmd.Output()
    if err != nil {
        panic(err)
    }

    fields := strings.Field(string(out))
    return fields[1]
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
