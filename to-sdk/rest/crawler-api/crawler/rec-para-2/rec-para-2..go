package recpara2

func recFunc(input string, outputChannel chan, wg &sync.WaitGroup) {
    defer wg.Done()
    for subInput := range getSubInputs(input) {
        wg.Add(1)
        go recFunc(subInput)
    }
    outputChannel <-getOutput(input)
}

func rec2() {
    outputChannel := make(chan []string)
    waitGroup := sync.WaitGroup{}

    waitGroup.Add(1)
    go recFunc("some_input", outputChannel, &waitGroup)

    result := []string{}
    go func() {
     for nextResult := range outputChannel {
      result = append(result, nextResult ...)
     }
    }
    waitGroup.Wait()
    close(outputChannel)        
}