package cont

import (
	"encoding/json"
	"fmt"
	cons "gym-api/utils"
	"net/http"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
	Username: cons.TWILIO_ACCOUNT_SID,
	Password: cons.TWILIO_AUTH_TOKEN,
})

func sendOtp(to string) {
	params := &openapi.CreateVerificationParams{}
	params.SetTo(to)
	params.SetChannel("sms")

	resp, err := client.VerifyV2.CreateVerification(cons.VERIFY_SERVICE_SID, params)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Sent verification '%s'\n", *resp.Sid)
	}
}
func checkOtp(to string, code string) bool {
	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo(to)
	params.SetCode(code)
	resp, err := client.VerifyV2.CreateVerificationCheck(cons.VERIFY_SERVICE_SID, params)

	if err != nil {
		fmt.Println("Error is :", err)
		return false
	} else if *resp.Status == "approved" {
		return true
	} else {
		return false
	}
}

func SendOTP(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	phNumber := r.URL.Query().Get("number")

	sendOtp("+91" + phNumber)
	w.Write([]byte("OTP sent successfully "))

}

func CheckOTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	phNumber := r.URL.Query().Get("number")
	var otp = make(map[string]string)
	json.NewDecoder(r.Body).Decode(&otp)
	if checkOtp("+91"+phNumber, otp["otp"]) {
		w.Write([]byte("Phone Number verified sucessfully"))
	} else {
		w.Write([]byte("Verifictaion failed"))
	}
}

// email api key := "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"
