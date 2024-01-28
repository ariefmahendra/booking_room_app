package controller

import (
	"booking-room/delivery/middleware"
	"booking-room/model/dto"
	"booking-room/shared/common"
	"booking-room/usecase"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type TrxRsvController struct {
	trxRsvpUC  usecase.TrxRsvUsecase
	middleware *middleware.Middleware
	rg         *gin.RouterGroup
}

func NewTrxRsvpController(trxRsvpUC usecase.TrxRsvUsecase, middleware *middleware.Middleware, rg *gin.RouterGroup) *TrxRsvController {
	return &TrxRsvController{
		trxRsvpUC:  trxRsvpUC,
		middleware: middleware,
		rg:         rg,
	}
}

func (t *TrxRsvController) Route() {
	t.rg.GET("/", t.getAll)
	t.rg.GET("/get/:id", t.getID)
	t.rg.GET("/employee/:id", t.getEmployee)
	t.rg.GET("/approval", t.getApprove)
	t.rg.PUT("/approval", t.acceptRSVP)
	t.rg.POST("/", t.createRSVP)
	t.rg.DELETE("/:id", t.deleteRSVP)
	t.rg.GET("/available", t.getAvailable)
}

func (t *TrxRsvController) getAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	list, paging, err := t.trxRsvpUC.List(page, size)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var response []interface{}
	for _, v := range list {
		response = append(response, v)
	}
	common.SendSuccessPagedResponse(c, http.StatusOK, response, paging)
}

func (t *TrxRsvController) getID(c *gin.Context) {
	id := c.Param("id")

	trx, err := t.trxRsvpUC.GetID(id)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "Transcaction with ID : "+id+" not found")
		return
	}
	common.SendSuccessResponse(c, http.StatusOK, trx)
}

func (t *TrxRsvController) getEmployee(c *gin.Context) {
	claims := t.middleware.GetUser(c)
	if ok := common.AuthorizationGaAdmin(claims); ok == false {
		common.SendErrorResponse(c, http.StatusForbidden, "Forbidden")
		return
	}

	id := c.Param("id")
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))

	trx, paging, err := t.trxRsvpUC.GetEmployee(id, page, size)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "Transcaction with Employee ID : "+id+" not found")
		return
	}

	common.SendSuccessPagedResponse(c, http.StatusOK, trx, paging)
}

func (t *TrxRsvController) createRSVP(c *gin.Context) {
	var payload dto.PayloadReservationDTO
	if err := c.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if payload.Id != "" || payload.Email == "" || payload.RoomCode == "" || payload.Note == "" || payload.StartDate == nil || payload.EndDate == nil {
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

func (t *TrxRsvController) acceptRSVP(c *gin.Context) {
	claims := t.middleware.GetUser(c)
	if ok := common.AuthorizationGa(claims); ok == false {
		common.SendErrorResponse(c, http.StatusForbidden, "Forbidden")
		return
	}

	var acc dto.TransactionDTO
	if err := c.ShouldBindJSON(&acc); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if acc.Id == "" && acc.ApproveStatus == "" {
		common.SendErrorResponse(c, http.StatusBadRequest, "accesment field failed")
		return
	}

	acc.ApproveStatus = strings.ToUpper(acc.ApproveStatus)
	if acc.ApproveStatus != "ACCEPT" && acc.ApproveStatus != "PENDING" && acc.ApproveStatus != "DECLINE" {
		common.SendErrorResponse(c, http.StatusBadRequest, "wrong approval status")
		return
	}

	a, err := t.trxRsvpUC.UpdateStatus(acc)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	common.SendSuccessResponse(c, http.StatusCreated, a)
}

func (t *TrxRsvController) deleteRSVP(c *gin.Context) {

	id := c.Param("id")

	del, err := t.trxRsvpUC.DeleteResv(id)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "Transcaction with ID : "+id+" not found")
		return
	}
	common.SendSuccessResponse(c, http.StatusOK, del)
}

func (t *TrxRsvController) getApprove(c *gin.Context) {
	claims := t.middleware.GetUser(c)
	if ok := common.AuthorizationGa(claims); ok == false {
		common.SendErrorResponse(c, http.StatusForbidden, "Forbidden")
		return
	}

	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	list, paging, err := t.trxRsvpUC.GetApprovalList(page, size)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var response []interface{}
	for _, v := range list {
		response = append(response, v)
	}
	common.SendPagedResponse(c, response, paging, "success")
}

func (t *TrxRsvController) getAvailable(c *gin.Context) {
	var avble dto.PayloadAvailable
	if err := c.ShouldBindJSON(&avble); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if avble.StartDate == "" || avble.EndDate == "" {
		common.SendErrorResponse(c, http.StatusBadRequest, "required time range")
		return
	}

	response, _ := t.trxRsvpUC.GetAvailableRoom(avble)
	if len(response) == 0 {
		common.SendSuccessResponse(c, http.StatusOK, "Room Full Booked")
		return
	}
	common.SendSuccessResponse(c, http.StatusOK, response)
}
