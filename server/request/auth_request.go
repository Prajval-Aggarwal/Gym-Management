package request

type RegisterRequest struct {
	Username string `json:"username"`
	Contact  string `json:"contact"`
}

type SendOtpRequest struct {
	Contact string `json:"contact"`
}

type VerifyOtpRequest struct {
	Contact string `json:"contact"`
	Otp     string `json:"otp"`
}
