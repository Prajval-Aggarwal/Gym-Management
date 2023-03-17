package routes

import (
	"fmt"
	db "gym-api/Database"
	cont "gym-api/controllers"
	"log"
	"net/http"
)
func Routes() {
	err := db.Connect()
	if err != nil {
		panic(err)

	}

	//cont.ScheduleDailyEntry()
	cont.UserDailyEntry()
	fmt.Println("Listening on port:8000")
	mux := http.NewServeMux()

	//User Routes
	mux.HandleFunc("/createuser", cont.CreateUserHandler)
	mux.HandleFunc("/getUsers", cont.GetUsers)
	mux.HandleFunc("/userAttendence", cont.UserAttendence)

	//Payment Routes
	mux.HandleFunc("/makepayment", cont.MakepaymentHandler)

	//subscription Routes
	mux.HandleFunc("/createsubs", cont.CreateSubsHandler)
	mux.HandleFunc("/end-memberShip", cont.EndSubscription)
	mux.HandleFunc("/updateMemberShip", cont.UpdateSubscription)

	//Employee Routes
	mux.HandleFunc("/createEmp", cont.CreateEmphandler)

	//Prices routes
	mux.HandleFunc("/getPrice", cont.GetPrices)
	mux.HandleFunc("/priceUpdate", cont.PriceUpdateHandler)

	//Listening to the server
	log.Fatal(http.ListenAndServe(":8000", mux))

}

