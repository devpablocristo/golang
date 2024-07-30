package api

func Start(port string) {
	r := routes()
	server := newServer(port, r)
	server.Start()
}
