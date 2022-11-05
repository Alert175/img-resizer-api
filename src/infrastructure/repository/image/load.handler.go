package image

import (
	"bytes"
	"crypto/tls"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"

	"github.com/davidbyttow/govips/v2/vips"
)

// Загрузить изображение по http и установить необходимые данные
func (image *ImageModel) loadFormNet(url string) error {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	response, err := client.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	image.Buffer = body

	return nil
}

func (image *ImageModel) loadToVips() error {
	imageRef, err := vips.LoadImageFromBuffer(image.Buffer, nil)
	if err != nil {
		return err
	}
	image.Ref = imageRef
	image.Width = imageRef.Width()
	image.Height = imageRef.Height()
	image.Type = vips.DetermineImageType(image.Buffer)
	image.Name = "origin"
	return nil
}

func (image *ImageModel) loadFileToImage(file *multipart.FileHeader) error {
	fileContent, err := file.Open()
	if err != nil {
		return err
	}
	defer fileContent.Close()
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, fileContent); err != nil {
		return err
	}
	image.Buffer = buf.Bytes()

	imageRef, err := vips.LoadImageFromBuffer(image.Buffer, nil)
	if err != nil {
		return err
	}
	image.Ref = imageRef
	image.Width = imageRef.Width()
	image.Height = imageRef.Height()
	image.Type = vips.DetermineImageType(image.Buffer)
	image.Name = file.Filename
	return nil
}
