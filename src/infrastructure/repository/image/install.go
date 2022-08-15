package image

import (
	"crypto/tls"
	"img-resizer-api/src/infrastructure/pkg/utils/logger"
	"io/ioutil"
	"net/http"

	"github.com/davidbyttow/govips/v2/vips"
)

func InstallNet(url string) error {
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

	// defer vips.Shutdown()

	imageRef, err := vips.LoadImageFromBuffer(body, nil)
	if err != nil {
		return err
	}

	width := imageRef.Width()
	height := imageRef.Height()
	imageType := vips.DetermineImageType(body)

	logger.Log(width)
	logger.Log(height)
	logger.Log(imageType)

	return nil
}
