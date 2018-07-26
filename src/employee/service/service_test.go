package service_test

import (
    "testing"
    ."employee/service"
)

func Test_DeleteEmployee_Input_ID_1_Should_Be_True(t *testing.T)  {
    id := "1"
    expected := true

    service := Service{}

    actual := service.DeleteEmployee(id)
    if expected != actual{
        t.Errorf("expected %v but got %v",expected,actual)
    }
}
