package imagerouter

import (
	"img-resizer-api/src/domain"
	"img-resizer-api/src/infrastructure/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func LoadFromNet(ctx *fiber.Ctx) error {
	imageDomain := domain.ImageModel{}

	var body = new(LoadFromNetDto)
	if err := ctx.BodyParser(body); err != nil {
		return err
	}
	if errV := utils.Validate(body); errV != nil {
		return ctx.Status(400).JSON(errV)
	}

	imageDomain.InstallFromNetwork(body.Url)
	return ctx.JSON("ok")
}
