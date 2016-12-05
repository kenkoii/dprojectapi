package main

import (
	"github.com/kenkoii/dprojectapi/api/handlers"
	"github.com/kenkoii/dprojectapi/api/routers"
	"github.com/codegangsta/negroni"
	"net/http"
	"github.com/rs/cors"
	//"github.com/kenkoii/dprojectapi/api/common"
	"github.com/kenkoii/dprojectapi/api/common"
)

func init() {
	//router := mux.NewRouter()
	/*http.Handle("/", router)*/
	//router.HandleFunc("/", handlers.Handler)*/
	//Categories
	// Get the mux router object
	common.StartUp()
	router := routers.InitRoutes()
	router.HandleFunc("/",handlers.Handler);
	n := negroni.Classic()
	handler := cors.Default().Handler(router)
	n.UseHandler(handler)
	http.Handle("/", n)
}
