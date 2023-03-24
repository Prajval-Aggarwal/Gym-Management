package cont

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	db "gym-api/Database"
	mod "gym-api/models"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"

	razorpay "github.com/razorpay/razorpay-go"
)
type Pagevar struct {
	Orderid string
}
var pagevar Pagevar


type PaymentStatusUpdate struct {
	Entity    string   `json:"entity"`
	AccountID string   `json:"account_id"`
	Event     string   `json:"event"`
	Contains  []string `json:"contains"`
	Payload   struct {
		Payment struct {
			Entity struct {
				ID             string `json:"id"`
				Entity         string `json:"entity"`
				Amount         int    `json:"amount"`
				Currency       string `json:"currency"`
				Status         string `json:"status"`
				OrderID        string `json:"order_id"`
				InvoiceID      string `json:"invoice_id"`
				International  bool   `json:"international"`
				Method         string `json:"method"`
				AmountRefunded int    `json:"amount_refunded"`
				RefundStatus   string `json:"refund_status"`
				Captured       bool   `json:"captured"`
				Description    string `json:"description"`
				CardID         string `json:"card_id"`
				Bank           string `json:"bank"`
				Wallet         string `json:"wallet"`
				Vpa            string `json:"vpa"`
				Email          string `json:"email"`
				Contact        string `json:"contact"`
				Notes          struct {
					Address string `json:"address"`
				} `json:"notes"`
				Fee              int    `json:"fee"`
				Tax              int    `json:"tax"`
				ErrorCode        string `json:"error_code"`
				ErrorDescription string `json:"error_description"`
				ErrorSource      string `json:"error_source"`
				ErrorStep        string `json:"error_step"`
				ErrorReason      string `json:"error_reason"`
				AcquirerData     struct {
					BankTransactionID string `json:"bank_transaction_id"`
				} `json:"acquirer_data"`
				CreatedAt  int64 `json:"created_at"`
				BaseAmount int   `json:"base_amount"`
			} `json:"entity"`
		} `json:"payment"`
	} `json:"payload"`
	CreatedAt int64 `json:"created_at"` 
}

//payment response struct

type paymentresponse struct {

	paymentID string	
	Amount int	
	Status string	
	orderId string

}
var paymentRes paymentresponse










func MakepaymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		
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
		billamount=billamount*100
		// fmt.Fprintln(w, "10% Discount applied")
		payment.Offer = "10%"

	} else if sub.Duration == 12 {
		//20% discount
		billamount = (memShip.Price * sub.Duration) * 0.8
		billamount=billamount*100

		// fmt.Fprintln(w, "20% Discount applied")
		payment.Offer = "20%"

	} else {
		billamount = memShip.Price * sub.Duration
		billamount=billamount*100

	}

	// payment.Amount = (memShip.Price * sub.Duration)
	// payment.OfferAmount = billamount
	payment.User_Id = id
	db.DB.Create(&payment)


	//ORDER CREATION------------------------------------------------------>


	client := razorpay.NewClient("rzp_test_MLjFMJxEVuaLjd", os.Getenv("API_SecretKey"))

	data := map[string]interface{}{
		"amount":   billamount,        
		"currency": "INR",
		"notes": map[string]interface{}{

        "subscription":sub.Subs_Name,
		"Duration":sub.Duration,},
	}
	Body, err := client.Order.Create(data, nil)

	if err != nil {
		fmt.Println("error")
	}

	order_id := Body["id"].(string)

	pagevar.Orderid = order_id

	// fmt.Println("orderId",order_id)
	// t, err := template.ParseFiles("/templates/app.html")

	// if err != nil {

	// 	fmt.Println("template parsing err", err)
	// }

	
// Template
	t, err := template.ParseFiles("controllers/app.html")

	err = t.Execute(w, pagevar)
	if err != nil {

		fmt.Println("template executing error", err)
	}

	// error:=body["error"].(string)
	// fmt.Println("error in failure",error)

	fmt.Println("body response", Body)
	fmt.Println("")

		

	//update during order creation
		payment.User_Id=id
		payment.OrderID=order_id
		db.DB.Create(&payment)
		
		
		
		sub.Payment_Id=paymentRes.paymentID
		
		db.DB.Where("user_id=?", id).Updates(&sub)

	

}

func Response(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	id := r.URL.Query().Get("id")
	fmt.Println("Id is :", id)
	w.Header().Set("Content-Type", "application/json")

	// fmt.Println("Response body",string(body))
	var response PaymentStatusUpdate
	json.Unmarshal(body, &response)
	fmt.Println("")
	// fmt.Println("response",response)
	fmt.Println("id", response.Payload.Payment.Entity.ID)
	fmt.Println("amount", (response.Payload.Payment.Entity.Amount)/100)
	fmt.Println("status", response.Payload.Payment.Entity.Status)
	//put all the response data to paymentresponse struct
	

	var payment mod.Payment
	
	//updates after response
	payment.Payment_Id=response.Payload.Payment.Entity.ID
	payment.Amount=float64(response.Payload.Payment.Entity.Amount)/100
	payment.Status=response.Payload.Payment.Entity.Status
	db.DB.Where("order_id",response.Payload.Payment.Entity.OrderID).Updates(&payment)
	
	var sub mod.Subscription

	sub.Payment_Id=response.Payload.Payment.Entity.ID
	


	//Signature verification
	signature := r.Header.Get("X-Razorpay-Signature")
	fmt.Println("signature", signature)
	if !VerifyWebhookSignature(body, signature, "Qi5HgPzuGUeKjQMOyiTUGjs8") {
		http.Error(w, "Invalid signature", http.StatusUnauthorized)
		return
	} else {

		fmt.Println("signature verified")
	}



	

}

func VerifyWebhookSignature(body []byte, signature string, secret string) bool {

	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	return err
	// }

	h := hmac.New(sha256.New, []byte(secret))
	h.Write(body)
	hash := hex.EncodeToString(h.Sum(nil))

	return hash == signature
}






