package service

import (
    "time"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error){
    return sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/sckemployee")
}

type IService interface {
    DeleteEmployee(id string) bool
}

type MockService struct {

}

func  (s *MockService) DeleteEmployee(id string) bool {
    return true;
}

type Service struct {
    DB *sql.DB 
}

type Employee struct{
    Id int   `json:"id"`
    NameTH string `json:"NameTH"`
    NameEN string `json:"NameEN"`
    CitizenID string `json:"CitizenID"`
    Birthday time.Time `json:"CitizenID"`
    Telephone string `json:"telephone"`

}
func  (s *Service) DeleteEmployee(id string) bool {

    statement,_ := s.DB.Prepare("DELETE FROM  employees WHERE id=?")
    defer statement.Close()
    _,err := statement.Exec(id)
    if err != nil {
        
        return false
    }
    return true
}