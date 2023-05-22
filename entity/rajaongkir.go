package entity

import "time"

type Rajaongkir struct {
	UID         string                 `json:"uid"         gorm:"column:uid;size:65;not null;uniqueIndex;primary_key"`
	HashData    string                 `hashData:"-"       gorm:"column:hashData;size:256"`
	Origin      int                    `json:"origin"      gorm:"column:origin"`
	Destination int                    `json:"destination" gorm:"column:destination"`
	Wight       float64                `json:"wight"       gorm:"column:weight"`
	Courier     string                 `json:"courier"     gorm:"column:courier"`
	APIResponse string                 `json:"-"           gorm:"column:apiResponse;type:longtext"`
	CreatedAt   *time.Time             `json:"createdAt"   gorm:"column:createdAt"`
	CreatedBy   string                 `json:"createdBy"   gorm:"column:createdBy"`
	Response    map[string]interface{} `json:"response"    gorm:"-"`
	ApiStatus   int                    `json:"-"   gorm:"column:apiStatus"`
}

// TableName specifies the table name for the PermissionPolicyUser model
func (Rajaongkir) TableName() string {
	return "rajaongkir"
}
