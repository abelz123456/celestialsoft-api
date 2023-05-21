package domain

import "mime/multipart"

type UploadFileData struct {
	UserOid string
	File    multipart.FileHeader
}

type ReplaceFileData struct {
	FileUID string
	File    multipart.FileHeader
}
