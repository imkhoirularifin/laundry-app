package controller

import (
	"laundry-app/entity"
	"laundry-app/entity/dto"
	"laundry-app/usecase"
	common "laundry-app/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmployeeController interface {
	Find(ctx *gin.Context)
	FindByIdHandler(ctx *gin.Context)
	CreateHandler(ctx *gin.Context)
	Route()
}

type employeeController struct {
	uc usecase.EmployeeUsecase
	rg *gin.RouterGroup
}

func (c *employeeController) Find(ctx *gin.Context) {
	var payload dto.GetAllParams

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response, err := c.uc.Find(payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendPagedResponse(ctx, http.StatusOK, "success", response, gin.H{
		"start": payload.Offset,
		"end":   payload.Offset + payload.Limit,
	})
}

func (c *employeeController) FindByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	response, err := c.uc.FindById(id)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
	}

	common.SendSingleResponse(ctx, http.StatusOK, "success", response)
}

func (c *employeeController) CreateHandler(ctx *gin.Context) {
	var payload entity.Employee

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
	}

	response, err := c.uc.Create(payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
	}

	common.SendSingleResponse(ctx, http.StatusOK, "success", response)
}

func (c *employeeController) Route() {
	c.rg.GET("/employees", c.Find)
	c.rg.GET("/employees/:id", c.FindByIdHandler)
	c.rg.POST("/employees", c.CreateHandler)
}

func NewEmployeeController(uc usecase.EmployeeUsecase, rg *gin.RouterGroup) EmployeeController {
	return &employeeController{uc: uc, rg: rg}
}
