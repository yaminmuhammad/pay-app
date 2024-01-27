package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yaminmuhammad/pay-app/config"
	"github.com/yaminmuhammad/pay-app/entity"
	"github.com/yaminmuhammad/pay-app/handler/middleware"
	"github.com/yaminmuhammad/pay-app/shared/common"
	"github.com/yaminmuhammad/pay-app/usecase"
)

type MerchantController struct {
	merchantUC     usecase.MerchantUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (c *MerchantController) registerHandler(ctx *gin.Context) {
	var data entity.Merchant
	if err := ctx.ShouldBindJSON(&data); err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	merchant, err := c.merchantUC.RegisterMerchant(data)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendSuccessResponse(ctx, merchant, "Created")

}

func (c *MerchantController) Route() {
	c.rg.POST(config.MasterMerchantRegister, c.registerHandler)
}

func NewMerchantController(merchantUC usecase.MerchantUseCase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *MerchantController {
	return &MerchantController{
		merchantUC:     merchantUC,
		rg:             rg,
		authMiddleware: authMiddleware,
	}
}
