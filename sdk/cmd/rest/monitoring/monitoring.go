package monitoring

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"

	gnic "github.com/devpablocristo/golang/sdk/pkg/gin-gonic/gin"
	gmw "github.com/devpablocristo/golang/sdk/pkg/go-micro-web"
)

func MonitoringRestAPI(ginInst gnic.GinClientPort, ms gmw.GoMicroClientPort) {
	r := ginInst.GetRouter()

	// TODO: Probar prometheus
	r.GET("/metrics", ginInst.WrapH(promhttp.Handler()))

	// Integrar Go Micro y Gin
	ms.GetService().Handle("/", r)
}
