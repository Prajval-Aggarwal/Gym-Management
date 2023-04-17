package payment

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"gym/server/db"
	"gym/server/model"
	"gym/server/request"
	"gym/server/response"
	"io/ioutil"
	"os"
	"text/template"
	"time"

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
}
var pagevar Pagevar


func MakePaymentService(context *gin.Context, PaymentData request.CreatePaymentRequest) {
	var payment model.Payment
	var subscription model.Subscription
	var membership model.Membership

	payment.Payment_Type = PaymentData.PaymentType
	err := db.FindById(&subscription, PaymentData.UserId, "user_id")
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

	//add razor pay to add payemnt id for now it is uuid
	err = db.CreateRecord(&payment)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}

	subscription.Payment_Id = payment.Payment_Id

	result := db.UpdateRecord(&subscription, PaymentData.UserId, "user_id")
	if result.Error != nil {
		response.ErrorResponse(context, 400, result.Error.Error())
		return
	}

	userid:="123"
	order_creation(userid,membership,context)
	response.ShowResponse("Success", 200, "Order Created", payment, context)

}


func order_creation(user_id string,membership model.Membership ,context *gin.Context){

	//ORDER CREATION------------------------------------------------------>

	var memship model.Membership
	memship.MemName=membership.MemName
	// er:=db.DB.Where("membership_name=?",membership.MemName).First(&memship).Error
	er:=db.FindById(&memship,membership.MemName,"membership_name")
	if er!=nil {
		// Res.Response("server error",500,er.Error(),"",writer)
		response.ShowResponse("server error", 500, er.Error(), "", context)

	}
	client := razorpay.NewClient("rzp_test_MLjFMJxEVuaLjd", os.Getenv("Razorpay_Key"))

	data := map[string]interface{}{
		"amount":   memship.Price,        
		"currency": "INR",
		"notes": map[string]interface{}{

        "subscription":membership.MemName,
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
	t, err := template.ParseFiles("services/payment/payment.html")
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
		
		payment.Order_id=order_id
		
		// Er:=db.DB.Where("user_id=?",user_id).Updates(&payment).Error
		Er:=db.FindById(&memship,user_id,"user_id")

		if Er!=nil{
			// Res.Response("server error",500,er.Error(),"",writer)
			response.ShowResponse("server error", 500, er.Error(), "", context)

		}


}




func Razorpay_Response(context *gin.Context, PaymentData PaymentStatusUpdate) {
	fmt.Println("Response function called./....")
	// w.Header().Set("Content-Type", "application/json")

	// body, err := ioutil.ReadAll(r.Body)
	
	
	// json.Unmarshal(body, &response)
	fmt.Println("")

	fmt.Println("id", PaymentData.Payload.Payment.Entity.ID)
	fmt.Println("order_id",PaymentData.Payload.Payment.Entity.OrderID)
	fmt.Println("amount", (PaymentData.Payload.Payment.Entity.Amount)/100)
	fmt.Println("status", PaymentData.Payload.Payment.Entity.Status)

	

	//put all the response data to paymentresponse struct
	

	var payment model.Payment
	// err1:=db.DB.Where("order_id=?",response.Payload.Payment.Entity.OrderID).First(&payment).Error
	err1:=db.FindById(&payment,PaymentData.Payload.Payment.Entity.OrderID,"order_id")

	if err1!=nil{
		fmt.Println("error is ",err1)
		// Res.Response("server error",500,err1.Error(),"",w)
		response.ShowResponse("server error", 500, err1.Error(), "", context)

	}
	//updates after response
	payment.Payment_Id=PaymentData.Payload.Payment.Entity.ID
	payment.Status=PaymentData.Payload.Payment.Entity.Status
	payment.CreatedAt=time.Now()

	if payment.Status=="captured"{

		var user model.User

		// db.DB.Where("user_id=?",payment.User_Id).First(&user)
		Er:=db.FindById(&user,payment.User_Id,"user_id")
		if Er!=nil{

		response.ShowResponse("server error", 500, err1.Error(), "", context)

		}

		var subscription model.Subscription
		er:=db.FindById(&subscription,payment.User_Id,"user_id")
		if er!=nil{

			response.ShowResponse("server error", 500, er.Error(), "", context)

		}
		subscription.Payment_Id=payment.Payment_Id
		

		
		

		// db.DB.Where("user_id=?",payment.User_id).Updates(&user)


		
	}

	fmt.Println("Payments is;",payment)
	// dbErr:=db.DB.Where("order_id=?",PaymentData.Payload.Payment.Entity.OrderID).Updates(&payment).Error
	dbErr:=db.UpdateRecord(&payment,PaymentData.Payload.Payment.Entity.OrderID,"order_id").Error

	if dbErr!=nil{
		fmt.Println("db error",dbErr)
		// Res.Response("Bad gateway",500,dbErr.Error(),"",w)
		response.ShowResponse("server error", 500, dbErr.Error(), "", context)

	}

	//Signature verification
	signature := context.Request.Header["X-Razorpay-Signature"]
	fmt.Println("signature", signature)
	reqBody, _ := ioutil.ReadAll(context.Request.Body)
	if !VerifyWebhookSignature(reqBody, signature[0], os.Getenv("API_SecretKey")) {

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