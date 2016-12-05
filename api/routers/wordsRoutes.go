package routers

import(
	"github.com/gorilla/mux"
	"github.com/kenkoii/dprojectapi/api/handlers"
)

func SetWordsRoutes(router *mux.Router) *mux.Router{
	wordsRouter := mux.NewRouter()
	wordsRouter.HandleFunc("/api/v1/words", handlers.GetWordsEndpoint).Methods("GET")
	wordsRouter.HandleFunc("/api/v1/words/{id}", handlers.GetWordEndpoint).Methods("GET")
	wordsRouter.HandleFunc("/api/v1/words/{id}", handlers.DeleteWordEndpoint).Methods("DELETE")
	wordsRouter.HandleFunc("/api/v1/words/{id}", handlers.UpdateWordEndpoint).Methods("PUT")
	wordsRouter.HandleFunc("/api/v1/words", handlers.PostWordEndpoint).Methods("POST")
	router.PathPrefix("/api/v1/words").Handler(wordsRouter)
	//router.
	return router
}