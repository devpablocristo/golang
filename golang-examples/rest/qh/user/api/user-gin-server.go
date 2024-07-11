package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

func GinServer(port string, router *gin.Engine) {
	log.Println("starting gin server")

	router.Run("localhost:" + port)

	// httpServer := &http.Server{
	// 	Addr:         ":" + port,
	// 	Handler:      routes,
	// 	ReadTimeout:  5 * time.Second,
	// 	WriteTimeout: 10 * time.Second,
	// 	IdleTimeout:  15 * time.Second,
	// }

	// go func() {
	// 	err := httpServer.ListenAndServe()
	// 	if err != nil {
	// 		if err != http.ErrServerClosed {
	// 			log.Fatalf("could not listen on %s due to %s", httpServer.Addr, err.Error())
	// 		}
	// 	}
	// }()

	//log.Printf("the chi server is ready to handle requests %s", httpServer.Addr)
}
