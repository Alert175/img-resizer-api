package domain

import "github.com/davidbyttow/govips/v2/vips"

type ImageModel struct {
	Buffer []byte         // буфер данных изображения
	Ref    *vips.ImageRef // тип данных необходимый для работы с libvips
	Name   string
	Type   vips.ImageType // тип изображения
	Height int
	Width  int
}

type ImageUseCase interface {
	InstallFromNetwork(url string) (string, error)
}
