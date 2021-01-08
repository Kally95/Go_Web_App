package main

import (
	"fmt"
	"net/http"

	"github.com/Kally95/Go_Web_App/controllers"
	"github.com/Kally95/Go_Web_App/middleware"
	"github.com/Kally95/Go_Web_App/models"
	"github.com/Kally95/Go_Web_App/rand"
	"github.com/gorilla/csrf"
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
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	services, err := models.NewServices(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer services.Close()
	services.AutoMigrate()
	//services.DestructiveReset()

	r := mux.NewRouter()
	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers(services.User)
	galleriesC := controllers.NewGalleries(services.Gallery, services.Image, r)

	// TODO: UPdate this to be a config variable
	isProd := false
	b, err := rand.Bytes(32)
	if err != nil {
		panic(err)
	}
	csrfMw := csrf.Protect(b, csrf.Secure(isProd))

	userMw := middleware.User{
		UserService: services.User,
	}
	requireUserMw := middleware.RequireUser{}

	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	r.Handle("/login", usersC.LoginView).Methods("GET")
	r.HandleFunc("/login", usersC.Login).Methods("POST")
	r.HandleFunc("/cookietest", usersC.CookieTest).Methods("GET")
	// Gallery routes
	r.Handle("/galleries",
		requireUserMw.ApplyFn(galleriesC.Index)).
		Methods("GET").
		Name(controllers.IndexGalleries)
	r.Handle("/galleries/new",
		requireUserMw.Apply(galleriesC.New)).
		Methods("GET")
	r.Handle("/galleries",
		requireUserMw.ApplyFn(galleriesC.Create)).
		Methods("POST")
	r.HandleFunc("/galleries/{id:[0-9]+}",
		galleriesC.Show).
		Methods("GET").
		Name(controllers.ShowGallery)
	r.HandleFunc("/galleries/{id:[0-9]+}/edit",
		requireUserMw.ApplyFn(galleriesC.Edit)).
		Methods("GET").
		Name(controllers.EditGallery)
	r.HandleFunc("/galleries/{id:[0-9]+}/update",
		requireUserMw.ApplyFn(galleriesC.Update)).
		Methods("POST")
	r.HandleFunc("/galleries/{id:[0-9]+}/delete",
		requireUserMw.ApplyFn(galleriesC.Delete)).
		Methods("POST")
	r.HandleFunc("/galleries/{id:[0-9]+}/images",
		requireUserMw.ApplyFn(galleriesC.ImageUpload)).
		Methods("POST")
	r.HandleFunc("/galleries/{id:[0-9]+}/images/{filename}/delete",
		requireUserMw.ApplyFn(galleriesC.ImageDelete)).
		Methods("POST")

	// Assets
	assetHandler := http.FileServer(http.Dir("./assets"))
	assetHandler = http.StripPrefix("/assets", assetHandler)
	r.PathPrefix("/assets/").Handler(assetHandler)

	// Image routes
	imageHandler := http.FileServer(http.Dir("./images/"))
	r.PathPrefix("/images/").Handler(http.StripPrefix("/images/", imageHandler))

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", csrfMw(userMw.Apply(r)))
}
