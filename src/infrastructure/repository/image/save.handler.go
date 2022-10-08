package image

import (
	"io/ioutil"
	"os"

	"github.com/davidbyttow/govips/v2/vips"
)

const IMAGE_QUALITY = 100

// Сохранить файл в папку
func (image *ImageModel) saveTo(path string) (string, error) {
	outBytes, err := image.getOutBytes()
	if err != nil {
		return "", err
	}

	saveFilePath := "public/" + path
	saveFileName := image.Name + image.Type.FileExt()

	_, errS := os.Stat(saveFilePath)
	if errS != nil {
		if err := os.MkdirAll(saveFilePath, 0o777); err != nil {
			return "", err
		}
	}

	file, err := os.Create(saveFilePath + "/" + saveFileName)
	if err != nil {
		return "", err
	}
	if err := ioutil.WriteFile(saveFilePath+"/"+saveFileName, outBytes, 0o777); err != nil {
		return "", err
	}
	if err := file.Close(); err != nil {
		return "", err
	}
	return saveFilePath + "/" + saveFileName, nil
}

// получить поток байтов для сохранении в файл
func (image *ImageModel) getOutBytes() ([]byte, error) {
	var resultByte []byte
	switch image.Type {
	case 3:
		param := vips.NewPngExportParams()
		param.Quality = IMAGE_QUALITY
		bytes, _, errE := image.Ref.ExportPng(param)
		if errE != nil {
			return resultByte, errE
		} else {
			resultByte = bytes
		}
	case 1:
		param := vips.NewJpegExportParams()
		param.Quality = IMAGE_QUALITY
		bytes, _, errE := image.Ref.ExportJpeg(param)
		if errE != nil {
			return bytes, errE
		} else {
			resultByte = bytes
		}
	case 2:
		param := vips.NewWebpExportParams()
		param.Quality = IMAGE_QUALITY
		bytes, _, errE := image.Ref.ExportWebp(param)
		if errE != nil {
			return bytes, errE
		} else {
			resultByte = bytes
		}
	case 11:
		param := vips.NewAvifExportParams()
		param.Quality = IMAGE_QUALITY
		bytes, _, errE := image.Ref.ExportAvif(param)
		if errE != nil {
			return bytes, errE
		} else {
			resultByte = bytes
		}
	default:
		bytes, _, errE := image.Ref.ExportNative()
		if errE != nil {
			return bytes, errE
		} else {
			resultByte = bytes
		}
	}
	return resultByte, nil
}
