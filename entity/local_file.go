package entity

import "time"

type LocalFile struct {
	UID          string    `json:"uid"           gorm:"column:uid;size:65;not null;uniqueIndex;primary_key"  bson:"uid"`
	LocalPath    string    `json:"localPath"     gorm:"column:localPath;size:256"                            bson:"localPath"`
	OriginalName string    `json:"originalName"  gorm:"column:originalName;size:256"                         bson:"originalName"`
	UploadedBy   string    `json:"uploadedBy"    gorm:"column:uploadedBy;size:256"                           bson:"uploadedBy"`
	CreatedAt    time.Time `json:"createdAt"     gorm:"column:createdAt"                                     bson:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"     gorm:"column:updatedAt"                                     bson:"updatedAt"`
	FileURL      string    `json:"fileURL"       gorm:"-"                                                    bson:"-"`
}

func (LocalFile) TableName() string {
	return "localFile"
}
