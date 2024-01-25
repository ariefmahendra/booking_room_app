package controller

import (
	"booking-room/model"
	"booking-room/shared/common"
	"booking-room/usecase"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type FacilitiesController struct {
	facilitiesUsecase usecase.FacilitiesUsecase
	fg                *gin.RouterGroup
}

// FindAllFacilities Get all facilities
func (f *FacilitiesController) FindAllFacilities(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	//if query param empty
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 5
	}
	facilities, paging, err := f.facilitiesUsecase.List(page, size)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendSuccessPagedResponse(c, http.StatusOK, facilities, paging)
	return

}

// FindFacilityById Get facility by id
func (f *FacilitiesController) FindFacilityById(c *gin.Context) {
	id := c.Param("id")
	facility, err := f.facilitiesUsecase.Get(id)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	common.SendSuccessResponse(c, http.StatusOK, facility)
}

// FindFacilityByName Get facility by name
func (f *FacilitiesController) FindFacilityByName(c *gin.Context) {
	name := strings.ToUpper(c.Param("codeName"))
	facility, err := f.facilitiesUsecase.GetByName(name)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	common.SendSuccessResponse(c, http.StatusOK, facility)
}

// FindFacilityByStatus Get facility by status
func (f *FacilitiesController) FindFacilityByStatus(c *gin.Context) {
	status := strings.ToUpper(c.Param("status"))
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 5
	}
	facility, paging, err := f.facilitiesUsecase.GetByStatus(status, page, size)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	common.SendSuccessPagedResponse(c, http.StatusOK, facility, paging)
}

// FindFacilityByType Get facility by Facilities Type
func (f *FacilitiesController) FindFacilityByType(c *gin.Context) {
	ftype := strings.ToLower(c.Param("FacilitiesType"))
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 5
	}
	facility, paging, err := f.facilitiesUsecase.GetByType(ftype, page, size)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	common.SendSuccessPagedResponse(c, http.StatusOK, facility, paging)
}

// CreateFacility Create new facility
func (f *FacilitiesController) CreateFacility(c *gin.Context) {
	var payload model.Facilities
	if err := c.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	CreateFacility, err := f.facilitiesUsecase.Create(payload)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	common.SendSuccessResponse(c, http.StatusOK, CreateFacility)
}

// UpdateFacility Update facility
func (f *FacilitiesController) UpdateFacility(c *gin.Context) {
	id := c.Param("id")
	var payload model.Facilities
	if err := c.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	facility, err := f.facilitiesUsecase.Update(payload, id)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	common.SendSuccessResponse(c, http.StatusOK, facility)
}

// DeleteFacility Delete facility by id
func (f *FacilitiesController) DeleteFacility(c *gin.Context) {
	id := c.Param("id")
	err := f.facilitiesUsecase.Delete(id)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	common.SendSuccessResponse(c, http.StatusOK, nil)
}

// DeleteFacilityByName Delete facility by name
func (f *FacilitiesController) DeleteFacilityByName(c *gin.Context) {
	name := strings.ToUpper(c.Param("codeName"))
	err := f.facilitiesUsecase.DeleteByName(name)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	common.SendSuccessResponse(c, http.StatusOK, nil)
}

// FindAllDeletedFacilities Get deleted facilities
func (f *FacilitiesController) FindAllDeletedFacilities(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 5
	}
	facility, paging, err := f.facilitiesUsecase.GetDeleted(page, size)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	common.SendSuccessPagedResponse(c, http.StatusOK, facility, paging)
}

// NewFacilitiesController constructor for facilities controller
func NewFacilitiesController(facilitiesUsecase usecase.FacilitiesUsecase, fg *gin.RouterGroup) *FacilitiesController {
	return &FacilitiesController{
		facilitiesUsecase: facilitiesUsecase,
		fg:                fg,
	}
}

// Route setup route for facilities
func (f *FacilitiesController) Route() {
	f.fg.GET("/", f.FindAllFacilities)
	f.fg.GET("/id/:id", f.FindFacilityById)
	f.fg.GET("/name/:codeName", f.FindFacilityByName)
	f.fg.GET("/status/:status", f.FindFacilityByStatus)
	f.fg.GET("/type/:FacilitiesType", f.FindFacilityByType)
	f.fg.POST("/", f.CreateFacility)
	f.fg.PUT("/:id", f.UpdateFacility)
	f.fg.DELETE("/:id", f.DeleteFacility)
	f.fg.DELETE("/name/:codeName", f.DeleteFacilityByName)
	f.fg.GET("/deleted", f.FindAllDeletedFacilities)
}
