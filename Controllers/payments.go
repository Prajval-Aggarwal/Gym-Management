package cont

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

	} else if sub.Duration == 12 {
		//20% discount
		billamount = (memShip.Price * sub.Duration) * 0.8
		fmt.Fprintln(w, "20% Discount applied")

	} else {
		billamount = memShip.Price * sub.Duration
	}

	payment.Amount = billamount
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
	// creating card params
	// cardParams := &stripe.CardParams{
	// 	Number: stripe.String("4242424242424242"),
	// 	ExpMonth: stripe.String("12"),
	// 	ExpYear: stripe.String("25"),
	// 	CVC: stripe.String("123"),
	// }
	
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

	// Return the payment status
	payment.Status = string(pi1.Status)
	db.DB.Create(&payment)

	// update payment id in subscription when payment is successful
	// if payment.Status == "succeeded" {
	// 	sub.Payment_Id = payment.Payment_Id
	// 	db.DB.Where("user_id=?", id).Updates(&sub)
	// }


	sub.Payment_Id = payment.Payment_Id
	db.DB.Where("user_id=?", id).Updates(&sub)
	// json.NewEncoder(w).Encode(&pi)
	json.NewEncoder(w).Encode(&payment)
	// w.Write([]byte("Payment processed successfully"))
}

func HandlePaymentStatus(w http.ResponseWriter, r *http.Request) {
	// Parse the payment ID from the request URL
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

	// Return the payment status
	status := pi.Status
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": string(status)})
}

