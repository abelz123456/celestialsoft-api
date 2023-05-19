package entity

import "time"

type PermissionPolicyUser struct {
	Oid                 string    `json:"oid" gorm:"column:oid"`
	EmailName           string    `json:"emailName" gorm:"column:emailName"`
	Password            string    `json:"-" gorm:"column:password"`
	LevelUser           int       `json:"levelUser" gorm:"column:levelUser"`
	AndroidToken        string    `json:"androidToken" gorm:"column:androidToken"`
	ExpiredTime         string    `json:"expiredTime" gorm:"column:expiredTime"`
	IsActive            bool      `json:"isActive" gorm:"column:isActive"`
	Description         string    `json:"description" gorm:"column:description"`
	OptimisticLockField int       `json:"optimisticLockField" gorm:"column:optimisticLockField"`
	GCRecord            int       `json:"gCRecord" gorm:"column:gCRecord"`
	Deleted             bool      `json:"deleted" gorm:"column:deleted"`
	UserInserted        string    `json:"userInserted" gorm:"column:userInserted"`
	InsertedDate        time.Time `json:"insertedDate" gorm:"column:insertedDate"`
	LastUserId          string    `json:"lastUserId" gorm:"column:lastUserId"`
	LastUpdate          time.Time `json:"lastUpdate" gorm:"column:lastUpdate"`
	AuthToken           string    `json:"authToken" gorm:"-"`
}

// TableName specifies the table name for the PermissionPolicyUser model
func (PermissionPolicyUser) TableName() string {
	return "permissionPolicyUser"
}
