package main

import (
	"fmt"
	"gym-api/routes"

	"github.com/joho/godotenv"
)
func main() {

	envErr:=godotenv.Load(".env")
	if envErr!=nil {

		fmt.Println("could not load environment")
	}

	
	//---------------------------------------------------------->
	routes.Routes()

	
}
