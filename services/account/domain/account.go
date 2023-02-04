package domain

type Account struct {
	Id                 int    `json:"id"`
	BankName           string `json:"bank_name"`
	BankIdentityNumber string `json:"bank_identity_number"`
}
