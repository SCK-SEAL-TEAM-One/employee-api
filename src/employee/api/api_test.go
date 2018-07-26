package api_test

import (
	. "employee/api"
	"employee/model"
	"employee/service"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func Test_DeleteEmployeeHandler_Input_ID_1_Should_Be_Status_OK(t *testing.T) {
	request := httptest.NewRequest("DELETE", "/employees?id=1", nil)
	responseRecorder := httptest.NewRecorder()

	expected := 200

	api := API{
		Service: &service.MockService{},
	}
	api.DeleteEmployeeHandler(responseRecorder, request)
	actual := responseRecorder.Result().StatusCode
	if expected != actual {
		t.Errorf("expected status %d but it got %d", expected, actual)
	}
}
func Test_PostAddEmpolyeeHandler_Input_JsonBody_Should_Be_EmployeeResponse(t *testing.T) {
	birthDay, _ := time.Parse(time.RFC3339, "1996-04-17T00:00:00+00:00")
	requestEmployee := model.EmployeeRequest{
		NameTH:    "นายภาณุมาศ แสนโท",
		NameEN:    "PANUMARS SEANTO",
		CitizenID: "1-3703-000012-54-8",
		Birthday:  birthDay,
		Telephone: "0914676299",
	}
	requestEmployeeJson, _ := json.Marshal(requestEmployee)
	request := httptest.NewRequest("POST", "/employees", strings.NewReader(string(requestEmployeeJson)))
	responseRecorder := httptest.NewRecorder()

	expected := model.EmployeeResponse{
		NameTH:     "นายภาณุมาศ แสนโท",
		NameEN:     "PANUMARS SEANTO",
		CitizenID:  "1-3703-000012-54-8",
		BirthDayBE: "17 เมษายน 2539",
		BirthDayBC: "17 April 1996",
		Telephone:  "0914676299",
	}

	api := API{
		Service: &service.MockService{},
	}
	api.PostAddEmpolyeeHandler(responseRecorder, request)
	response := responseRecorder.Result()
	body, _ := ioutil.ReadAll(response.Body)
	var actual model.EmployeeResponse
	json.Unmarshal(body, &actual)

	if expected != actual {
		t.Errorf("expected status %v but it got %v", expected, actual)
	}
}
