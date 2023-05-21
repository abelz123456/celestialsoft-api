package services

import (
	"context"
	"fmt"
	"path"
	"time"

	"github.com/abelz123456/celestial-api/api/file/domain"
	"github.com/abelz123456/celestial-api/api/file/repositories"
	"github.com/abelz123456/celestial-api/entity"
	"github.com/abelz123456/celestial-api/package/manager"
	"github.com/google/uuid"
)

type service struct {
	manager    manager.Manager
	repository domain.Repository
}

func NewService(mgr manager.Manager) domain.Service {
	return &service{
		manager:    mgr,
		repository: repositories.NewRepository(mgr),
	}
}

func (s *service) getFileURL(localFile entity.LocalFile) string {
	return fmt.Sprintf(
		"%s://%s/%s",
		s.manager.Config.AppScheme, s.manager.Config.AppHost, localFile.LocalPath,
	)
}

func (s *service) UploadFile(ctx context.Context, fileData domain.UploadFileData) (*entity.LocalFile, error) {
	var (
		fileUID  string = uuid.NewString()
		filePath string = fmt.Sprintf("%s/%s%s", s.manager.Config.StaticFilePath, fileUID, path.Ext(fileData.File.Filename))
	)

	if err := s.repository.SaveLocalStorage(ctx, fileData.File, filePath); err != nil {
		return nil, err
	}

	newData := entity.LocalFile{
		UID:          fileUID,
		LocalPath:    filePath,
		OriginalName: fileData.File.Filename,
		UploadedBy:   fileData.UserOid,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := s.repository.Save(ctx, newData); err != nil {
		s.repository.DeleteLocalStorage(ctx, filePath)
		return nil, err
	}

	newData.FileURL = s.getFileURL(newData)
	return &newData, nil
}

func (s *service) GetCollection(ctx context.Context) ([]entity.LocalFile, error) {
	results, err := s.repository.GetCollection(ctx)
	if err != nil {
		return results, err
	}

	for i, result := range results {
		results[i].FileURL = s.getFileURL(result)
	}

	return results, nil
}

func (s *service) GetInfo(ctx context.Context, uid string) (*entity.LocalFile, error) {
	result, err := s.repository.GetOneByUID(ctx, uid)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, fmt.Errorf("no file data with uid '%s'", uid)
	}

	result.FileURL = s.getFileURL(*result)
	return result, nil
}

func (s *service) ReplaceFile(ctx context.Context, fileData domain.ReplaceFileData) (*entity.LocalFile, error) {
	prevData, err := s.repository.GetOneByUID(ctx, fileData.FileUID)
	if err != nil || prevData == nil {
		if prevData == nil {
			err = fmt.Errorf("no file data with uid '%s'", fileData.FileUID)
		}

		return nil, err
	}

	filePath := fmt.Sprintf("%s/%s%s", s.manager.Config.StaticFilePath, fileData.FileUID, path.Ext(fileData.File.Filename))
	if err := s.repository.SaveLocalStorage(ctx, fileData.File, filePath); err != nil {
		return nil, err
	}

	prevData.LocalPath = filePath
	prevData.OriginalName = fileData.File.Filename
	prevData.UpdatedAt = time.Now()
	prevData.FileURL = s.getFileURL(*prevData)

	if err := s.repository.UpdateOne(ctx, *prevData, *prevData); err != nil {
		s.repository.DeleteLocalStorage(ctx, filePath)
		return nil, err
	}

	return prevData, nil
}

func (s *service) UnlinkFile(ctx context.Context, uid string) error {
	fileData, err := s.repository.GetOneByUID(ctx, uid)
	if err != nil || fileData == nil {
		if fileData == nil {
			err = fmt.Errorf("no file data with uid '%s'", uid)
		}

		return err
	}

	if err := s.repository.Delete(ctx, *fileData); err != nil {
		return err
	}

	s.repository.DeleteLocalStorage(ctx, fileData.LocalPath)
	return nil
}
