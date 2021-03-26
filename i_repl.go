package minsk

type IRepl interface {
    EvaluateMetaCommand(input string)
    EvaluateSubmission(text string) 
    IsCompleteSubmission(text string) bool 
    RenderLine(string)
}
