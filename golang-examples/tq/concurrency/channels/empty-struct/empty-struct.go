// The important use of empty struct is to show the developer that we do not have any value. 
// The purpose is purely informational. Some of the examples where the empty struct is useful are as follows:
// While implementing a data set: We can use the empty struct to implement a dataset. 
// Consider an example as shown below.

map_obj := make(map[string]struct{})
for _, value := range []string{"interviewbit", "golang", "questions"} {
   map_obj[value] = struct{}{}
}
fmt.Println(map_obj)

// The output of this code would be:
map[interviewbit:{} golang:{} questions:{}]

// Here, we are initializing the value of a key to an empty struct and initializing the map_obj 
// to an empty struct.

// In graph traversals in the map of tracking visited vertices. 
// For example, consider the below piece of code where we are initializing the value of vertex 
// visited empty struct.
visited := make(map[string]struct{})
for _, isExists := visited[v]; !isExists {
    // First time visiting a vertex.
    visited[v] = struct{}{}
}

// When a channel needs to send a signal of an event without the need for sending any data. 
// From the below piece of code, we can see that we are sending a signal using sending empty struct 
// to the channel which is sent to the workerRoutine.
func workerRoutine(ch chan struct{}) {
    // Receive message from main program.
    <-ch
    println("Signal Received")

    // Send a message to the main program.
    close(ch)
}

func main() {
    //Create channel
    ch := make(chan struct{})
    
    //define workerRoutine
    go workerRoutine(ch)

    // Send signal to worker goroutine
    ch <- struct{}{}

    // Receive a message from the workerRoutine.
    <-ch
    println("Signal Received")
    
}

// The output of the code would be:
// Signal Received
// Signal Received
