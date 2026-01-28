package main

import (
	"fmt"
	"log"
	"os"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main(){

	godotenv.Load(".env")
	
	portString := os.Getenv("PORT")
	if portString == ""{
		log.Fatal("PORT is not found in the env")
	}
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://", "http://"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))


	router.HandleFunc("/healthz", handlerReadiness)
	router.HandleFunc("/error", handlerErr)
	router.Mount("/v1", router)
	
	srv := &http.Server{
		Handler: router,
		Addr: ":"+portString,
	}
	log.Printf("Server starting on port %v",portString)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Port:", portString)

}

