package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
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

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELTET", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/health", handlerReadines)
	v1Router.Get("/error", handlerError)

	router.Mount("/v1", v1Router)




	srver := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}


	log.Printf("The server is runing in port %v", portString);



	err := srver.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(portString)
}
