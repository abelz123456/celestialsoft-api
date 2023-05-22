package entity

import "time"

type EmailSent struct {
	UID              string                 `json:"uid"       gorm:"column:uid;size:65;not null;uniqueIndex;primary_key"  bson:"uid"`
	SentBy           string                 `json:"sentBy"    gorm:"column:sentBy"`
	Payload          map[string]interface{} `json:"payload"   gorm:"-"`
	StringifyPayload string                 `json:"-"         gorm:"column:payload;type:longtext"`
	SentAt           *time.Time             `json:"sentAt"    gorm:"column:sentAt"`
	SentError        string                 `json:"sentError" gorm:"column:sentError"`
}

// TableName specifies the table name for the PermissionPolicyUser model
func (EmailSent) TableName() string {
	return "emailSent"
}
