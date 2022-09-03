package imagerouter

import (
	"os"
	"strings"

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

	appPath := os.Getenv("HTTP_PATH_PREFIX")
	return ctx.JSON(strings.ReplaceAll(result, "public", appPath))
}

func LoadFromNetResize(ctx *fiber.Ctx) error {
	image := imageRepo.ImageModel{}

	body := new(LoadFromNetDtoAndResize)
	if err := ctx.BodyParser(body); err != nil {
		return err
	}
	if errV := utils.Validate(body); errV != nil {
		return ctx.Status(400).JSON(errV)
	}
	if body.Height == 0 && body.Width == 0 {
		return ctx.Status(400).JSON("not valid height and width")
	}

	result, err := image.InstallFromNetworkAndResize(body.Url, body.Height, body.Width)
	if err != nil {
		return ctx.Status(500).JSON(err)
	}

	appPath := os.Getenv("HTTP_PATH_PREFIX")
	return ctx.JSON(strings.ReplaceAll(result, "public", appPath))
}

func LoadFromNetResizeConvert(ctx *fiber.Ctx) error {
	image := imageRepo.ImageModel{}

	body := new(LoadFromNetDtoAndResizeAndConvert)
	if err := ctx.BodyParser(body); err != nil {
		return err
	}
	if errV := utils.Validate(body); errV != nil {
		return ctx.Status(400).JSON(errV)
	}
	if body.Height == 0 && body.Width == 0 {
		return ctx.Status(400).JSON("not valid height and width")
	}

	result, err := image.InstallFromNetworkAndResizeAndConvert(body.Url, body.Height, body.Width, body.Format)
	if err != nil {
		return ctx.Status(500).JSON(err)
	}

	appPath := os.Getenv("HTTP_PATH_PREFIX")
	return ctx.JSON(strings.ReplaceAll(result, "public", appPath))
}

func Optimize(ctx *fiber.Ctx) error {
	image := imageRepo.ImageModel{}

	body := new(OptimizeDto)
	if err := ctx.BodyParser(body); err != nil {
		return err
	}
	if errV := utils.Validate(body); errV != nil {
		return ctx.Status(400).JSON(errV)
	}
	appPath := os.Getenv("HTTP_PATH_PREFIX")

	resultList, err := image.InstallFromNetworkAndOptimize(body.Url, body.Points)
	if err != nil {
		return ctx.Status(500).JSON(err)
	}
	validResultList := []string{}
	for _, str := range resultList {
		validResultList = append(validResultList, strings.ReplaceAll(str, "public", appPath))
	}

	return ctx.JSON(validResultList)
}
