package routes

import (
	"fmt"
	cont "gym-api/Controllers"
	db "gym-api/Database"
	mod "gym-api/Models"
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
	mux.HandleFunc("/getUsers", cont.GetUsersHandler)
	mux.HandleFunc("/userAttendence", cont.UserAttendenceHandler)
	mux.HandleFunc("/getUserID", cont.GetUserHandler)
	
	//Payment Routes
	mux.HandleFunc("/makepayment", cont.MakePaymentHandler)
	mux.HandleFunc("/payment/status", cont.PaymentStatusHandler)

	//subscription Routes
	mux.HandleFunc("/createsubs", cont.CreatSubscriptionHandler)
	mux.HandleFunc("/end-memberShip", cont.EndSubscriptionHandler)
	mux.HandleFunc("/updateMemberShip", cont.UpdateSubscriptionHandler)
	mux.HandleFunc("/getMemberShip", cont.GetSubscriptionsHandler)

	//Employee Routes
	mux.HandleFunc("/createEmp", cont.CreateEmployeehandler)
	mux.HandleFunc("/getEmp", cont.GetEmployeesHandler)
	mux.HandleFunc("/updateEmp", cont.EmployeeRoleHandler)
	mux.HandleFunc("/getEmpRole", cont.GetEmployeeRoleHandler)
	mux.HandleFunc("/empWithuser", cont.GetUsersWithEmployeesHandler)
	mux.HandleFunc("/empAttendence", cont.EmployeeAttendenceHandler)

	//Prices routes
	mux.HandleFunc("/getPrice", cont.GetPricesHandler)
	mux.HandleFunc("/priceUpdate", cont.PriceUpdateHandler)

	//slots routes
	mux.HandleFunc("/slotUpdate", cont.SlotUpdateHandler)

	//Authentication Handler
	mux.HandleFunc("/register", cont.RegisterHandler)
	mux.HandleFunc("/login", cont.LoginHandler)
	mux.HandleFunc("/forgotPassword", cont.ForgotPasswordHandler)
	mux.HandleFunc("/resetPassword", cont.ResetPasswordHandler)

	//OTP verfication routes
	mux.HandleFunc("/sendotp", cont.SendOTPHandler)
	mux.HandleFunc("/verifyotp", cont.CheckOTPHandler)


	//Equipment routes
	mux.HandleFunc("/createEquipment", cont.CreateEquipmentHandler)
	mux.HandleFunc("/getEquipment", cont.EquimentListHandler)

	//Api documentation Route
	mux.HandleFunc("/api-doc", cont.APIdocsHandler)

	//Listening to the server
	log.Fatal(http.ListenAndServe(":8000", mux))

}
