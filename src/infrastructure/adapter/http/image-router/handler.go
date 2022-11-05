package imagerouter

import (
	"errors"
	"net/http"
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
// @Summary         Загрузить и оптимизировать изображение
// @Description   	Загрузить изображение, сделать ресайз, и перевести в нужный формат
// @Tags            Image
// @Accept          json
// @Produce       	json
// @Param data body RequestLoadOptimize false "Загрузить FormData c файлом"
// @Success       	200  {string}  string    "файл"
// @Failure         400  {string}  string    "error"
// @Failure         500  {string}  string    "error"
// @Router         /api/image/v2/image/optimize/load [post]
func OptimizeLoad(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		return err
	}
	height := ctx.FormValue("height")
	width := ctx.FormValue("width")
	if height == "" {
		height = "0"
	}
	if width == "" {
		width = "0"
	}
	validHeigth, err1 := strconv.Atoi(height)
	validWidth, err2 := strconv.Atoi(width)

	format := ctx.FormValue("format")

	if err1 != nil || err2 != nil {
		return errors.New("fail validate heigth or width arguments")
	}

	image := imageRepo.ImageModel{}
	if err := image.OptimizeLoad(file, validHeigth, validWidth, format); err != nil {
		return err
	}

	ctx.Response().Header.Add("Content-Type", http.DetectContentType(image.Buffer))
	ctx.Response().Header.Add("Content-Desposition", http.DetectContentType(image.Buffer))
	return ctx.JSON("OK")
}
