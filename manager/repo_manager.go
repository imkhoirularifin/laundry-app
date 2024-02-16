package manager

import "laundry-app/repository"

type RepoManager interface {
	EmployeeRepo() repository.EmployeeRepository
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) EmployeeRepo() repository.EmployeeRepository {
	return repository.NewEmployeeRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}