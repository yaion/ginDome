package common

import (
	"ginDome/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Verify(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.Verify())
}
