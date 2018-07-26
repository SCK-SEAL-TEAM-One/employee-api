package main

import (
    "employee/api"
    "employee/service"
    "log"
    "net/http"
)

func main() {
    ConnectDB,_ := service.ConnectDB()
    apiService := api.API{Service:&service.Service{DB:ConnectDB}}
	http.HandleFunc("/employees", apiService.DeleteEmployeeHandler)

	log.Fatal(http.ListenAndServe(":3000", nil))
}