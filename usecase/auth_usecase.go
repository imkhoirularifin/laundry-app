package usecase

import (
	"laundry-app/entity/dto"
	"laundry-app/utils/common"
)

type AuthUsecase interface {
	Login(payload dto.LoginParams) (dto.AuthResponse, error)
}

type authUsecase struct {
	employeeUsecase EmployeeUsecase
	jwtService      common.JwtService
}

func (a *authUsecase) Login(payload dto.LoginParams) (dto.AuthResponse, error) {
	employee, err := a.employeeUsecase.FindByUsernameAndPassword(payload)
	if err != nil {
		return dto.AuthResponse{}, err
	}

	token, err := a.jwtService.GenerateToken(employee)
	if err != nil {
		return dto.AuthResponse{}, err
	}

	return token, nil
}

func NewAuthUsecase(employeeUsecase EmployeeUsecase, jwtToken common.JwtService) AuthUsecase {
	return &authUsecase{
		employeeUsecase: employeeUsecase,
		jwtService:      jwtToken,
	}
}
