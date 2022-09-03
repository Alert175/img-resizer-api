package image

import (
	"crypto/tls"
	"io/ioutil"
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

	imageRef, err := vips.LoadImageFromBuffer(body, nil)
	if err != nil {
		return err
	}

	image.Ref = imageRef
	image.Buffer = body
	image.Width = imageRef.Width()
	image.Height = imageRef.Height()
	image.Type = vips.DetermineImageType(body)
	image.Name = "origin"

	return nil
}
