package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lekan-pvp/gokee/gateway-svc/pkg/auth"
	"github.com/lekan-pvp/gokee/gateway-svc/pkg/config"
	"github.com/lekan-pvp/gokee/gateway-svc/pkg/doc"
	"log"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	doc.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
