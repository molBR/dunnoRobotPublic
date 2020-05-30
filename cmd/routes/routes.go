package routes

import ( 
	"github.com/gorilla/mux"
	"dunnorobot/cmd/comunication/in"
)


func routes (router *mux.Router) bool {
	router.HandleFunc("/", inputInterface.HelloWorld).Methods("POST")
	return true
	
}

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	routes(router)
	return router
}
