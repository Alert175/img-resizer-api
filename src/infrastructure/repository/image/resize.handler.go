package image

import (
	"strconv"

	"github.com/davidbyttow/govips/v2/vips"
)

func (image *ImageModel) resize(width int, height int) error {
	validWidth := width
	validHeight := height
	image.Name = ""
	if width != 0 {
		image.Name += "w" + strconv.Itoa(width)
	}
	if height != 0 {
		image.Name += "h" + strconv.Itoa(height)
	}
	if width > image.Width || height > image.Height {
		validWidth = image.Width
		validHeight = image.Height
	}
	if validWidth == 0 {
		validWidth = image.Width / int(image.Height/height)
	}
	if validHeight == 0 {
		validHeight = image.Height / int(image.Width/width)
	}
	if err := image.Ref.Thumbnail(validWidth, validHeight, vips.InterestingEntropy); err != nil {
		return err
	}
	return nil
}
