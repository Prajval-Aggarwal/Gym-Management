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

	//USER
	mux.HandleFunc("/createuser", cont.CreateUserHandler)
	mux.HandleFunc("/getUsers", cont.GetUsers)
	mux.HandleFunc("/end-memberShip", cont.EndSubscription)
	

	//PAYMENT
	mux.HandleFunc("/makepayment", cont.MakepaymentHandler)

	//SUBSCRIPTION
	mux.HandleFunc("/createsubs", cont.CreateSubsHandler)

	//GYM EMPLOYEE
	mux.HandleFunc("/createEmp", cont.CreateEmphandler)

	//PRICE
	mux.HandleFunc("/getPrice", cont.GetPrices)
	mux.HandleFunc("/priceUpdate", cont.PriceUpdateHandler)


	//EQUIPMENTS
	mux.HandleFunc("/updateEquipment",cont.UpdateEquipHandler)
	mux.HandleFunc("/equipmentList",cont.GetEquipList)
	
	log.Fatal(http.ListenAndServe(":8000", mux))
}
