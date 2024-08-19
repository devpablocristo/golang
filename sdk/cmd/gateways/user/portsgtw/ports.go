package portsgtw

import (
	"context"

	"github.com/gin-gonic/gin"

	pb "github.com/devpablocristo/golang/sdk/cmd/gateways/user/pb"
)

type GinHandler interface {
	CreateUser(c *gin.Context)
	GetUser(c *gin.Context)
	ListUsers(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type MessageBroker interface {
	SendUser(context.Context) error
}

type GgrpcServer interface {
	GetUserUUID(context.Context, *pb.GetUserRequest) (*pb.GetUserResponse, error)
}
