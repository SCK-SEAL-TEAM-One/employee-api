package api

import (
    "net/http"
    "employee/service"
)

type API struct {
    Service service.IService
}

func (api *API) DeleteEmployeeHandler(writer http.ResponseWriter, request *http.Request)  {
    id := request.URL.Query().Get("id")
    status := api.Service.DeleteEmployee(id)
    if status{
        writer.WriteHeader(http.StatusOK)
        return
    }
    http.Error(writer,"Delete Error",http.StatusInternalServerError)
}
