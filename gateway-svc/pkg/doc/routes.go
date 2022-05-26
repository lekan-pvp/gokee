package doc

import (
	"github.com/gin-gonic/gin"
	auth2 "github.com/lekan-pvp/gokee/gateway-svc/pkg/auth"
	"github.com/lekan-pvp/gokee/gateway-svc/pkg/config"
	routes2 "github.com/lekan-pvp/gokee/gateway-svc/pkg/doc/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth2.ServiceClient) {
	a := auth2.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/doc")
	routes.Use(a.AuthRequired)
	routes.POST("/add", svc.AddDocument)
	routes.GET("/get", svc.GetDocument)
	routes.POST("/update", svc.UpdateDocument)
	routes.POST("/remove", svc.RemoveDocument)
}

func (svc *ServiceClient) AddDocument(ctx *gin.Context) {
	routes2.AddDocument(ctx, svc.Client)
}

func (svc *ServiceClient) GetDocument(ctx *gin.Context) {
	routes2.GetDocument(ctx, svc.Client)
}

func (svc *ServiceClient) UpdateDocument(ctx *gin.Context) {
	routes2.UpdateDocument(ctx, svc.Client)
}

func (svc *ServiceClient) RemoveDocument(ctx *gin.Context) {
	routes2.RemoveDocument(ctx, svc.Client)
}
