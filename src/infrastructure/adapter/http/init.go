package http

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	_ "img-resizer-api/docs"
	imagerouter "img-resizer-api/src/infrastructure/adapter/http/image-router"

	"github.com/gofiber/swagger"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func InitHttpAdapter() {
	app := fiber.New()

	pathPrefix := os.Getenv("HTTP_PATH_PREFIX")
	port := os.Getenv("HTTP_APP_PORT")

	if len(pathPrefix) == 0 || len(port) == 0 {
		log.Fatal("http env not found")
	}

	configLogger(app)
	configCors(app)

	appPath := app.Group(pathPrefix)
	appPath.Get("/swagger/*", swagger.HandlerDefault)
	appPath.Static("/", "./public")

	imagerouter.Register(appPath)

	app.Listen(":" + port)
}

// настройка логгера
func configLogger(app *fiber.App) {
	file, err := os.OpenFile("./logs/http.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
	if err != nil {
		panic(fmt.Sprintf("error opening file: %v", err))
	}

	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "[\n request_id ${locals:requestid}\n status ${status}\n method ${method}\n path ${path}\n time ${time}\n latency ${latency}\n reqHeaders ${reqHeaders}\n params ${queryParams}\n query ${query}\n body ${body}\n]\n",
		Output: file,
	}))
}

func configCors(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, HEAD, PUT, DELETE, PATCH, OPTION",
	}))
}
