package routers

import(
	"github.com/gorilla/mux"
	"github.com/kenkoii/dprojectapi/api/handlers"
)

func SetTopicsRoutes(router *mux.Router) *mux.Router{
	topicRouter := mux.NewRouter()
	//Topics
	topicRouter.HandleFunc("/api/v1/topics", handlers.GetTopicsEndpoint).Methods("GET")
	topicRouter.HandleFunc("/api/v1/topics/{id}", handlers.GetTopicEndpoint).Methods("GET")
	topicRouter.HandleFunc("/api/v1/topics/{id}", handlers.DeleteTopicEndpoint).Methods("DELETE")
	topicRouter.HandleFunc("/api/v1/topics/{id}", handlers.UpdateTopicEndpoint).Methods("PUT")
	topicRouter.HandleFunc("/api/v1/topics", handlers.PostTopicEndpoint).Methods("POST")
	router.PathPrefix("/api/v1/topics").Handler(topicRouter)
	//router.
	return router
}