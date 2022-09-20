package handlers

import (
	"img-resizer-api/src/infrastructure/db/psql"
	"img-resizer-api/src/infrastructure/db/psql/models"
)

type ImageStore struct {
	models.ImageLog
}

func (image *ImageStore) CreateUpdate() error {
	content := []models.ImageLog{}
	if err := psql.DB.Find(&content, models.ImageLog{Url: image.Url}).Error; err != nil {
		return err
	}

	record := models.ImageLog{}
	record.Url = image.Url
	record.Status = image.Status

	if len(content) == 0 {
		if err := psql.DB.Create(&record).Error; err != nil {
			return err
		}
	} else {
		record.ID = content[0].ID
		if err := psql.DB.Save(&record).Error; err != nil {
			return err
		}
	}
	return nil
}
