package main

import (
	"fmt"
	"net/http"

	"github.com/Kally95/Go_Web_App/controllers"
	"github.com/Kally95/Go_Web_App/middleware"
	"github.com/Kally95/Go_Web_App/models"
	"github.com/gorilla/mux"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "secret"
	dbname   = "go_web_dev"
)

func main() {
	// Create a DB connection string and then use it to
	// create the model services.
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	services, err := models.NewServices(psqlInfo)
	if err != nil {
		panic(err)
	}
	// TODO: Fix this
	defer services.Close()
	// services.DestructiveReset()
	services.AutoMigrate()

	r := mux.NewRouter()

	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers(services.User)
	galleriesC := controllers.NewGalleries(services.Gallery, r)
	requireUserMw := middleware.RequireUser{
		UserService: services.User,
	}

	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	r.Handle("/login", usersC.LoginView).Methods("GET")
	r.HandleFunc("/login", usersC.Login).Methods("POST")
	r.HandleFunc("/cookietest", usersC.CookieTest).Methods("GET")

	// Gallery routes
	galleryNew := requireUserMw.Apply(galleriesC.New)
	r.Handle("/galleries/new", requireUserMw.Apply(galleryNew)).Methods("GET")
	r.HandleFunc("/galleries", requireUserMw.ApplyFn(galleriesC.Create)).Methods("POST")
	r.HandleFunc("/galleries/{id:[0-9]+}", galleriesC.Show).Methods("GET").Name(controllers.ShowGallery)

	http.ListenAndServe(":3000", r)
}
