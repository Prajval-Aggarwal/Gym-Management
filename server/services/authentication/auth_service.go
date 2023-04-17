package authentication

import (
	"fmt"
	"gym/server/db"
	"gym/server/model"
	"gym/server/request"
	"gym/server/response"
	"gym/server/utils"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

var twilioClient *twilio.RestClient

func TwilioInit(password string) {
	twilioClient = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: utils.TWILIO_ACCOUNT_SID,
		Password: password,
	})
}

func AdminRegisterService(context *gin.Context, adminRequest request.RegisterRequest) {

	var credential model.Credential
	var existRecord model.Credential

	credential.UserName = adminRequest.Username
	credential.Contact = adminRequest.Contact
	credential.Role = "admin"

	err := db.FindById(&existRecord, adminRequest.Contact, "contact")
	if err == nil {
		response.ErrorResponse(context, 400, "Account already exists")
		return
	}
	err = db.CreateRecord(&credential)
	if err != nil {
		response.ErrorResponse(context, 500, err.Error())
		return
	}
	response.Response(context, 200, credential)
}

func UserRegisterService(context *gin.Context, userRequest request.RegisterRequest) {

	var credential model.Credential
	var existRecord model.Credential

	credential.UserName = userRequest.Username
	credential.Contact = userRequest.Contact
	credential.Role = "admin"

	err := db.FindById(&existRecord, userRequest.Contact, "contact")
	if err == nil {
		response.ErrorResponse(context, 400, "Account already exists")
		return
	}
	err = db.CreateRecord(&credential)
	if err != nil {
		response.ErrorResponse(context, 500, err.Error())
		return
	}
	response.Response(context, 200, credential)
}

func SendOtpService(context *gin.Context, phoneNumber request.SendOtpRequest) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM credentials WHERE contact=?)"
	err := db.QueryExecutor(query, &exists, phoneNumber.Contact)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	// Response
	if !exists {
		response.ErrorResponse(context, 409, "Number do not exists, please register first")
		return
	}
	ok, sid := sendOtp("+91" + phoneNumber.Contact)
	fmt.Println("SID is", sid)
	if ok {
		response.ShowResponse("Success", 200, "OTP send sucessfully", sid, context)
	}
}
func sendOtp(to string) (bool, *string) {
	fmt.Println("sahdvasasjfjasfjsaf")
	params := &openapi.CreateVerificationParams{}
	params.SetTo(to)

	params.SetChannel("sms")

	resp, err := twilioClient.VerifyV2.CreateVerification(os.Getenv("VERIFY_SERVICE_SID"), params)
	if err != nil {
		fmt.Println("bbkjfbkdsfbkaj")
		return false, nil
	} else {
		return true, resp.Sid
	}

}
func VerifyOtpService(context *gin.Context, verifyOtp request.VerifyOtpRequest) {
	if CheckOtp("+91"+verifyOtp.Contact, verifyOtp.Otp) {
		fmt.Println("verification sucess")
		//create a jwt token adn store it in a cookie
	} else {
		response.ErrorResponse(context, 401, "Verification Failed")
		return
	}
}

// OTP code verification
func CheckOtp(to string, code string) bool {
	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo(to)
	params.SetCode(code)
	resp, err := twilioClient.VerifyV2.CreateVerificationCheck(os.Getenv("VERIFY_SERVICE_SID"), params)

	if err != nil {
		return false
	} else if *resp.Status == "approved" {
		return true
	} else {
		return false
	}
}
