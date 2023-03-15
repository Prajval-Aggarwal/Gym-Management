package main

import (
	"fmt"
	cont "gym-api/controllers"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Listening on port:8000")
	mux := http.NewServeMux()
	mux.HandleFunc("/createuser", cont.CreateUserHandler)
	mux.HandleFunc("/makepayent", cont.MakepaymentHandler)
	mux.HandleFunc("/createsubs", cont.CreateSubsHandler)
	mux.HandleFunc("/createEmp", cont.CreateEmphandler)
	mux.HandleFunc("/getPrice", cont.GetPrices)
	mux.HandleFunc("/priceUpdate", cont.PriceUpdateHandler)

	mux.HandleFunc("/getUsers", cont.GetUsers)
	mux.HandleFunc("/end-memberShip", cont.EndSubscription)
	log.Fatal(http.ListenAndServe(":8000", mux))
}
