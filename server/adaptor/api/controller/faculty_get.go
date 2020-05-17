package api

import (
	"net/http"

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
	faculties, err := c.FacultyGetService.ListFaculty()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "failed to list faculties")
		return
	}
	ctx.JSON(http.StatusOK, faculties)
	return
}

func (c *FacultyGetController) ListDepartment(ctx *gin.Context) {
	departments, err := c.FacultyGetService.ListDepartment()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "failed to list faculties")
		return
	}
	ctx.JSON(http.StatusOK, departments)
	return
}
