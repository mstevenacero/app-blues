package main

import (
	"fmt"
	"net/http"

	database "github.com/mstevenacero/application-core-blues/database"
	routes "github.com/mstevenacero/application-core-blues/routes"
)

func main() {
	database.ConnectToMongo()
	router := routes.RegisterRouter()
	fmt.Println("Server started at PORT:3000")
	if err := http.ListenAndServe(":3000", router); err != nil {
		fmt.Println("error with  route in backend")
	}

}
