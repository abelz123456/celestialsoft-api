package domain

type PayloadRegister struct {
	EmailName string `json:"emailName" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}

type PayloadLogin struct {
	EmailName string `json:"emailName" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}

type PermissionPolicyUserVm struct {
	Oid         string  `json:"oid"`
	CompanyName *string `json:"companyName"`
	Address     *string `json:"address"`
	EmailName   string  `json:"emailName"`
	Password    string  `json:"password"`
	Token       *string `json:"token"`
}

type PermissionPolicyUserAuthVm struct {
	Oid          string  `json:"oid"`
	CompanyName  *string `json:"companyName"`
	Address      *string `json:"address"`
	EmailName    string  `json:"emailName"`
	Token        *string `json:"token"`
	RefreshToken *string `json:"refreshToken"`
	Role         *string `json:"role"`
}
