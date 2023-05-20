package entity

import "time"

type PermissionPolicyUser struct {
	Oid                 string    `json:"oid"                   gorm:"column:oid"                   bson:"oid"`
	EmailName           string    `json:"emailName"             gorm:"column:emailName"             bson:"emailName"`
	Password            string    `json:"-"                     gorm:"column:password"              bson:"password"`
	LevelUser           int       `json:"levelUser"             gorm:"column:levelUser"             bson:"levelUser"`
	AndroidToken        string    `json:"androidToken"          gorm:"column:androidToken"          bson:"androidToken"`
	ExpiredTime         string    `json:"expiredTime"           gorm:"column:expiredTime"           bson:"expiredTime"`
	IsActive            bool      `json:"isActive"              gorm:"column:isActive"              bson:"isActive"`
	Description         string    `json:"description"           gorm:"column:description"           bson:"description"`
	OptimisticLockField int       `json:"optimisticLockField"   gorm:"column:optimisticLockField"   bson:"optimisticLockField"`
	GCRecord            int       `json:"gCRecord"              gorm:"column:gCRecord"              bson:"gCRecord"`
	Deleted             bool      `json:"deleted"               gorm:"column:deleted"               bson:"deleted"`
	UserInserted        string    `json:"userInserted"          gorm:"column:userInserted"          bson:"userInserted"`
	InsertedDate        time.Time `json:"insertedDate"          gorm:"column:insertedDate"          bson:"insertedDate"`
	LastUserId          string    `json:"lastUserId"            gorm:"column:lastUserId"            bson:"lastUserId"`
	LastUpdate          time.Time `json:"lastUpdate"            gorm:"column:lastUpdate"            bson:"lastUpdate"`
	AuthToken           string    `json:"authToken"             gorm:"-"                            bson:"-"`
}

// TableName specifies the table name for the PermissionPolicyUser model
func (PermissionPolicyUser) TableName() string {
	return "permissionPolicyUser"
}
