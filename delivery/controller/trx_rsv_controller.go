package controller

import (
	"booking-room/model/dto"
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
	t.rg.GET("/",t.getAll )
	t.rg.GET("/get/:id",t.getID)
	t.rg.GET("/employee/:id",t.getEmployee)
	t.rg.GET("/approval",t.getApprove)
	t.rg.PUT("/approval", t.acceptRSVP)
	t.rg.POST("/", t.createRSVP)
	t.rg.PUT("/:id", t.editRSVP)
	t.rg.DELETE("/:id", t.deleteRSVP)
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
	common.SendSuccessResponse(c, http.StatusOK,trx)
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

func (t *TrxRsvController) createRSVP(c *gin.Context)  {
	var payload dto.PayloadReservationDTO
	if err := c.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if payload.Id != "" || payload.Email == "" || payload.RoomCode == "" || payload.Note == "" || payload.StartDate == nil || payload.EndDate == nil{
		common.SendErrorResponse(c, http.StatusBadRequest, "All field must be filled")
		return
	}

	trx, err := t.trxRsvpUC.PostReservation(payload)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	common.SendSuccessResponse(c, http.StatusCreated, trx)
}

func (t *TrxRsvController) acceptRSVP(c *gin.Context)  {
	// id := c.Param("id")
	var acc dto.TransactionDTO
	if err := c.ShouldBindJSON(&acc); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if acc.Id == "" && acc.ApproveStatus == "" {
		common.SendErrorResponse(c, http.StatusBadRequest, "accesment field failed")
		return
	}

	a, err := t.trxRsvpUC.UpdateStatus(acc)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	common.SendCreatedResponse(c, a, "updated")
}

func (t *TrxRsvController) deleteRSVP(c *gin.Context)  {
	id := c.Param("id")

	del, err := t.trxRsvpUC.DeleteResv(id)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "Transcaction with ID : " + id + " not found")
		return
	}
	common.SendSingleResponse(c, del, "success")
}

func (t *TrxRsvController) getApprove(c *gin.Context)  {
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	list, paging, err := t.trxRsvpUC.GetApprovalList(page,size)
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

func (t *TrxRsvController) editRSVP(c *gin.Context)  {
	id := c.Param("id")

	trx, err := t.trxRsvpUC.GetID(id)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "Transcaction with ID : " + id + " not found")
		return
	}
	
	var payload dto.PayloadReservationDTO
	payload.Id = trx.Id
	
	if err := c.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if payload.Email != ""{
		common.SendErrorResponse(c, http.StatusBadRequest, "Cant change user reservation")
		return
	}
	if payload.RoomCode == "" || payload.Note == "" || payload.StartDate == nil || payload.EndDate == nil{
		common.SendErrorResponse(c, http.StatusBadRequest, "All field must be filled")
		return
	}
	// trx, err := t.trxRsvpUC.UpdateResv(payload)
	// if err != nil {
	// 	common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
	// 	return
	// }
	// common.SendCreatedResponse(c, trx, "updated")

	common.SendSuccessResponse(c, http.StatusOK, payload)
}
