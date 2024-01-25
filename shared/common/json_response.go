package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendCreatedResponse(ctx *gin.Context, data interface{}, message string) {
	ctx.JSON(http.StatusCreated, &model)
}
