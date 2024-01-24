package controller

import (
	"booking-room/shared/common"
	"booking-room/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TrxRsvController struct {
	trxRsvpUC usecase.TrxRsvUsecase
	rg *gin.RouterGroup
}

func NewTrxRsvpController (trxRsvpUC usecase.TrxRsvUsecase, rg *gin.RouterGroup) *TrxRsvController {
	return &TrxRsvController{
		trxRsvpUC: trxRsvpUC,
		rg: rg,
	}
}

func (t *TrxRsvController) Route() {
	t.rg.GET("/list",t.getAll )
	t.rg.GET("/get/:id",t.getID)
	t.rg.GET("/employee/:id",t.getEmployee)
}

func (t *TrxRsvController) getAll(c *gin.Context)  {
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	list, paging, err := t.trxRsvpUC.List(page,size)
	if err != nil{
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var response []interface{}
	for _, v := range list{
		response = append(response, v)
	}
	common.SendPagedResponse(c, response, paging, "success")
}

func (t *TrxRsvController) getID(c *gin.Context)  {
	id := c.Param("id")

	trx, err := t.trxRsvpUC.GetID(id)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "Transcaction with ID : " + id + " not found")
		return
	}
	common.SendSingleResponse(c, trx, "success")
}

func (t *TrxRsvController) getEmployee(c *gin.Context)  {
	id := c.Param("id")
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))

	trx, paging, err := t.trxRsvpUC.GetEmployee(id, page, size)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "Transcaction with Employee ID : " + id + " not found")
		return
	}

	common.SendPagedResponse(c, trx, paging, "success")
}