package controller

import (
	"booking-room/model"
	"booking-room/shared/common"
	"booking-room/usecase"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoomController struct {
	roomUC usecase.RoomUseCase
	rg     *gin.RouterGroup
}

func (r *RoomController) Route() {
	r.rg.POST("/create", r.createHandler)
	r.rg.PUT("/:id", r.updateHandler)
	r.rg.GET("/", r.listHandler)
	r.rg.GET("/:id", r.getHandler)
}

// sudah benar
func (r *RoomController) getHandler(c *gin.Context) {
	id := c.Param("id")
	room, err := r.roomUC.FindRoomById(id)
	if err != nil {
		common.SendErrorResponse(c, http.StatusNotFound, fmt.Sprintf("Room with ID %s not found", id))
		return
	}
	common.SendSingleResponse(c, room, "Ok")
}

func (r *RoomController) createHandler(c *gin.Context) {
	var payload model.Room
	if err := c.ShouldBindJSON(&payload); err != nil {
		log.Println("Error binding JSON:", err.Error())
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	room, err := r.roomUC.RegisterNewRoom(payload)
	if err != nil {
		log.Println("Error registering new room:", err.Error())
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendCreateResponse(c, room, "Created")
}

func (r *RoomController) updateHandler(c *gin.Context) {
	var payload model.Room
	if err := c.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// room, err := r.roomUC.UpdateRoom(payload)
	// if err != nil {
	//  common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
	//  return
	// }
	// common.SendCreateResponse(c, room, "Updated")

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    payload,
	})
}

func (r *RoomController) listHandler(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	sizeStr := c.DefaultQuery("size", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid page parameter: %s must be an integer", pageStr))
		return
	}

	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid size parameter: %s must be an integer", sizeStr))
		return
	}

	rooms, paging, err := r.roomUC.FindAllRoom(page, size)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response := gin.H{
		"rooms":  rooms,
		"paging": paging,
	}

	common.SendListResponse(c, response, "Ok")
}

func NewRoomController(roomUC usecase.RoomUseCase, rg *gin.RouterGroup) *RoomController {
	return &RoomController{
		roomUC: roomUC,
		rg:     rg,
	}
}
