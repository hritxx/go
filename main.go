package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main()  {
	fmt.Println("Hello, World!")

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if(portString==""){
		fmt.Println("PORT is not found in the environment file")
	}


	fmt.Println("PORT: ", portString)
}