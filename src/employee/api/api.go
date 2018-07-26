package api

import (
	"employee/service"

	"employee/model"
	"encoding/json"
	"net/http"
)

type API struct {
	Service service.IService
}

func (api *API) DeleteEmployeeHandler(writer http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")
	status := api.Service.DeleteEmployee(id)
	if status {
		writer.WriteHeader(http.StatusOK)
		return
	}
	http.Error(writer, "Delete Error", http.StatusInternalServerError)
}

func (api *API) PostAddEmpolyeeHandler(write http.ResponseWriter, request *http.Request) {
	employee := model.EmployeeRequest{}
	err := json.NewDecoder(request.Body).Decode(&employee)
	if err != nil {
		http.Error(write, err.Error(), http.StatusInternalServerError)
		return
	}
	newEmployeeResponse, err := api.Service.AddEmployee(employee)
	if err != nil {
		http.Error(write, err.Error(), http.StatusInternalServerError)
		return
	}
	newEmployeeResponseJson, _ := json.Marshal(newEmployeeResponse)
	write.Write(newEmployeeResponseJson)
}
