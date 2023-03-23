package Controllers

import (
	"encoding/json"
	"fmt"
	db "gym-api/Database"
	mod "gym-api/Models"
	"log"
	"net/http"

	"github.com/stripe/stripe-go/v72"

	"github.com/stripe/stripe-go/v72/paymentintent"
)

func MakePaymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	id := r.URL.Query().Get("id")
	fmt.Println("Id is :", id)
	var user mod.User
	db.DB.Where("user_id = ?", id).Find(&user)
	fmt.Println("user: ", user)
	if user.User_Id == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "User with id %s not found", id)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var payment mod.Payment
	var subscription mod.Subscription
	json.NewDecoder(r.Body).Decode(&payment)

	db.DB.Where("user_id=?", id).First(&subscription)

	var memShip mod.SubsType
	db.DB.Where("subs_name=?", subscription.Subs_Name).First(&memShip)

	var billamount float64
	if subscription.Duration == 6 {
		//10% discount
		billamount = (memShip.Price * subscription.Duration) * 0.9
		fmt.Fprintln(w, "10% Discount applied")
		payment.OfferAmount=billamount
		payment.Offer="10%"

	} else if subscription.Duration == 12 {
		//20% discount
		billamount = (memShip.Price * subscription.Duration) * 0.8
		fmt.Fprintln(w, "20% Discount applied")
		payment.OfferAmount=billamount
		payment.Offer="20%"



	} else {
		billamount = memShip.Price * subscription.Duration
		payment.OfferAmount=billamount

	}

	payment.Amount =  memShip.Price * subscription.Duration
	payment.User_Id = id

	// stripe payment integration
	stripe.Key = "sk_test_51MnxVTSGT1jvrl9CIDO2h1vvRKS0yKYBu0MRagvAcLn9ZshNY7P5CpLLamz6U7rUhx4Bch0Onv03vsoYfg9Bitpv006VIbV229"

	// Get the amount from the request
	amount := billamount
	fmt.Println("amount", amount)
	// Create a new PaymentIntent
	params := &stripe.PaymentIntentParams{
		Amount:             stripe.Int64(int64(amount * 100)),
		Currency:           stripe.String("inr"),
		Description:        stripe.String("Payment"),
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
	}
	pi, err := paymentintent.New(params)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error processing payment", http.StatusInternalServerError)
		return
	}

	params1 := &stripe.PaymentIntentConfirmParams{
		PaymentMethod: stripe.String("pm_card_visa"),
	}

	pi1, err := paymentintent.Confirm(pi.ID, params1)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error processing payment", http.StatusInternalServerError)
		return
	}

	// Check the payment intent status
	switch pi1.Status {
	case "succeeded":
		// Payment succeeded
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Payment processed successfully"))
	case "requires_payment_method":
		// Payment failed
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Payment failed"))
	case "requires_action":
		// Additional action required
		clientSecret := pi1.ClientSecret
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"client_secret": clientSecret,
		})
	default:
		// Unknown status
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Payment requires more actions"))
	}

	payment.Payment_Id = pi.ID
	payment.Status = string(pi1.Status)
	payment.Payment_Type = pi.PaymentMethodTypes[0]
	db.DB.Create(&payment)

	// update payment id in subscription when payment is successful
	// if payment.Status == "succeeded" {
	// 	sub.Payment_Id = payment.Payment_Id
	// 	db.DB.Where("user_id=?", id).Updates(&sub)
	// }

	subscription.Payment_Id = payment.Payment_Id
	db.DB.Where("user_id=?", id).Updates(&subscription)
	json.NewEncoder(w).Encode(&payment)

}

func PaymentStatusHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the payment ID from the request URL
	w.Header().Set("Content-Type", "application/json")
	paymentID := r.URL.Query().Get("payment_id")
	if paymentID == "" {
		http.Error(w, "Payment ID not provided", http.StatusBadRequest)
		return
	}

	// Retrieve the payment details from Stripe
	pi, err := paymentintent.Get(paymentID, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error retrieving payment details", http.StatusInternalServerError)
		return
	}
	status := pi.Status
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": string(status)})
}
