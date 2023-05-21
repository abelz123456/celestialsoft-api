package mockdata

import (
	"time"

	"github.com/abelz123456/celestial-api/entity"
	"github.com/google/uuid"
)

var LocalFileMock = []entity.LocalFile{
	{
		UID:          uuid.New().String(),
		LocalPath:    ".test_public/test.jpg",
		OriginalName: "test.jpg",
		UploadedBy:   uuid.New().String(),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	},
	{
		UID:          uuid.New().String(),
		LocalPath:    ".test_public/ltest.jpg",
		OriginalName: "ltest.jpg",
		UploadedBy:   uuid.New().String(),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	},
}
