package controller

import (
	"laundry-app/entity/dto"
	"laundry-app/usecase"
	common "laundry-app/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	LoginHandler(ctx *gin.Context)
	Route()
}

type authController struct {
	authUsecase usecase.AuthUsecase
	rg          *gin.RouterGroup
	jwtService  common.JwtService
}

func (c *authController) LoginHandler(ctx *gin.Context) {
	var payload dto.LoginParams
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response, err := c.authUsecase.Login(payload)
	if err != nil {
		common.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendSingleResponse(ctx, http.StatusOK, "success", response)
}

func (c *authController) Route() {
	// ag = auth group
	ag := c.rg.Group("/auth")
	ag.POST("/login", c.LoginHandler)
}

func NewAuthController(uc usecase.AuthUsecase, rg *gin.RouterGroup, jwtService common.JwtService) AuthController {
	return &authController{
		authUsecase: uc,
		rg:          rg,
		jwtService:  jwtService,
	}
}
