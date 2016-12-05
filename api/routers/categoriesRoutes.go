package routers

import(
	"github.com/gorilla/mux"
	"github.com/kenkoii/dprojectapi/api/handlers"
)

func SetCategoriesRoutes(router *mux.Router) *mux.Router{
	categoriesRouter := mux.NewRouter()
	categoriesRouter.HandleFunc("/api/v1/categories", handlers.GetCategoriesEndpoint).Methods("GET")
	categoriesRouter.HandleFunc("/api/v1/categories/{id}", handlers.GetCategoryEndpoint).Methods("GET")
	categoriesRouter.HandleFunc("/api/v1/categories/{id}", handlers.DeleteCategoryEndpoint).Methods("DELETE")
	categoriesRouter.HandleFunc("/api/v1/categories/{id}", handlers.UpdateCategoryEndpoint).Methods("PUT")
	categoriesRouter.HandleFunc("/api/v1/categories", handlers.PostCategoryEndpoint).Methods("POST")
	router.PathPrefix("/api/v1/categories").Handler(categoriesRouter)
	//router.
	return router
}