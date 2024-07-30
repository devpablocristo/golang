1. Basic HTML: we use jet a library 
2. We use 3 external tools:
   1. routing: bmizerany/pat
   2. template engine: CloudyKit
   3. websocket: gorilla websocket
3. renderPage renders a jet template, creates pages.
   1. We use 3 arguments: w (http.ResponseWriter), tmpl (string), data (jet.VarMap) <- can be empty

---

1. Somebody connects to the webpage: func Home <- display the webpage
2. func WsEndpoint is called to connect the WS
3. When we connect to WS:
   1. go ListenForWs(&conn) <- is called
   2. go ListenForWs(&conn) <- is a goroutine on a infite loop
   3. wsChan <- payload <- everything sent on the payload is pass to wsChan
4. func ListenToWsChannel() <- everything on the the payload is store on the variable e
   1. Here is where we process the payload, the logic, where we stract the username or the message.
5. func broadcastToAll <- send the payload to all the clients