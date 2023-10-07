package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"

	"git.tashilcar.com/core/tashilcar/app/di"
	errs "git.tashilcar.com/core/tashilcar/core/yerrors"
	"git.tashilcar.com/core/tashilcar/presentation/rest"
)

type Router struct {
	Router *mux.Router
}

func StartApplication() {
	fmt.Println("start application")
	// migrate tables
	migrate()
	fmt.Println("migration successful!")

	rest.Start()
	fmt.Println("start routes")
}

func StopApplication() {
	fmt.Println("* SERVER STOPPED *")
	db, err := di.DB().DB()
	if errs.IsNotNil(err) {
		panic("* FAILED TO GET GENERIC DB OBJECT *")
	}

	err = db.Close()
	if errs.IsNotNil(err) {
		panic("* FAIL TO CLOSE DB CONNECTION *")
	}
	fmt.Println("* DB CONNECTION STOPPED *")
}

func migrate() {
	err := di.DBDS().CreateTables()
	if errs.IsNotNil(err) {
		panic(err)
	}
}

func (server *Server) run() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	server.Run(":8080")

}

type Server struct {
	Router *mux.Router
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func (s *Router) initializeRoutes() {

	fmt.Println("initializeRoutes")
}
