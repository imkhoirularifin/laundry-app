package usecase

import (
	"laundry-app/entity"
	"laundry-app/entity/dto"
	"laundry-app/repository"
	"laundry-app/utils/common"
)

type EmployeeUsecase interface {
	Find(payload dto.GetAllParams) ([]entity.Employee, error)
	FindById(id string) (entity.Employee, error)
	FindByUsernameAndPassword(payload dto.LoginParams) (entity.Employee, error)
	Create(employee entity.Employee) (entity.Employee, error)
}

type employeeUsecase struct {
	employeeRepository repository.EmployeeRepository
}

func (e *employeeUsecase) Find(payload dto.GetAllParams) ([]entity.Employee, error) {
	employees, err := e.employeeRepository.Find(payload)
	if err != nil {
		return nil, err
	}

	for i := range employees {
		employees[i].Password = ""
	}

	return employees, nil
}

func (e *employeeUsecase) FindById(id string) (entity.Employee, error) {
	employee, err := e.employeeRepository.FindById(id)
	if err != nil {
		return entity.Employee{}, err
	}

	employee.Password = ""

	return employee, nil
}

func (e *employeeUsecase) Create(employee entity.Employee) (entity.Employee, error) {
	password, err := common.GeneratePassword(employee.Password)
	if err != nil {
		return entity.Employee{}, err
	}

	employee.Password = password

	employee, err = e.employeeRepository.Create(employee)
	if err != nil {
		return entity.Employee{}, err
	}

	employee.Password = ""

	return employee, nil
}

func (e *employeeUsecase) FindByUsernameAndPassword(payload dto.LoginParams) (entity.Employee, error) {
	employee, err := e.employeeRepository.FindByUsername(payload.Username)
	if err != nil {
		return entity.Employee{}, err
	}

	// check password
	err = common.ComparePassword(employee.Password, payload.Password)
	if err != nil {
		return entity.Employee{}, err
	}

	employee.Password = ""

	return employee, nil
}

func NewEmployeeUsecase(employeeRepository repository.EmployeeRepository) EmployeeUsecase {
	return &employeeUsecase{employeeRepository: employeeRepository}
}
