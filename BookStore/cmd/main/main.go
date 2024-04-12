package main

import (
	"fmt"
	//"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/nikhilkumar9687/GoLangProjects/BookStore/pkg/routes"
)

func main() {
	fmt.Println("In Main.go")
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Starting Server at 8000")
	if err := http.ListenAndServe("localhost:9010", r); err != nil {
        fmt.Println("Server failed to start: ", err)
	}

}