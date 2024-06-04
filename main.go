package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)


func main(){
	fmt.Println("hello")

	godotenv.Load()

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("port is not geting from the env")
	}

	router := chi.NewRouter()

	srver := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}

	err := srver.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

	print(srver)

	fmt.Println(portString)
}
