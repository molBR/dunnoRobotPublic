package routes

import ( 
	"github.com/gorilla/mux"
	"dunnorobot/cmd/comunication/in"
	"github.com/rs/cors"
	"net/http"

)


func routes (router *mux.Router) http.Handler {
	router.HandleFunc("/", inputInterface.HelloWorld).Methods("POST")
	
    c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowCredentials: true,
		})	

	handler := c.Handler(router)
	return handler
	
}

func CreateRouter() http.Handler {
	router := mux.NewRouter()
	routesWithHandlers:=routes(router)
	return routesWithHandlers
}
