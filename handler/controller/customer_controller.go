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

type CustomerController struct {
	customerUC     usecase.CustomerUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (c *CustomerController) getHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	customer, err := c.customerUC.FindCustomerByID(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusNotFound, "customer with ID "+id+"not found")
		return
	}
	common.SendSuccessResponse(ctx, customer, "OK")
}

func (c *CustomerController) registerHandler(ctx *gin.Context) {
	var data entity.Customer
	if err := ctx.ShouldBindJSON(&data); err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	customer, err := c.customerUC.RegisterCustomer(data)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendSuccessResponse(ctx, customer, "Created")

}

func (c *CustomerController) Route() {
	c.rg.GET(config.MasterCustomerGetById, c.authMiddleware.RequireToken("customer"), c.getHandler)
	c.rg.POST(config.MasterCustomerRegister, c.registerHandler)
}

func NewCustomerController(customerUC usecase.CustomerUseCase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *CustomerController {
	return &CustomerController{customerUC: customerUC, rg: rg, authMiddleware: authMiddleware}
}
