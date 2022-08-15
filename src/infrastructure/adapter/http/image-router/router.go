package imagerouter

import "github.com/gofiber/fiber/v2"

func Register(httpPath fiber.Router) {
	router := httpPath.Group("/image")

	router.Post("/load-from-net", LoadFromNet)
}
