package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
	Name  string
	Email string `gorm:"not null; unique_index"`
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if err := db.DB().Ping(); err != nil {
		panic(err)
	}
	db.LogMode(true)
	//db.DropTableIfExists(&User{})
	//db.AutoMigrate(&User{})
	name, email := getInfo()
	u := User{
		Name:  name,
		Email: email,
	}

	if err = db.Create(&u).Error; err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", u)
}

func getInfo() (name, email string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What;s your name?")
	name, _ = reader.ReadString('\n')
	fmt.Println("What's your email address?")
	email, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)
	return name, email
}
