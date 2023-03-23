package cont

import (
	"encoding/json"
	"fmt"
	db "gym-api/Database"
	mod "gym-api/Models"
	"net/http"
)

func MakepaymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
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

	var billamount float64
	if sub.Duration == 6 {
		//10% discount
		billamount = (memShip.Price * sub.Duration) * 0.9
		fmt.Fprintln(w, "10% Discount applied")
		payment.Offer = "10%"

	} else if sub.Duration == 12 {
		//20% discount
		billamount = (memShip.Price * sub.Duration) * 0.8
		fmt.Fprintln(w, "20% Discount applied")
		payment.Offer = "10%"

	} else {
		billamount = memShip.Price * sub.Duration
	}

	payment.Amount = (memShip.Price * sub.Duration)
	payment.OfferAmount = billamount
	payment.User_Id = id

	fmt.Println("payment.User.User_Id", payment.User_Id)

	db.DB.Create(&payment)

	// update payment id in subscription when payment is successful
	sub.Payment_Id = payment.Payment_Id
	db.DB.Where("user_id=?", id).Updates(&sub)
	json.NewEncoder(w).Encode(&payment)

}
