package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	pb2 "github.com/lekan-pvp/gokee/gateway-svc/pkg/doc/pb"
	"net/http"
)

type AddDocRequestBody struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  string `json:"userId"`
}

func AddDocument(ctx *gin.Context, c pb2.DocServiceClient) {
	body := AddDocRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId, _ := ctx.Get("userId")

	res, err := c.Add(context.Background(), &pb2.AddRequest{
		Title:   body.Title,
		Content: body.Content,
		UserId:  userId.(string),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
