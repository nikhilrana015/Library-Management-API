package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	config "github.com/nikhilrana/Go-LibraryMgmt/pkg/config"
	routes "github.com/nikhilrana/Go-LibraryMgmt/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookInLibraryRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:"+config.EnvConfigs.LocalServerPort, r))
}
