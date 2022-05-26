package auth

import (
	"github.com/gin-gonic/gin"
	routes2 "github.com/lekan-pvp/gokee/gateway-svc/pkg/auth/routes"
	"github.com/lekan-pvp/gokee/gateway-svc/pkg/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/auth-svc")
	routes.POST("/register", svc.Register)
	routes.POST("/login", svc.Login)

	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes2.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes2.Login(ctx, svc.Client)
}
