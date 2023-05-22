package domain

type CostInfoPayload struct {
	Origin      int     `json:"origin"      validate:"required"`
	Destination int     `json:"destination" validate:"required"`
	Weight      float64 `json:"weight"      validate:"required"`
	Courier     string  `json:"courier"     validate:""`
	CreatedBy   string  `json:"-"`
}

type RajaongkirProvince struct {
	ProvinceID   int    `json:"provinceId"`
	ProvinceName string `json:"provinceName"`
}
