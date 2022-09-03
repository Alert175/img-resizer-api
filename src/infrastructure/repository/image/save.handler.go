package image

import (
	"io/ioutil"
	"os"

	"github.com/davidbyttow/govips/v2/vips"
)

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
		bytes, _, errE := image.Ref.ExportPng(vips.NewPngExportParams())
		if errE != nil {
			return resultByte, errE
		} else {
			resultByte = bytes
		}
	case 1:
		bytes, _, errE := image.Ref.ExportJpeg(vips.NewJpegExportParams())
		if errE != nil {
			return bytes, errE
		} else {
			resultByte = bytes
		}
	case 2:
		bytes, _, errE := image.Ref.ExportWebp(vips.NewWebpExportParams())
		if errE != nil {
			return bytes, errE
		} else {
			resultByte = bytes
		}
	case 11:
		bytes, _, errE := image.Ref.ExportAvif(vips.NewAvifExportParams())
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
