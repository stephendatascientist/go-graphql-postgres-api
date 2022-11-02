package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/stephendatascientist/go-graphql-postgres-api/graph/generated"
	"github.com/stephendatascientist/go-graphql-postgres-api/graph/model"
)

// CreateEmployee is the resolver for the createEmployee field.
func (r *mutationResolver) CreateEmployee(ctx context.Context, input model.NewEmployee) (*model.Employee, error) {
	// panic(fmt.Errorf("not implemented: CreateEmployee - createEmployee"))
	employee, err := r.EmployeeRepository.CreateEmployee(&input)
	employeeCreated := &model.Employee{
		ID:     employee.ID,
		Name:   employee.Name,
		Email:  employee.Email,
		Mobile: employee.Mobile,
	}
	if err != nil {
		return nil, err
	}
	return employeeCreated, nil
}

// DeleteEmployee is the resolver for the DeleteEmployee field.
func (r *mutationResolver) DeleteEmployee(ctx context.Context, id int) (string, error) {
	// panic(fmt.Errorf("not implemented: DeleteEmployee - DeleteEmployee"))
	err := r.EmployeeRepository.DeleteEmployee(id)
	if err != nil {
		return "", err
	}
	successMessage := "successfully deleted"
	return successMessage, nil
}

// UpdateEmployee is the resolver for the UpdateEmployee field.
func (r *mutationResolver) UpdateEmployee(ctx context.Context, id int, input model.NewEmployee) (string, error) {
	// panic(fmt.Errorf("not implemented: UpdateEmployee - UpdateEmployee"))
	err := r.EmployeeRepository.UpdateEmployee(&input, id)
	if err != nil {
		return "nil", err
	}
	successMessage := "successfully updated"

	return successMessage, nil
}

// GetEmployees is the resolver for the GetEmployees field.
func (r *queryResolver) GetEmployees(ctx context.Context) ([]*model.Employee, error) {
	// panic(fmt.Errorf("not implemented: GetEmployees - GetEmployees"))
	employees, err := r.EmployeeRepository.GetEmployees()
	if err != nil {
		return nil, err
	}
	return employees, nil
}

// GetEmployee is the resolver for the GetEmployee field.
func (r *queryResolver) GetEmployee(ctx context.Context, id int) (*model.Employee, error) {
	// panic(fmt.Errorf("not implemented: GetEmployee - GetEmployee"))
	employee, err := r.EmployeeRepository.GetEmployee(id)
	selectedEmployee := &model.Employee{
		ID:     employee.ID,
		Name:   employee.Name,
		Email:  employee.Email,
		Mobile: employee.Mobile,
	}
	if err != nil {
		return nil, err
	}
	return selectedEmployee, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Employees(ctx context.Context) ([]*model.Employee, error) {
	// panic(fmt.Errorf("not implemented: Employees - employees"))
	employees, err := r.EmployeeRepository.GetEmployees()
	if err != nil {
		return nil, err
	}
	return employees, nil
}
