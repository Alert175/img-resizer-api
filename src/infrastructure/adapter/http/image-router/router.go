package imagerouter

import "github.com/gofiber/fiber/v2"

func Register(httpPath fiber.Router) {
	router := httpPath.Group("/image")

	router.Post("/", LoadFromNet)
	router.Post("/resize", LoadFromNetResize)
	router.Post("/resize/convert", LoadFromNetResizeConvert)
	router.Post("/optimize", Optimize)
	router.Post("/get", GetFormNet)
}
