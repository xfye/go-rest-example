package main

import (
	"github.com/emicklei/go-restful"
	"log"
	"net/http"
	"os"
	"rest-example/database"
	"rest-example/resources/user_resource"
)

func main() {
	dbop, _ := database.NewOperator(os.Getenv("DB_URL"))

	restful.Add(user_resource.New(dbop))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
