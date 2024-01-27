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
	common.SendSuccessResponse(c, http.StatusOK, room)
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

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    payload,
	})
}

func (r *RoomController) listHandler(c *gin.Context) {
    page, _ := strconv.Atoi(c.Query("page"))
    size, _ := strconv.Atoi(c.Query("size"))

    rooms, paging, err := r.roomUC.FindAllRoom(page, size)
    if err != nil {
        common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }

    common.SendSuccessPagedResponse(c, http.StatusOK, rooms, paging)
}


func NewRoomController(roomUC usecase.RoomUseCase, rg *gin.RouterGroup) *RoomController {
	return &RoomController{
		roomUC: roomUC,
		rg:     rg,
	}
}
