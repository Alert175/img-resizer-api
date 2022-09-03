package imagerouter

import (
	"img-resizer-api/src/infrastructure/pkg/utils"
	imageRepo "img-resizer-api/src/infrastructure/repository/image"

	"github.com/gofiber/fiber/v2"
)

func LoadFromNet(ctx *fiber.Ctx) error {
	image := imageRepo.ImageModel{}

	body := new(LoadFromNetDto)
	if err := ctx.BodyParser(body); err != nil {
		return err
	}
	if errV := utils.Validate(body); errV != nil {
		return ctx.Status(400).JSON(errV)
	}

	result, err := image.InstallFromNetwork(body.Url)
	if err != nil {
		return ctx.Status(500).JSON(err)
	}
	return ctx.JSON(result)
}
