package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/k-ueki/AGPlus/server/adaptor/api/input"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/k-ueki/AGPlus/server/adaptor/repository"
	"github.com/k-ueki/AGPlus/server/application/service"
)

type (
	FacultyGetController struct {
		service.FacultyGetService
	}
)

func NewFacultyGetController(db *gorm.DB) *FacultyGetController {
	return &FacultyGetController{
		FacultyGetService: service.FacultyGetService{
			FacultyGetRepository: repository.FacultyGetRepository{
				DB: db,
			},
		},
	}
}

func (c *FacultyGetController) ListFaculty(ctx *gin.Context) {
	q := input.ListFacultyByCampusID{}
	if err := ctx.BindQuery(&q); err != nil {
		ctx.JSON(http.StatusBadRequest, errors.New("failed to bind query"))
		return
	}

	if q.CampusID != 0 {
		faculties, err := c.FacultyGetService.ListFacultyByCampusID(q.CampusID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, "failed to list faculties")
			return
		}
		ctx.JSON(http.StatusOK, faculties)
		return
	}

	faculties, err := c.FacultyGetService.ListFaculty()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "failed to list faculties")
		return
	}
	ctx.JSON(http.StatusOK, faculties)
	return
}

func (c *FacultyGetController) ShowFaculty(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errors.New("failed to bind parameters"))
		return
	}

	faculty, err := c.FacultyGetService.ShowFacultyByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "failed to show faculty")
		return
	}
	ctx.JSON(http.StatusOK, faculty)
	return
}

func (c *FacultyGetController) ListDepartment(ctx *gin.Context) {
	q := input.ListDepartmentByCampusID{}
	if err := ctx.BindQuery(&q); err != nil {
		ctx.JSON(http.StatusBadRequest, errors.New("failed to bind query"))
		return
	}

	if q.CampusID != 0 {
		departments, err := c.FacultyGetService.ListDepartmentsByCampusID(q.CampusID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, "failed to list departments")
			return
		}
		ctx.JSON(http.StatusOK, departments)
		return
	}

	departments, err := c.FacultyGetService.ListDepartment()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "failed to list departments")
		return
	}
	ctx.JSON(http.StatusOK, departments)
	return
}

func (c *FacultyGetController) ShowDepartment(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errors.New("failed to bind parameters"))
		return
	}

	department, err := c.FacultyGetService.ShowDepartmentByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "failed to show department")
		return
	}
	ctx.JSON(http.StatusOK, department)
	return
}
