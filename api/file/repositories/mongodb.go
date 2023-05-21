package repositories

import (
	"context"
	"errors"
	"io"
	"mime/multipart"
	"os"

	"github.com/abelz123456/celestial-api/entity"
	"github.com/abelz123456/celestial-api/package/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongodb struct {
	Mongo *mongo.Database
	Log   log.Log
}

func (r *mongodb) SaveLocalStorage(ctx context.Context, fileHeader multipart.FileHeader, destination string) error {
	file, err := fileHeader.Open()
	if err != nil {
		r.Log.Error(err, "mongodb.SaveLocalStorage Exception", nil)
		return err
	}
	defer file.Close()

	destinationFile, err := os.Create(destination)
	if err != nil {
		r.Log.Error(err, "mongodb.SaveLocalStorage Exception", nil)
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, file)
	if err != nil {
		r.Log.Error(err, "mongodb.SaveLocalStorage Exception", nil)
		return err
	}

	return nil
}

func (r *mongodb) DeleteLocalStorage(ctx context.Context, filePath string) error {
	return os.Remove(filePath)
}

func (r *mongodb) Save(ctx context.Context, fileData entity.LocalFile) error {
	collection := r.Mongo.Collection("localFile")

	if _, err := collection.InsertOne(ctx, &fileData); err != nil {
		r.Log.Error(err, "mongodb.Save Exception", nil)
		return err
	}

	return nil
}

func (r *mongodb) GetCollection(ctx context.Context) ([]entity.LocalFile, error) {
	var results = make([]entity.LocalFile, 0)
	cursor, err := r.Mongo.Collection("localFile").
		Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var result = new(entity.LocalFile)
		if err := cursor.Decode(&result); err != nil {
			r.Log.Error(err, "mongodb.GetCollection Exception", nil)
			cursor.Close(ctx)
			return nil, err
		}

		results = append(results, *result)
	}

	cursor.Close(ctx)
	return results, nil
}

func (r *mongodb) GetOneByUID(ctx context.Context, uid string) (*entity.LocalFile, error) {
	mongoResult := r.Mongo.Collection("localFile").
		FindOne(ctx, bson.M{"uid": uid})

	if err := mongoResult.Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}

		r.Log.Error(err, "mongodb.GetOneByUID Exception", nil)
		return nil, err
	}

	var dataResult entity.LocalFile
	if err := mongoResult.Decode(&dataResult); err != nil {
		r.Log.Error(err, "mongodb.GetOneByUID Exception", nil)
		return nil, err
	}

	if dataResult.UID == "" {
		return nil, nil
	}

	return &dataResult, nil
}

func (r *mongodb) UpdateOne(ctx context.Context, data entity.LocalFile, newData entity.LocalFile) error {
	collection := r.Mongo.Collection("localFile")

	updateData := bson.M{
		"$set": bson.M{
			"localPath":    newData.LocalPath,
			"originalName": newData.OriginalName,
			"updatedAt":    newData.UpdatedAt,
		},
	}

	_, err := collection.UpdateOne(ctx, bson.M{"uid": data.UID}, updateData)
	if err != nil {
		r.Log.Error(err, "mongodb.UpdateOne Exception", nil)
		return err
	}

	return nil
}

func (r *mongodb) Delete(ctx context.Context, fileData entity.LocalFile) error {
	if _, err := r.Mongo.Collection("localFile").DeleteOne(ctx, bson.M{"uid": fileData.UID}); err != nil {
		r.Log.Error(err, "mongodb.Delete Exception", nil)
		return err
	}

	return nil
}
