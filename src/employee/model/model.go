package model

import "time"

type Employee struct {
	Id        int       `json:"id"`
	NameTH    string    `json:"NameTH"`
	NameEN    string    `json:"NameEN"`
	CitizenID string    `json:"CitizenID"`
	Birthday  time.Time `json:"Birthday"`
	Telephone string    `json:"telephone"`
}

type EmployeeRequest Employee

type EmployeeResponse struct {
	NameTH     string `json:"NameTH"`
	NameEN     string `json:"NameEN"`
	CitizenID  string `json:"CitizenID"`
	BirthDayBE string `json:"BirthDayBE"`
	BirthDayBC string `json:"BirthDayBC"`
	Telephone  string `json:"Telephone"`
}
