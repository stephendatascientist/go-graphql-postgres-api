package repositories

import (
	"github.com/stephendatascientist/go-graphql-postgres-api/entities"
	"github.com/stephendatascientist/go-graphql-postgres-api/graph/model"
	"gorm.io/gorm"
)

type EmployeeRepository struct {
	DB *gorm.DB
}

func (r *EmployeeRepository) CreateEmployee(employeeInput *model.NewEmployee) (*entities.Employee, error) {
	employee := &entities.Employee{
		Name:   employeeInput.Name,
		Email:  employeeInput.Email,
		Mobile: employeeInput.Mobile,
	}
	err := r.DB.Create(&employee).Error

	return employee, err
}

func (r *EmployeeRepository) GetEmployees() ([]*model.Employee, error) {
	employees := []*model.Employee{}
	err := r.DB.Find(&employees).Error
	return employees, err

}

func (r *EmployeeRepository) GetEmployee(id int) (*entities.Employee, error) {
	employee := &entities.Employee{}
	err := r.DB.Where("id = ?", id).First(employee).Error
	return employee, err
}

func (r *EmployeeRepository) UpdateEmployee(employeeInput *model.NewEmployee, id int) error {
	employee := entities.Employee{
		ID:     id,
		Name:   employeeInput.Name,
		Email:  employeeInput.Email,
		Mobile: employeeInput.Mobile,
	}
	err := r.DB.Model(&employee).Where("id = ?", id).Updates(employee).Error
	return err
}

func (r *EmployeeRepository) DeleteEmployee(id int) error {
	employee := &entities.Employee{}
	err := r.DB.Delete(employee, id).Error
	return err
}
