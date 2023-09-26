package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

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

func startServer() {
	mux := http.NewServeMux()

	now := time.Now()

	mux.HandleFunc("/today", func(rw http.ResponseWriter, _ *http.Request) {

		rw.Write([]byte(now.Format(time.ANSIC)))
	})

	a := &Router{}
	mux.HandleFunc("/today2", a.today2)
	//a.initializeRoutes()

	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)

	// server.Run(":8080")

}

func example(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, "success"))
	JSON(w, http.StatusCreated, "exampleeeeeeeeeee")
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func (s *Router) initializeRoutes() {

	fmt.Println("initializeRoutes")
	// Home Route
	s.Router.HandleFunc("/test", s.Home).Methods("GET")
}

func (s *Router) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home")

	JSON(w, http.StatusOK, "Welcome To This Awesome API")
}

func (s *Router) today2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("today2today2today2")

	JSON(w, http.StatusOK, "Welcome To This today2 API")
}
