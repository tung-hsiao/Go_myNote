package main

import (
    "log"

    "github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

var httpApplication *fiber.App
const (bodyLimit = 100 * 1024 * 1024) // 上傳大小限制100MB

func main() {
    // Initialize a new Fiber app
    httpApplication = fiber.New(fiber.Config{BodyLimit: bodyLimit})
    httpApplication.Use(compress.New()) // 啟用壓縮
    httpApplication.Use(cors.New())   // 啟用CORS
    httpApplication.Use(logger.New()) // 啟用日誌

    addRoute_HelloWorld()
    addRoute_TungTest()

    // Start the server on port 3000
    log.Fatal(httpApplication.Listen(":3000"))
}

func addRoute_HelloWorld(){
    // Define a route for the GET method on the root path '/'
    httpApplication.Get("/", func(c fiber.Ctx) error {
        return c.SendString("Hello, World 👋!") // Send a string response to the client
    })
}

func addRoute_TungTest(){
    // Define a route for the GET method on the root path '/'
    httpApplication.Get("/tung", func(c fiber.Ctx) error {
        return c.SendString("Tung Test!") // Send a string response to the client
    })
}
