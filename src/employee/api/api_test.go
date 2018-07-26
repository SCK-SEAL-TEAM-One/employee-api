package api_test

import (
    ."employee/api"
    "testing"
    "net/http/httptest"
    "employee/service"
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