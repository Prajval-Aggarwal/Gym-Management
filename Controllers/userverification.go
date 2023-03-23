package cont

import (
	"fmt"
	"net/http"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

var TWILIO_ACCOUNT_SID string = "AC5869424c6ae2d0b27f66a5d5b9b90485"
var TWILIO_AUTH_TOKEN string = "28625c8fcd3ce5ab1a67b30ebc86221f"
var VERIFY_SERVICE_SID string = "VA6602f8535f8f1369b0ed68eed5d6af67"
var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
	Username: TWILIO_ACCOUNT_SID,
	Password: TWILIO_AUTH_TOKEN,
})

func sendOtp(to string) {
	params := &openapi.CreateVerificationParams{}
	params.SetTo(to)
	params.SetChannel("sms")

	resp, err := client.VerifyV2.CreateVerification(VERIFY_SERVICE_SID, params)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Sent verification '%s'\n", *resp.Sid)
	}
}
func checkOtp(to string) bool {
	var code string
	fmt.Println("Please check your phone and enter the code:")
	fmt.Scanln(&code)

	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo(to)
	params.SetCode(code)

	resp, err := client.VerifyV2.CreateVerificationCheck(VERIFY_SERVICE_SID, params)

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
	if checkOtp("+91" + phNumber) {
		w.Write([]byte("Phone Number verified sucessfully"))
	} else {
		w.Write([]byte("Verifictaion failed"))
	}
}

// email api key := "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"
