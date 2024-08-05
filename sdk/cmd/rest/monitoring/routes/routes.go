package monitoring

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"

	wire "github.com/devpablocristo/golang/sdk/cmd/rest"
	gnic "github.com/devpablocristo/golang/sdk/pkg/gin-gonic/gin"
	gmw "github.com/devpablocristo/golang/sdk/pkg/go-micro-web"
	is "github.com/devpablocristo/golang/sdk/pkg/init-setup"
)

func Routes(ginInst gnic.GinClientPort, ms gmw.GoMicroClientPort) {
	r := ginInst.GetRouter()

	handler, err := wire.InitializeMonitoring()
	if err != nil {
		is.MicroLogError("userHandler error: %v", err)
	}	

	// Ruta de Salud
	r.GET("/health", handler.Health)
	r.GET("/ping", handler.Ping)

	// TODO: Probar prometheus
	r.GET("/metrics", ginInst.WrapH(promhttp.Handler()))

	// Integrar Go Micro y Gin
	ms.GetService().Handle("/", r)
}
