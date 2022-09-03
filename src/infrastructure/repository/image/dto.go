package image

type Point struct {
	ResizeDto
	FormatDto
}

type FormatDto struct {
	Format string `json:"format" validate:"required"`
}
type ResizeDto struct {
	Height int `json:"height"`
	Width  int `json:"width"`
}
