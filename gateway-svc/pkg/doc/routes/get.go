package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	pb2 "github.com/lekan-pvp/gokee/gateway-svc/pkg/doc/pb"
	"net/http"
)

type GetDocRequestBody struct {
	ID string `json:"id"`
}

func GetDocument(ctx *gin.Context, c pb2.DocServiceClient) {
	body := GetDocRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Get(context.Background(), &pb2.GetRequest{
		Id: body.ID,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(http.StatusOK, &res)
}
