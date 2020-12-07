package main

import (
	"fmt"

	"github.com/Kally95/Go_Web_App/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "secret"
	dbname   = "go_web_dev"
)

type User struct {
	gorm.Model
	Name   string
	Email  string `gorm:"not null; unique_index"`
	Orders []Order
}

type Order struct {
	gorm.Model
	UserID      uint
	Amount      int
	Description string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	us.DestructiveReset()
	user := models.User{
		Name:  "James Bond",
		Email: "007@gmail.com",
	}
	if err := us.Create(&user); err != nil {
		panic(err)
	}
	if err := us.Delete(user.ID); err != nil {
		panic(err)
	}
	userById, err := us.ByID(user.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println(userById)
}
