package entity

import (
	"time"
)

type Bank struct {
	Oid          string    `json:"oid" gorm:"column:oid;size:65;not null;uniqueIndex;primary_key"`
	BankCode     string    `json:"bankCode" gorm:"column:bankCode;size:65;not null;uniqueIndex"`
	BankName     string    `json:"bankName" gorm:"column:bankName;size:100;not null"`
	UserInserted string    `json:"userInserted" gorm:"column:userInserted;size:65"`
	InsertedDate time.Time `json:"insertedDate" gorm:"column:insertedDate"`
	LastUserId   string    `json:"lastUserId" gorm:"column:lastUserId;size:65"`
	LastUpdate   time.Time `json:"lastUpdate" gorm:"column:lastUpdate"`
}

// TableName specifies the table name for the PermissionPolicyUser model
func (Bank) TableName() string {
	return "bank"
}
