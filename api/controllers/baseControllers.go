package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres
	"log"
	"net/http"
	"testApi/api/responses"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (app *App) Initialize(DbHost, DbPort, DbUser, DbName, DbPassword string) {
	var err error
	DBURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	app.DB, err = gorm.Open("postgres", DBURI)
	if err != nil {
		fmt.Printf("\n Cannot connect to database %s", DbName)
		log.Fatal("this is the server error:", err)
	} else {
		fmt.Printf("We are connected the database %s", DbName)
	}

	app.Router = mux.NewRouter().StrictSlash(true)
	app.InitializeRoute()
}
func (app *App) InitializeRoute() {

	app.Router.HandleFunc("/", home).Methods("GET")
	app.Router.HandleFunc("/user/{id}", app.GetUser).Methods("GET")
	app.Router.HandleFunc("/allUsers", app.GetAllUsers).Methods("GET")
}
func (app *App) RunServer()  {
	log.Printf("\n Server Starting on port 5000")
	log.Fatal(http.ListenAndServe(":5000",app.Router))
}
func home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome TestApi")
}
