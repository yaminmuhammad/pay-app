package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/yaminmuhammad/pay-app/dto"
	"github.com/yaminmuhammad/pay-app/entity"
	"github.com/yaminmuhammad/pay-app/shared/common"
	"github.com/yaminmuhammad/pay-app/usecase"
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
	rsp, err := a.authUC.Login(payload, ctx)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendCreatedResponse(ctx, rsp, "Ok")
}

func (a *AuthController) logoutHandler(ctx *gin.Context) {
	customerIdInterface, exist := ctx.Get("customerId")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: CustomerId not found in context"})
		return
	}

	customerId, ok := customerIdInterface.(string)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid CustomerId type"})
		return
	}

	activity := entity.Activities{
		CustomerId:   customerId,
		Activity:     "Logout",
		ActivityTime: time.Now(),
	}

	if err := a.authUC.LogActivity(activity); err != nil {
		log.Printf("Error logging activity: %v \n", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error: Unable to log activity"})
		return
	}

	session := sessions.Default(ctx)
	session.Clear()
	session.Save()

	// Respon JSON untuk memberi tahu keberhasilan logout
	ctx.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

func (a *AuthController) Route() {
	a.rg.POST("/auth/login", a.loginHandler)
	a.rg.POST("/auth/logout", a.logoutHandler)
}

func NewAuthController(authUC usecase.AuthUseCase, rg *gin.RouterGroup) *AuthController {
	return &AuthController{authUC: authUC, rg: rg}
}
