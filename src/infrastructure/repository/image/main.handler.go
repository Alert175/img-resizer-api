package image

import (
	"strings"

	"img-resizer-api/src/domain"
	"img-resizer-api/src/infrastructure/db/psql/handlers"
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
	if err := image.loadToVips(); err != nil {
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
	if err := image.loadToVips(); err != nil {
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
	if err := image.loadToVips(); err != nil {
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

// Загрузить изображения по сети и установить необходимые параметры
func (image *ImageModel) InstallFromNetworkAndOptimize(url string, points []Point) ([]string, error) {
	err := image.loadFormNet(url)

	imageDb := handlers.ImageStore{}
	imageDb.ImageLog.Url = url
	imageDb.ImageLog.Status = "success"

	if err != nil {
		return nil, err
	}

	validPath, err := utils.ConvertToPathFormat(url)
	if err != nil {
		imageDb.ImageLog.Status = "error"
		return nil, err
	}

	resultList := []string{}
	originExtension := strings.Split(url, ".")[len(strings.Split(url, "."))-1]

	for _, point := range points {
		imageElement := ImageModel{}
		imageElement.Buffer = image.Buffer
		imageElement.Name = image.Name
		imageElement.Ref = image.Ref
		imageElement.Type = image.Type
		imageElement.Height = image.Height
		imageElement.Width = image.Width

		if err := imageElement.loadToVips(); err != nil {
			imageDb.ImageLog.Status = "error"
			return nil, err
		}
		if err := imageElement.resize(point.Width, point.Height); err != nil {
			imageDb.ImageLog.Status = "error"
			return nil, err
		}
		if err := imageElement.convertTo(point.Format); err != nil {
			imageDb.ImageLog.Status = "error"
			return nil, err
		}
		result, err := imageElement.saveTo(validPath)
		if err != nil {
			imageDb.ImageLog.Status = "error"
			return nil, err
		}
		resultList = append(resultList, result)

		if errOrigin := imageElement.convertTo(originExtension); err != nil {
			imageDb.ImageLog.Status = "error"
			return nil, errOrigin
		}
		_, errOrigin := imageElement.saveTo(validPath)
		if errOrigin != nil {
			imageDb.ImageLog.Status = "error"
			return nil, err
		}
	}

	if err := imageDb.CreateUpdate(); err != nil {
		return nil, err
	}

	return resultList, nil
}

// Загрузить изображения по сети и установить необходимые параметры
// func (image *ImageModel) GetFromNetworkAndResizeAndConvert(url string, height int, width int, format string) (*vips.ImageRef, error) {
// 	err := image.loadFormNet(url)
// 	if err != nil {
// 		logger.Error(err)
// 	}
// 	if err := image.loadToVips(); err != nil {
// 		return nil, err
// 	}
// 	if err := image.resize(width, height); err != nil {
// 		return nil, err
// 	}
// 	if err := image.convertTo(format); err != nil {
// 		return nil, err
// 	}
// 	bytes, err := image.getOutBytes()
// 	if err != nil {
// 	}
// 	return nil, nil
// }
