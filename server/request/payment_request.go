package request

type CreatePaymentRequest struct {
	UserId      string `json:"userId" validate:"required"`
	PaymentType string `json:"paymentType" validate:"required"`
}
