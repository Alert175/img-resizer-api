package image

import "github.com/davidbyttow/govips/v2/vips"

// конвертировать изображение в необходимый формат
func (image *ImageModel) convertTo(format string) error {
	switch format {
	case "webp":
		bytes, _, err := image.Ref.ExportWebp(vips.NewWebpExportParams())
		if err != nil {
			return err
		}
		image.Buffer = bytes
		image.Type = vips.ImageTypeWEBP
	case "avif":
		bytes, _, err := image.Ref.ExportAvif(vips.NewAvifExportParams())
		if err != nil {
			return err
		}
		image.Buffer = bytes
		image.Type = vips.ImageTypeAVIF
	case "jpeg":
		export := vips.NewJpegExportParams()
		export.Quality = 100
		bytes, _, err := image.Ref.ExportJpeg(export)
		if err != nil {
			return err
		}
		image.Buffer = bytes
		image.Type = vips.ImageTypeJPEG
	case "jpg":
		bytes, _, err := image.Ref.ExportJpeg(vips.NewJpegExportParams())
		if err != nil {
			return err
		}
		image.Buffer = bytes
		image.Type = vips.ImageTypeJPEG
	case "png":
		bytes, _, err := image.Ref.ExportPng(vips.NewPngExportParams())
		if err != nil {
			return err
		}
		image.Buffer = bytes
		image.Type = vips.ImageTypePNG
	}

	return nil
}
