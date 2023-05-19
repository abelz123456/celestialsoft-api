package mockdata

import (
	"time"

	"github.com/abelz123456/celestial-api/entity"
	"github.com/google/uuid"
)

var PermissionPolicyUserMock = []entity.PermissionPolicyUser{
	{
		Oid:                 uuid.New().String(),
		EmailName:           "test@email.com",
		Password:            "passwordHash",
		Description:         "",
		OptimisticLockField: 0,
		GCRecord:            0,
		Deleted:             false,
		UserInserted:        "",
		InsertedDate:        time.Now(),
		LastUserId:          "",
		LastUpdate:          time.Now(),
	},
	{
		Oid:                 uuid.New().String(),
		EmailName:           "test1@email.com",
		Password:            "passwordHash",
		Description:         "",
		OptimisticLockField: 0,
		GCRecord:            0,
		Deleted:             false,
		UserInserted:        "",
		InsertedDate:        time.Now(),
		LastUserId:          "",
		LastUpdate:          time.Now(),
	},
	{
		Oid:                 uuid.New().String(),
		EmailName:           "test2@email.com",
		Password:            "passwordHash",
		Description:         "",
		OptimisticLockField: 0,
		GCRecord:            0,
		Deleted:             false,
		UserInserted:        "",
		InsertedDate:        time.Now(),
		LastUserId:          "",
		LastUpdate:          time.Now(),
	},
}
