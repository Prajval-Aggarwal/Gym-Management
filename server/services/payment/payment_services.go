package payment

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"gym/server/db"
	"gym/server/model"
	"gym/server/provider"
	"gym/server/request"
	"gym/server/response"
	"os"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/razorpay/razorpay-go"
)


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


type Pagevar struct {
	Orderid string
	Amount float64
}
var pagevar Pagevar


func MakePaymentService(context *gin.Context, PaymentData request.CreatePaymentRequest) {
	var payment model.Payment
	var subscription model.Subscription
	var membership model.Membership

	payment.Payment_Type = PaymentData.PaymentType
	err := db.FindById(&subscription, PaymentData.UserId, "user_id")

	fmt.Println("subscription", subscription)
	// er:=db.DB.Where("membership_name=?",membership.MemName).First(&memship).Error
	//
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	err = db.FindById(&membership, subscription.Subs_Name, "mem_name")
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	var billamount float64
	totalAmount := membership.Price * subscription.Duration
	if subscription.Duration == 6 {
		//10% discount
		billamount = totalAmount * 0.9
		payment.Offer = "10%"

	} else if subscription.Duration == 12 {
		//20% discount
		billamount = totalAmount * 0.8
		payment.Offer = "20%"

	} else {
		payment.Offer = "0"
		billamount = totalAmount
	}
	payment.Amount = totalAmount
	payment.OfferAmount = billamount
	payment.User_Id = PaymentData.UserId

	fmt.Println("paymetn ",payment)
	fmt.Println("/n")

	//add razor pay to add payemnt id for now it is uuid
	err = db.CreateRecord(&payment)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}

	fmt.Println("same Paymnt",payment)
	cookie,_:=context.Request.Cookie("cookie")
	

	
	


	claims,_:=provider.DecodeToken(cookie.Value)
	
	order_creation(claims.Id,billamount,context)
	response.ShowResponse("Success", 200, "Order Created", payment, context)

}


func order_creation(user_id string,billamount float64,context *gin.Context){

	//ORDER CREATION------------------------------------------------------>

	
	var subscription model.Subscription
	db.FindById(&subscription,user_id,"user_id")

	
	client := razorpay.NewClient("rzp_test_MLjFMJxEVuaLjd", os.Getenv("Razorpay_Key"))

	data := map[string]interface{}{
		"amount":  billamount ,        
		"currency": "INR",
		"notes": map[string]interface{}{

        "subscription":subscription.Subs_Name,
		},
	}
	Body, err := client.Order.Create(data, nil)

	if err != nil {
		fmt.Println("error in order create request")
		// Res.Response("access denied",402,err.Error(),"",writer)
		response.ShowResponse("access denied", 402, err.Error(), "", context)

		return
	}

	order_id := Body["id"].(string)

	pagevar.Orderid = order_id

	
// Template
	t, err := template.ParseFiles("server/services/payment/payment.html")
	if err!=nil{
		fmt.Println("template parsing error",err)
	}

	err = t.Execute(context.Writer, pagevar)
	if err != nil {

		fmt.Println("template executing error", err)
	}



	fmt.Println("body response", Body)
	fmt.Println("")

		

	//update during order creation
		var payment model.Payment
		db.FindById(&payment,user_id,"user_id")
		payment.Order_id=order_id
		fmt.Println(";bkdsjkfdsfkjfadand:",payment.Order_id)
		// fmt.Println("payment is",payment)
		// Er:=db.DB.Where("user_id=?",user_id).Updates(&payment).Error
		// Er:=db.FindById(&memship,user_id,"user_id")
		Er:=db.UpdateRecord(&payment, user_id,"user_id").Error
		
	// fmt.Println("fghfhf")
		if Er!=nil{
			// Res.Response("server error",500,er.Error(),"",writer)
			fmt.Println("errro is ",Er.Error())
			//sresponse.ErrorResponse(context, 500, er.Error())
		}
		// fmt.Println("yujghjgj")

}




func Razorpay_Response(context *gin.Context, body []byte) {
	fmt.Println("Response function called./....")
	// w.Header().Set("Content-Type", "application/json")

	// body, err := ioutil.ReadAll(r.Body)
	
	

	var paymentresponse PaymentStatusUpdate
	json.Unmarshal(body, &paymentresponse)
	fmt.Println("")

	fmt.Println("id", paymentresponse.Payload.Payment.Entity.ID)
	fmt.Println("order_id",paymentresponse.Payload.Payment.Entity.OrderID)
	fmt.Println("amount", (paymentresponse.Payload.Payment.Entity.Amount)/100)
	fmt.Println("status", paymentresponse.Payload.Payment.Entity.Status)

	

	//put all the response data to paymentresponse struct
	

	var payment model.Payment
	// err1:=db.DB.Where("order_id=?",response.Payload.Payment.Entity.OrderID).First(&payment).Error
	err1:=db.FindById(&payment,paymentresponse.Payload.Payment.Entity.OrderID,"order_id")

	if err1!=nil{
		fmt.Println("error is ",err1)
		// Res.Response("server error",500,err1.Error(),"",w)
		response.ShowResponse("server error", 500, err1.Error(), "", context)

	}
	//updates after response
	fmt.Println("payment before updation",payment)
	payment.Payment_Id=paymentresponse.Payload.Payment.Entity.ID
	payment.Status=paymentresponse.Payload.Payment.Entity.Status
//	payment.Order_id=paymentresponse.Payload.Payment.Entity.OrderID
	fmt.Println("payment after updation is:",payment)
	
//update payment 

	// err:=db.UpdateRecord(&payment,paymentresponse.Payload.Payment.Entity.OrderID,"order_id").Error
	// if err!=nil{
	// 	fmt.Println("eroronbsjdfkjsdkjfbsd:",err.Error())
	// 	return
	// }
	db.QueryExecutor("UPDATE payments SET status=? ,payment_id=? WHERE order_id=?",nil,paymentresponse.Payload.Payment.Entity.Status,paymentresponse.Payload.Payment.Entity.ID,paymentresponse.Payload.Payment.Entity.OrderID)


	if payment.Status=="captured"{
		
		var subscription model.Subscription
		er:=db.FindById(&subscription,payment.User_Id,"user_id")
		if er!=nil{

			response.ShowResponse("server error", 500, er.Error(), "", context)

		}
		subscription.Payment_Id=payment.Payment_Id


		err:=db.UpdateRecord(&subscription,payment.User_Id,"user_id").Error
		if er!=nil{

			response.ShowResponse("server error", 500, err.Error(), "", context)

		}



		
	}
	fmt.Println("full paymewnt info",payment)
	// fmt.Println("Payments is;",payment.Status)
	// fmt.Println("order id ",paymentresponse.Payload.Payment.Entity.OrderID)
	fmt.Println("orderidd",payment.Order_id)
	// dbErr:=db.DB.Where("order_id=?",PaymentData.Payload.Payment.Entity.OrderID).Updates(&payment).Error


	//Signature verification
	signature := context.Request.Header.Get("X-Razorpay-Signature")
	fmt.Println("signature", signature)
	
	if !VerifyWebhookSignature(body, signature, os.Getenv("Razorpay_Key")) {

		response.ShowResponse("Unauthorized",401,"Invalid signature","",context)
		return
	} else {

		fmt.Println("signature verified")
		response.ShowResponse("OK",200,"Success","",context)
	}


	

}

func VerifyWebhookSignature(body []byte, signature string, secret string) bool {


	h := hmac.New(sha256.New, []byte(secret))
	h.Write(body)
	hash := hex.EncodeToString(h.Sum(nil))

	return hash == signature
}