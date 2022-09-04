package image

import (
	"strconv"

	"github.com/davidbyttow/govips/v2/vips"
)

func (image *ImageModel) resize(width int, height int) error {
	var widthScale float64 = 0
	var heightScale float64 = 0
	image.Name = ""
	if width != 0 {
		widthScale = float64(width) / float64(image.Width)
		image.Name += "w" + strconv.Itoa(width)
	}
	if height != 0 {
		heightScale = float64(height) / float64(image.Height)
		image.Name += "h" + strconv.Itoa(height)
	}
	if widthScale == 0 {
		widthScale = heightScale
	}
	if heightScale == 0 {
		heightScale = widthScale
	}
	if err := image.Ref.ResizeWithVScale(heightScale, widthScale, vips.KernelAuto); err != nil {
		return err
	}
	return nil
}
