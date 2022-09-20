package imagerouter

import (
	"os"
	"strconv"
	"strings"

	"img-resizer-api/src/infrastructure/pkg/utils"
	imageRepo "img-resizer-api/src/infrastructure/repository/image"

	"github.com/gofiber/fiber/v2"
)

// Create godoc
// @Summary         Загрузить изображение
// @Description   	Загрузить изображение и сохранить на директории сервера
// @Tags            Image
// @Accept          json
// @Produce       	json
// @Param data body LoadFromNetDto false "-"
// @Success       	200  {string}  string    "image url"
// @Failure         400  {string}  string    "error"
// @Failure         404  {string}  string    "error"
// @Failure         500  {string}  string    "error"
// @Router         /api/image/v2/image [post]
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

// Create godoc
// @Summary         Изменить размер изображения
// @Description   	Загрузить изображение и сохранить на директории сервера, сделать ресайз
// @Tags            Image
// @Accept          json
// @Produce       	json
// @Param data body LoadFromNetDtoAndResize false "-"
// @Success       	200  {string}  string    "image url"
// @Failure         400  {string}  string    "error"
// @Failure         404  {string}  string    "error"
// @Failure         500  {string}  string    "error"
// @Router         /api/image/v2/image/resize [post]
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

// Create godoc
// @Summary         Конвертировать изображение
// @Description   	Загрузить изображение и сохранить на директории сервера, сделать ресайз, и перевести в нужный формат
// @Tags            Image
// @Accept          json
// @Produce       	json
// @Param data body LoadFromNetDtoAndResizeAndConvert false "-"
// @Success       	200  {string}  string    "image url"
// @Failure         400  {string}  string    "error"
// @Failure         404  {string}  string    "error"
// @Failure         500  {string}  string    "error"
// @Router         /api/image/v2/image/resize/convert [post]
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

// Create godoc
// @Summary         Оптимизировать изображение
// @Description   	Загрузить изображение и сохранить на директории сервера, сделать ресайз, и перевести в нужный формат
// @Tags            Image
// @Accept          json
// @Produce       	json
// @Param data body OptimizeDto false "-"
// @Success       	200  {string}  string    "image urls"
// @Failure         400  {string}  string    "error"
// @Failure         404  {string}  string    "error"
// @Failure         500  {string}  string    "error"
// @Router         /api/image/v2/image/optimize [post]
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

// Create godoc
// @Summary         Оптимизировать изображение
// @Description   	Загрузить изображение и сохранить на директории сервера, сделать ресайз, и перевести в нужный формат
// @Tags            Image
// @Accept          json
// @Produce       	json
// @Param data body OptimizeDto false "-"
// @Success       	200  {string}  string    "image urls"
// @Failure         400  {string}  string    "error"
// @Failure         404  {string}  string    "error"
// @Failure         500  {string}  string    "error"
// @Router         /api/image/v2/image/optimize [post]
func GetFormNet(ctx *fiber.Ctx) error {
	image := imageRepo.ImageModel{}
	query := GetFromNet{}

	if err := ctx.QueryParser(query); err != nil {
		return ctx.Status(400).JSON(err)
	}

	h, errH := strconv.Atoi(query.Height)
	w, errW := strconv.Atoi(query.Width)

	if errH != nil || errW != nil {
		return ctx.Status(400).JSON("not valid height or width")
	}

	image.GetFromNetworkAndResizeAndConvert(query.Url, h, w, query.Format)
	ctx.Write()
}
