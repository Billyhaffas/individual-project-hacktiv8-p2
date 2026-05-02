package payment

type PaymentInfo struct {
	Name      string `json:"name"`
	BankCover string `json:"bank_cover"`
}
type GetRentBookPaymentMethod struct {
	Name string `json:"name"`
}
