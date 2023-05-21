package mockdata

import (
	"time"

	"github.com/abelz123456/celestial-api/entity"
	"github.com/google/uuid"
)

var BankMock = []entity.Bank{
	{
		Oid:          uuid.New().String(),
		BankCode:     "MOCK-BRI",
		BankName:     "Mock Bank Rakyat Indonesia",
		UserInserted: uuid.New().String(),
		InsertedDate: time.Now(),
		LastUpdate:   time.Now(),
	},
	{
		Oid:          uuid.New().String(),
		BankCode:     "MOCK-BTN",
		BankName:     "Mock Bank Tabungan Negara",
		UserInserted: uuid.New().String(),
		InsertedDate: time.Now(),
		LastUpdate:   time.Now(),
	},
}
