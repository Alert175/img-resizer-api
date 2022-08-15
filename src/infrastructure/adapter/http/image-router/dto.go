package imagerouter

type LoadFromNetDto struct {
	Url string `json:"url" validate:"required"`
}
