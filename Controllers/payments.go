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

func StripePayment(amount float64, w http.ResponseWriter) (pi, pi1 *stripe.PaymentIntent) {
	// stripe payment integration
	stripe.Key = "sk_test_51MnxVTSGT1jvrl9CIDO2h1vvRKS0yKYBu0MRagvAcLn9ZshNY7P5CpLLamz6U7rUhx4Bch0Onv03vsoYfg9Bitpv006VIbV229"

	// Get the amount from the request
	// amount := billamount
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

	pi1, err = paymentintent.Confirm(pi.ID, params1)
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

	return pi, pi1

}

func MakePaymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	id := r.URL.Query().Get("id")
	fmt.Println("Id is :", id)
	var u mod.User
	db.DB.Where("user_id = ?", id).Find(&u)
	fmt.Println("user: ", u)
	if u.User_Id == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "User with id %s not found", id)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var payment mod.Payment
	var sub mod.Subscription
	json.NewDecoder(r.Body).Decode(&payment)

	db.DB.Where("user_id=?", id).First(&sub)

	var memShip mod.SubsType
	db.DB.Where("subs_name=?", sub.Subs_Name).First(&memShip)

	var billamount float64
	totalAmount := memShip.Price * sub.Duration
	if sub.Duration == 6 {
		//10% discount
		billamount = totalAmount * 0.9
		payment.Offer = "10%"
		fmt.Fprintln(w, "10% Discount applied")

	} else if sub.Duration == 12 {
		//20% discount
		billamount = totalAmount * 0.8
		payment.Offer = "20%"
		fmt.Fprintln(w, "20% Discount applied")

	} else {
		payment.Offer = "0"
		billamount = totalAmount
	}
	payment.Amount = totalAmount
	payment.OfferAmount = billamount
	payment.User_Id = id
	pi, pi1 := StripePayment(billamount, w)
	payment.Payment_Id = pi.ID
	// Return the payment status
	payment.Status = string(pi1.Status)
	db.DB.Create(&payment)

	sub.Payment_Id = payment.Payment_Id
	db.DB.Where("user_id=?", id).Updates(&sub)
	json.NewEncoder(w).Encode(&payment)
}

func PaymentStatusHandler(w http.ResponseWriter, r *http.Request) {
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
