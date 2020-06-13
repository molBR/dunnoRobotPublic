package main

import (
	"fmt"
  "dunnorobot/cmd/routes"
  "net/http"
  "log"
  "os"
)


func main() {
  port := os.Getenv("PORT")
  fmt.Println("Hello world")
  router := routes.CreateRouter()
  log.Fatal(http.ListenAndServe(":"+ port , router))
}