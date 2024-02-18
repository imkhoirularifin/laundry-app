package repository

import (
	"database/sql"
	"laundry-app/entity"
	"laundry-app/entity/dto"
	"time"
)

type EmployeeRepository interface {
	Find(payload dto.GetAllParams) ([]entity.Employee, error)
	FindById(id string) (entity.Employee, error)
	FindByUsername(username string) (entity.Employee, error)
	Create(employee entity.Employee) (entity.Employee, error)
	// Update(employee entity.Employee) (entity.Employee, error)
	// Delete(id string) error
}

type employeeRepository struct {
	db *sql.DB
}

func (e *employeeRepository) FindById(id string) (entity.Employee, error) {
	var employee entity.Employee
	err := e.db.QueryRow("SELECT * FROM employees WHERE id = $1", id).Scan(
		&employee.Id,
		&employee.Name,
		&employee.Email,
		&employee.Username,
		&employee.Password,
		&employee.CreatedAt,
		&employee.UpdatedAt,
	)

	if err != nil {
		return entity.Employee{}, err
	}

	return employee, nil
}

func (e *employeeRepository) Find(payload dto.GetAllParams) ([]entity.Employee, error) {
	var employees []entity.Employee

	rows, err := e.db.Query("SELECT * FROM employees ORDER BY name LIMIT $1 OFFSET $2", payload.Limit, payload.Offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var employee entity.Employee
		err := rows.Scan(
			&employee.Id,
			&employee.Name,
			&employee.Email,
			&employee.Username,
			&employee.Password,
			&employee.CreatedAt,
			&employee.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		employees = append(employees, employee)
	}

	return employees, nil
}

func (e *employeeRepository) Create(employee entity.Employee) (entity.Employee, error) {
	err := e.db.QueryRow("INSERT INTO employees (name, email, username, password, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, name, email, username, password, created_at, updated_at", employee.Name, employee.Email, employee.Username, employee.Password, time.Now()).Scan(
		&employee.Id,
		&employee.Name,
		&employee.Email,
		&employee.Username,
		&employee.Password,
		&employee.CreatedAt,
		&employee.UpdatedAt,
	)

	if err != nil {
		return entity.Employee{}, err
	}

	return employee, nil
}

func (e *employeeRepository) FindByUsername(username string) (entity.Employee, error) {
	var employee entity.Employee

	err := e.db.QueryRow("SELECT * FROM employees WHERE username = $1", username).Scan(
		&employee.Id,
		&employee.Name,
		&employee.Email,
		&employee.Username,
		&employee.Password,
		&employee.CreatedAt,
		&employee.UpdatedAt,
	)

	if err != nil {
		return entity.Employee{}, err
	}

	return employee, nil
}

func NewEmployeeRepository(db *sql.DB) EmployeeRepository {
	return &employeeRepository{db: db}
}
