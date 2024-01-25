package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yaminmuhammad/pay-app/handler/middleware"
	"github.com/yaminmuhammad/pay-app/usecase"
)

type CustomerController struct {
	customerUC     usecase.CustomerUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (c *CustomerController) getHandler(ctx *gin.Context)  {
	id := ctx.Param("id")
	customer, err := c.customerUC.
}

func (c *CustomerController) Route() {
	c.rg.GET("/customers/:id", c.authMiddleware.RequireToken("customer"), c.getHandler)
}

func NewCustomerController(customerUC usecase.CustomerUseCase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *CustomerController{
	return  &CustomerController{customerUC: customerUC, rg: rg, authMiddleware: authMiddleware}
}
