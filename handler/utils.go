package handler

import (
	"net/http"

	"direxplay.com/game-portal-backend-go/core"
	"github.com/gin-gonic/gin"
)

func SendMessage(ctx *gin.Context, code string, message string) {
	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "code": code, "message": message})
}

func SendData(ctx *gin.Context, code string, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "code": code, "data": data})
}

func SendError(ctx *gin.Context, status int, code string, message string) {
	ctx.JSON(status, gin.H{"status": status, "code": code, "error": message})
}

func SendUnexptectedError(ctx *gin.Context) {
	SendError(ctx, http.StatusInternalServerError, core.Code_InternalServerError, "Server encountered unexpted error")
}
