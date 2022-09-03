package imagerouter

import imageRepo "img-resizer-api/src/infrastructure/repository/image"

type LoadFromNetDto struct {
	Url string `json:"url" validate:"required"`
}

type LoadFromNetDtoAndResize struct {
	LoadFromNetDto
	imageRepo.ResizeDto
}

type LoadFromNetDtoAndResizeAndConvert struct {
	LoadFromNetDtoAndResize
	imageRepo.FormatDto
}

type OptimizeDto struct {
	LoadFromNetDto
	Points []imageRepo.Point `json:"points" validate:"required"`
}
