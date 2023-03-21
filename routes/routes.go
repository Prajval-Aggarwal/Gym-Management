package routes

import (
	"fmt"
	db "gym-api/Database"
	cont "gym-api/controllers"
	mod "gym-api/models"
	"log"
	"net/http"
)

func Routes() {
	fmt.Println("Listening on port:8000")
	mux := http.NewServeMux()

	err := db.Connect()
	if err != nil {
		panic(err)

	}
	if db.DB.Migrator().HasTable(&mod.Slot{}) {
		var slots []mod.Slot
		query := "SELECT slots.id,slots.start_time,slots.end_time,slots.available_space FROM slots ORDER BY id ASC;"
		db.DB.Raw(query).Scan(&slots)
		if slots == nil {
			cont.SlotDistribution()
		}
	}

	//User Routes
	mux.HandleFunc("/createuser", cont.CreateUserHandler)
	mux.HandleFunc("/getUsers", cont.GetUsers)
	mux.HandleFunc("/userAttendence", cont.UserAttendence)
	mux.HandleFunc("/getUserID", cont.GetUserbyId)

	//Payment Routes
	mux.HandleFunc("/makepayment", cont.MakepaymentHandler)

	//subscription Routes
	mux.HandleFunc("/createsubs", cont.CreateSubsHandler)
	mux.HandleFunc("/end-memberShip", cont.EndSubscription)
	mux.HandleFunc("/updateMemberShip", cont.UpdateSubscription)
	mux.HandleFunc("/getMemberShip", cont.GetSubscriptionsHandler)

	//Employee Routes
	mux.HandleFunc("/createEmp", cont.CreateEmphandler)
	mux.HandleFunc("/getEmp", cont.GetEmployees)
	mux.HandleFunc("/updateEmp", cont.SetEmpRole)
	mux.HandleFunc("/getEmpRole", cont.GetEmpRole)
	mux.HandleFunc("/empWithuser", cont.GetEmployeesWithUsers)
	mux.HandleFunc("/empAttendence", cont.EmpAttendence)

	//Prices routes
	mux.HandleFunc("/getPrice", cont.GetPrices)
	mux.HandleFunc("/priceUpdate", cont.PriceUpdateHandler)

	//slots routes
	mux.HandleFunc("/slotUpdate", cont.SlotUpdateHandler)


	//Authentication Handler
	mux.HandleFunc("/register", cont.RegisterHandler)
	mux.HandleFunc("/login", cont.LoginHandler)

	//OTP verfication routes
	mux.HandleFunc("/sendotp", cont.SendOTP)
	mux.HandleFunc("/verifyotp", cont.CheckOTP)



	//Listening to the server
	log.Fatal(http.ListenAndServe(":8000", mux))



}
