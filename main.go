package main

import (
	"fmt"
  "dunnorobot/cmd/routes"
  "net/http"
  "log"
)


func main() {
  fmt.Println("Hello world")
  router := routes.CreateRouter()
  log.Fatal(http.ListenAndServe(":8000", router))
}