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

func main()  {


	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if(portString==""){
		fmt.Println("PORT is not found in the environment file")
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge: 300, // Maximum value not ignored by any of major browsers
	},
	

		))

	srv := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}

	log.Printf("Server is running on port %v",portString)
	err:=srv.ListenAndServe()
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("PORT: ", portString)
}