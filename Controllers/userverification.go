package cont

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
	Username: os.Getenv("TWILIO_ACCOUNT_SID"),
	Password: os.Getenv("TWILIO_AUTH_TOKEN"),
})

func SendOtp(to string) {
	params := &openapi.CreateVerificationParams{}
	params.SetTo(to)
	params.SetChannel("sms")

	resp, err := client.VerifyV2.CreateVerification(os.Getenv("VERIFY_SERVICE_SID"), params)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Sent verification '%s'\n", *resp.Sid)
	}
}
func CheckOtp(to string, code string) bool {
	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo(to)
	params.SetCode(code)
	resp, err := client.VerifyV2.CreateVerificationCheck(os.Getenv("VERIFY_SERVICE_SID"), params)

	if err != nil {
		fmt.Println("Error is :", err)
		return false
	} else if *resp.Status == "approved" {
		return true
	} else {
		return false
	}
}

func SendOTPHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	phNumber := r.URL.Query().Get("number")

	SendOtp("+91" + phNumber)
	w.Write([]byte("OTP sent successfully "))

}

func CheckOTPHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	phNumber := r.URL.Query().Get("number")
	var otp = make(map[string]string)
	json.NewDecoder(r.Body).Decode(&otp)
	if CheckOtp("+91"+phNumber, otp["otp"]) {
		w.Write([]byte("Phone Number verified sucessfully"))
	} else {
		w.Write([]byte("Verifictaion failed"))
	}
}
