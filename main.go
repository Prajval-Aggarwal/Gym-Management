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
	mux.HandleFunc("/makepayment", cont.MakepaymentHandler)

	mux.HandleFunc("/createsubs", cont.CreateSubsHandler)
	mux.HandleFunc("/getsubs", cont.GetSubscriptionsHandler)

	mux.HandleFunc("/createEmp", cont.CreateEmphandler)
	mux.HandleFunc("/getPrice", cont.GetPrices)
	mux.HandleFunc("/priceUpdate", cont.PriceUpdateHandler)

	mux.HandleFunc("/getUsers", cont.GetUsers)
	mux.HandleFunc("/getUserbyId", cont.GetUserbyId)

	mux.HandleFunc("/end-memberShip", cont.EndSubscription)

	mux.HandleFunc("/emp", cont.GetEmployees)
	mux.HandleFunc("/emprole", cont.SetorUpdateEmpRole)

	mux.HandleFunc("/empwithusers", cont.GetEmployeesWithUsers)


	log.Fatal(http.ListenAndServe(":8000", mux))
}