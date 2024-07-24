package monitoring

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"

	gingonic "github.com/devpablocristo/qh/events/pkg/gin-gonic/gin"
	gmw "github.com/devpablocristo/qh/events/pkg/go-micro-web"
)

func MonitoringRestAPI(ginInst gingonic.GinClientPort, ms gmw.GoMicroClientPort) {
	r := ginInst.GetRouter()

	// TODO: Probar prometheus
	r.GET("/metrics", ginInst.WrapH(promhttp.Handler()))

	// Integrar Go Micro y Gin
	ms.GetService().Handle("/", r)

}
