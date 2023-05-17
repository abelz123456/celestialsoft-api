package entity

import "time"

type PermissionPolicyUser struct {
	Oid                 string    `json:"oid"`
	EmailName           string    `json:"emailName"`
	Password            string    `json:"password"`
	LevelUser           int       `json:"levelUser"`
	AndroidToken        *string   `json:"androidToken"`
	ExpiredTime         *string   `json:"expiredTime"`
	IsActive            bool      `json:"isActive"`
	Description         *string   `json:"description"`
	OptimisticLockField int       `json:"optimisticLockField"`
	GCRecord            int       `json:"gCRecord"`
	Deleted             bool      `json:"deleted"`
	UserInserted        *string   `json:"userInserted"`
	InsertedDate        time.Time `json:"insertedDate"`
	LastUserId          *string   `json:"lastUserId"`
	LastUpdate          time.Time `json:"lastUpdate"`
}
