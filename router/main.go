package main

import (
    "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var httpApplication *fiber.App

const (
	bodyLimit = 100 * 1024 * 1024 // 上傳大小限制100MB
)

var Prefork = false // 多線程
var localServer = ""

// 監聽網頁服務
func main() {
	httpApplication = fiber.New(fiber.Config{Prefork: Prefork, BodyLimit: bodyLimit, UnescapePath: true})
	setupRoute()
    httpApplication.Listen(":3000")
}

// 設定路由
func setupRoute() {
	httpApplication.Use(compress.New()) // 啟用壓縮
	httpApplication.Use(cors.New())   // 啟用CORS
	httpApplication.Use(logger.New()) // 啟用日誌

    // API 路徑
	apiGroup := httpApplication.Group("")

    apiGroup.Get("/", helloWorld)
    apiGroup.Get("/hello/:name", helloName)

}

func helloWorld(context *fiber.Ctx) error {
    return context.Status(200).JSON(fiber.Map{"Message": "Hello Word"})
}

// Define a route with dynamic parameter
func helloName(context *fiber.Ctx) error {
    name := context.Params("name")
    return context.SendString("Hello, " + name + "!")
}
