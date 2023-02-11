package implementation

import (
	"github.com/joyzem/proxy-project/services/base"
	"github.com/joyzem/proxy-project/services/employee/backend/repo"
	"github.com/joyzem/proxy-project/services/employee/backend/service"
	"github.com/joyzem/proxy-project/services/employee/domain"
)

type employeeService struct {
	employeeRepo repo.EmployeeRepo
}

func NewEmployeeService(repo repo.EmployeeRepo) service.EmployeeService {
	return &employeeService{
		employeeRepo: repo,
	}
}

func (s *employeeService) CreateEmployee(
	firstName string,
	lastName string,
	middleName string,
	post string,
	passportSeries string,
	passportNumber string,
	passportIssuedBy string,
	passportDateOfIssue string,
) (*domain.Employee, error) {
	employee := domain.Employee{
		FirstName:           firstName,
		LastName:            lastName,
		MiddleName:          middleName,
		Post:                post,
		PassportSeries:      passportSeries,
		PassportNumber:      passportNumber,
		PassportIssuedBy:    passportIssuedBy,
		PassportDateOfIssue: passportDateOfIssue,
	}
	result, err := s.employeeRepo.CreateEmployee(employee)
	if err != nil {
		base.LogError(err)
	}
	return result, nil
}

func (s *employeeService) GetEmployees() ([]domain.Employee, error) {
	employees, err := s.employeeRepo.GetEmployees()
	if err != nil {
		base.LogError(err)
	}
	return employees, err
}

func (s *employeeService) EmployeeById(id int) (*domain.Employee, error) {
	employee, err := s.employeeRepo.EmployeeById(id)
	if err != nil {
		base.LogError(err)
	}
	return employee, err
}

func (s *employeeService) DeleteEmployee(id int) error {
	if err := s.employeeRepo.DeleteEmployee(id); err != nil {
		base.LogError(err)
		return err
	}
	return nil
}

func (s *employeeService) UpdateEmployee(empl domain.Employee) (*domain.Employee, error) {
	employee, err := s.employeeRepo.UpdateEmployee(empl)
	if err != nil {
		base.LogError(err)
	}
	return employee, err
}
