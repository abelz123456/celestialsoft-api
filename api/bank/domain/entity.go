package domain

type CreateBankDto struct {
	BankCode     string `json:"bankCode"  validate:"required"`
	BankName     string `json:"bankName"  validate:"required"`
	UserInserted string `json:"-"`
}

type UpdateBankDto struct {
	BankName string `json:"bankName"  validate:"required"`
}
