package domain

import (
	"img-resizer-api/src/infrastructure/pkg/utils/logger"
	"img-resizer-api/src/infrastructure/repository/image"
)

type ImageModel struct {
	Buffer    []byte // буфер данных изображения
	Name      string
	Extension string
	Height    int
	Width     int
}

// Загрузить изображения по сети и установить необходимые параметры
func (imageModel *ImageModel) InstallFromNetwork(url string) error {
	err := image.InstallNet(url)
	if err != nil {
		logger.Error(err)
	}
	return nil
}

// type ImageUseCase interface {
// 	InstallFromNetwork(url string) error
// }
