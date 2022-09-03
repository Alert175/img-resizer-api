package image

import (
	"io/ioutil"
	"os"

	"github.com/davidbyttow/govips/v2/vips"
)

// Сохранить файл в папку
func (image *ImageModel) saveTo(path string) (string, error) {
	// var export *vips.ExportParams
	// switch image.Type {
	// case 3:
	// 	export = vips.NewPngExportParams()
	// case 1:
	// 	export = vips.NewJpegExportParams()
	// case 2:
	// 	export = vips.NewWebpExportParams()
	// default:
	// 	export = vips.NewDefaultExportParams()
	// }

	outBytes, _, errE := image.Ref.ExportPng(vips.NewPngExportParams())
	if errE != nil {
		return "", errE
	}

	saveFilePath := "public/" + path
	saveFileName := image.Name + image.Type.FileExt()

	_, errS := os.Stat(saveFilePath)
	if errS != nil {
		if err := os.MkdirAll(saveFilePath, 0o700); err != nil {
			return "", err
		}
	}

	file, err := os.Create(saveFilePath + "/" + saveFileName)
	if err != nil {
		return "", err
	}

	// file, err := os.OpenFile(saveFilePath+"/"+saveFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModeAppend.Perm())
	// if err != nil {
	// 	return "", err
	// }
	if err := ioutil.WriteFile(saveFilePath+"/"+saveFileName, outBytes, 0o700); err != nil {
		return "", err
	}
	if err := file.Close(); err != nil {
		return "", err
	}
	return saveFilePath, nil
}
