package manager

import "laundry-app/usecase"

type UsecaseManager interface {
	EmployeeUsecase() usecase.EmployeeUsecase
}

type usecaseManager struct {
	repo RepoManager
}

func (u *usecaseManager) EmployeeUsecase() usecase.EmployeeUsecase {
	return usecase.NewEmployeeUsecase(u.repo.EmployeeRepo())
}

/*
	Overwrite usecase manager struct with repo manager struct
*/
func NewUsecaseManager(repo RepoManager) UsecaseManager {
	return &usecaseManager{repo: repo}
}
