package service

import (
	"database/sql"
	"employee/model"

	_ "github.com/go-sql-driver/mysql"
)

func connectDB() (*sql.DB, error) {
	return sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/testsck")
}

type IService interface {
	AddEmployee(employee model.EmployeeRequest) (model.EmployeeResponse, error)
	DeleteEmployee(id string) bool
}

type MockService struct {
}

func (s *MockService) AddEmployee(employee model.EmployeeRequest) (model.EmployeeResponse, error) {
	return model.EmployeeResponse{
		NameTH:     "นายภาณุมาศ แสนโท",
		NameEN:     "PANUMARS SEANTO",
		CitizenID:  "1-3703-000012-54-8",
		BirthDayBE: "17 เมษายน 2539",
		BirthDayBC: "17 April 1996",
		Telephone:  "0914676299",
	}, nil
}

func (s *MockService) DeleteEmployee(id string) bool {
	return true
}

type Service struct {
	DB *sql.DB
}

func (s *Service) AddEmployee(employee model.EmployeeRequest) (model.EmployeeResponse, error) {
	insertStatement, err := s.DB.Prepare("INSERT INTO employees(nameTH,nameEN,citizenID,birthday,telephone) VALUES( ?, ?, ?, ?, ? )")
	if err != nil {
		return model.EmployeeResponse{}, err
	}
	defer insertStatement.Close()
	_, err = insertStatement.Exec(employee.NameTH, employee.NameEN, employee.CitizenID, employee.Birthday, employee.Telephone)
	if err != nil {
		return model.EmployeeResponse{}, err
	}
	queryStatement, err := s.DB.Prepare("SELECT * FROM employees WHERE citizenID = ?")
	if err != nil {
		return model.EmployeeResponse{}, err
	}
	defer queryStatement.Close()
	var newEmployee model.Employee
	err = queryStatement.QueryRow(1).Scan(&newEmployee)

	var employeeResponse model.EmployeeResponse

	return employeeResponse, nil
}

func (s *Service) DeleteEmployee(id string) bool {

	statement, _ := s.DB.Prepare("DELETE FROM  employees WHERE id=?")
	defer statement.Close()
	_, err := statement.Exec(id)
	if err != nil {

		return false
	}
	return true
}
