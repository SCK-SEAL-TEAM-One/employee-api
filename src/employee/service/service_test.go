package service_test

import (
	"employee-api/atdd/src/api"
	. "employee/model"
	. "employee/service"
	"testing"
)

func Test_DeleteEmployee_Input_ID_1_Should_Be_True(t *testing.T) {
	id := "1"
	expected := true

	service := Service{}

	actual := service.DeleteEmployee(id)
	if expected != actual {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}

func Test_AddEmployee_Input_Employee_Should_Be_Employee(t *testing.T) {
	employee := api.Employee{
		NameTH:     "นายภาณุมาศ แสนโท",
		NameEN:     "PANUMARS SEANTO",
		CitizenID:  "1-3703-000012-54-8",
		BirthDayBE: "17 เมษายน 2539",
		BirthDayBC: "17 April 1996",
		Telephone:  "0914676299",
	}
	expected := api.Employee{
		NameTH:     "นายภาณุมาศ แสนโท",
		NameEN:     "PANUMARS SEANTO",
		CitizenID:  "1-3703-000012-54-8",
		BirthDayBE: "17 เมษายน 2539",
		BirthDayBC: "17 April 1996",
		Telephone:  "09-1467-6299",
	}

	actual := AddEmployee(employee)

	if expected != actual {
		t.Errorf("Should be %v but got %v", expected, actual)
	}
}
