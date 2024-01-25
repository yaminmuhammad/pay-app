package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yaminmuhammad/pay-app/dto"
	"github.com/yaminmuhammad/pay-app/shared/common"
	"github.com/yaminmuhammad/pay-app/usecase"
	"net/http"
)

type AuthController struct {
	authUC usecase.AuthUseCase
	rg     *gin.RouterGroup
}

func (a *AuthController) loginHandler(ctx *gin.Context) {
	var payload dto.AuthRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	rsp, err := a.authUC.Login(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendCreatedResponse(ctx, rsp, "Ok")
}

func (a *AuthController) Route() {
	a.rg.POST("/auth/login", a.loginHandler)
}

func NewAuthController(authUC usecase.AuthUseCase, rg *gin.RouterGroup) *AuthController {
	return &AuthController{authUC: authUC, rg: rg}
}
