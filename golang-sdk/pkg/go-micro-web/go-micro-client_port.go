package gomicro

import "go-micro.dev/v4/web"

type GoMicroClientPort interface {
	Start() error
	Stop() error
	GetService() web.Service
}
