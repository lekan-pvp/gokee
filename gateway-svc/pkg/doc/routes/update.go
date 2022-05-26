package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	pb2 "github.com/lekan-pvp/gokee/gateway-svc/pkg/doc/pb"
	"net/http"
)

type UpdateDocRequestBody struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func UpdateDocument(ctx *gin.Context, c pb2.DocServiceClient) {
	body := UpdateDocRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Update(context.Background(), &pb2.UpdateRequest{
		Id:      body.ID,
		Title:   body.Title,
		Content: body.Content,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
