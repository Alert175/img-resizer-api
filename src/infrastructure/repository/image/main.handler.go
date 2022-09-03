package image

import (
	"img-resizer-api/src/domain"
	"img-resizer-api/src/infrastructure/pkg/utils"
	"img-resizer-api/src/infrastructure/pkg/utils/logger"
)

type ImageModel struct {
	domain.ImageModel
}

// Загрузить изображения по сети и установить необходимые параметры
func (image *ImageModel) InstallFromNetwork(url string) (string, error) {
	err := image.loadFormNet(url)
	if err != nil {
		logger.Error(err)
	}

	validPath, err := utils.ConvertToPathFormat(url)
	if err != nil {
		return "", err
	}
	result, err := image.saveTo(validPath)
	if err != nil {
		return "", err
	}
	return result, nil
}

// Загрузить изображения по сети и установить необходимые параметры
func (image *ImageModel) InstallFromNetworkAndResize(url string, height int, width int) (string, error) {
	err := image.loadFormNet(url)
	if err != nil {
		logger.Error(err)
	}

	validPath, err := utils.ConvertToPathFormat(url)
	if err != nil {
		return "", err
	}
	if err := image.resize(width, height); err != nil {
		return "", err
	}
	result, err := image.saveTo(validPath)
	if err != nil {
		return "", err
	}
	return result, nil
}

// Загрузить изображения по сети и установить необходимые параметры
func (image *ImageModel) InstallFromNetworkAndResizeAndConvert(url string, height int, width int, format string) (string, error) {
	err := image.loadFormNet(url)
	if err != nil {
		logger.Error(err)
	}

	validPath, err := utils.ConvertToPathFormat(url)
	if err != nil {
		return "", err
	}
	if err := image.resize(width, height); err != nil {
		return "", err
	}
	if err := image.convertTo(format); err != nil {
		return "", err
	}
	result, err := image.saveTo(validPath)
	if err != nil {
		return "", err
	}
	return result, nil
}
