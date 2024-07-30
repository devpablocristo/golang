Here's the final solution you can try out in case no other solution was helpful to you. This one's applicable and useful in some cases and could possiblty be of some help. No worries if you're unsure about it but I'd recommend going through it.

I can start a separate go function for this, but how do I ever quit this separate go function?

You can range over the output channel in the separate go-routine. The go-routine, in that case, will exit safely, when the channel is closed

go func() {
   for nextResult := range outputChannel {
     result = append(result, nextResult ...)
   }
}
So, now the thing that we need to take care of is that the channel is closed after all the go-routines spawned as part of the recursive function call have successfully existed

For that, you can use a shared waitgroup across all the go-routines and wait on that waitgroup in your main function, as you are already doing. Once the wait is over, close the outputChannel, so that the other go-routine also exits safely

func recFunc(input string, outputChannel chan, wg &sync.WaitGroup) {
    defer wg.Done()
    for subInput := range getSubInputs(input) {
        wg.Add(1)
        go recFunc(subInput)
    }
    outputChannel <-getOutput(input)
}

func main() {
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
PS: If you want to have bounded parallelism to limit the exponential growth, check this out


Final Words
Go really fits well for performance-oriented cloud software. The popular DevOps tools have been written in Go, such as Docker, and also the open-source container orchestration system Kubernetes.. These were a few of many solutions that were found helpful for your issue. Hope it turns out helpful for you. Please upvote the solutions if it worked for you.