package authentication

import (
	"errors"
	"fmt"
	"gym/server/db"
	"gym/server/model"
	"gym/server/provider"
	"gym/server/request"
	"gym/server/response"
	"gym/server/utils"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
	"gorm.io/gorm"
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
	credential.UserName = adminRequest.Username
	credential.Contact = adminRequest.Contact
	credential.Role = "admin"

	if db.RecordExist("credentials", "contact", adminRequest.Contact) {
		response.ErrorResponse(context, 400, "Admin already registerd")
		return
	}

	if db.RecordExist("users", "contact", adminRequest.Contact) {
		response.ErrorResponse(context, 400, "Admin cannot register as user")
		return
	}

	err := db.CreateRecord(&credential)
	if err != nil {
		response.ErrorResponse(context, 500, err.Error())
		return
	}

	response.Response(context, 200, credential)
}

func SendOtpService(context *gin.Context, phoneNumber request.SendOtpRequest) {
	var exists1 bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE contact_no=?)"
	err := db.QueryExecutor(query, &exists1, phoneNumber.Contact)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	// Response
	if !exists1 {
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
		//phone number check
		var tokenClaims model.Claims
		var admin model.Credential
		var user model.User
		fmt.Println("sdgsg")
		err := db.FindById(&admin, verifyOtp.Contact, "contact")
		if errors.Is(err, gorm.ErrRecordNotFound) {

			err := db.FindById(&user, verifyOtp.Contact, "contact_no")
			if err != nil {
				response.ErrorResponse(context, 500, "Error finding in DB")
				return
			}
			tokenClaims.Id = user.User_Id
			tokenClaims.Role = "user"
		} else {
			tokenClaims.Id = admin.UserID
			tokenClaims.Role = "admin"
		}
		user.IsActive = true
		db.UpdateRecord(&user, user.User_Id, "user_id")
		tokenString := provider.GenerateToken(tokenClaims, context)
		provider.SetCookie(context, tokenString)

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

func LogoutService(context *gin.Context, tokenString string) {

	provider.DeleteCookie(context)
	var blacklist model.BlackListedToken
	blacklist.Token = tokenString
	db.CreateRecord(&blacklist)

	var user model.User
	claims, err := provider.DecodeToken(tokenString)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
	}
	db.FindById(&user, &claims.RegisteredClaims.ID, "user_id")
	user.IsActive = false
	db.UpdateRecord(&user, &claims.RegisteredClaims.ID, "user_id")

}
