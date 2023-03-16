package cont

import (
	"encoding/json"
	"fmt"
	db "gym-api/Database"
	mod "gym-api/models"
	"net/http"
)

func MakepaymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id := r.URL.Query().Get("id")
	fmt.Println("Id is :", id)
	w.Header().Set("Content-Type", "application/json")
	var payment mod.Payment
	var sub mod.Subscription
	json.NewDecoder(r.Body).Decode(&payment)

	db.DB.Where("user_id=?", id).First(&sub)

	var memShip mod.SubsType
	db.DB.Where("subs_name=?", sub.Subs_Name).First(&memShip)

	payment.Amount = float64(sub.Duration) * memShip.Price
	payment.User_Id = id

	fmt.Println("payment.User.User_Id", payment.User.User_Id)

	db.DB.Create(&payment)
	
	// update payment id in subscription when payment is successful
	sub.Payment_Id = payment.Payment_Id
	db.DB.Where("user_id=?", id).Updates(&sub)
	json.NewEncoder(w).Encode(&payment)

}
