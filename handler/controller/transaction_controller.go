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

type TransactionController struct {
	transactionUC  usecase.TransactionUsecase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (t *TransactionController) CreateTransaction(c *gin.Context) {

	var data entity.Transactions

	if err := c.ShouldBindJSON(&data); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	data, err := t.transactionUC.CreateTransaction(data)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSuccessResponse(c, data, "Created")
}

func (t *TransactionController) Route() {
	t.rg.POST(config.MasterTransaction, t.authMiddleware.RequireToken("customer"), t.CreateTransaction)
}

func NewTransactionController(transactionUC usecase.TransactionUsecase, rg *gin.RouterGroup, auth middleware.AuthMiddleware) *TransactionController {
	return &TransactionController{
		transactionUC:  transactionUC,
		rg:             rg,
		authMiddleware: auth,
	}
}
