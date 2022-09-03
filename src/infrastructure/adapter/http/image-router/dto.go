package imagerouter

type LoadFromNetDto struct {
	Url string `json:"url" validate:"required"`
}

type LoadFromNetDtoAndResize struct {
	LoadFromNetDto
	Height int `json:"height"`
	Width  int `json:"width"`
}

type LoadFromNetDtoAndResizeAndConvert struct {
	LoadFromNetDtoAndResize
	Format string `json:"format" validate:"required"`
}
