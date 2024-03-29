package router

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

var httpApplication *fiber.App
const (bodyLimit = 100 * 1024 * 1024) // 上傳大小限制100MB
var Prefork = false // 多線程
var localServer = ""

var BackendPort string = "3002"

// 監聽網頁服務
func Run() {
	httpApplication = fiber.New(fiber.Config{Prefork: Prefork, 
		                                     BodyLimit: bodyLimit, 
											 UnescapePath: true, 
											 ErrorHandler: notFoundHandler})

	setupRoute() // 設定路由

	httpApplication.Listen(":" + BackendPort)
}

// 未找到路徑轉到首頁
func notFoundHandler(context *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	fmt.Println("notFoundHandler, code", code, "Error", err.Error())
	return context.Status(code).SendString(err.Error())
}
func setupRoute() {
	httpApplication.Use(compress.New()) // 啟用壓縮
	// httpApplication.Use(recover.New())  // 啟用錯誤處理
	httpApplication.Use(cors.New())   // 啟用CORS
	httpApplication.Use(logger.New()) // 啟用日誌

	// API 路徑
	apiGroup := httpApplication.Group("")

	apiGroup.Get("/", test)
}


// 檢查伺服器狀況
func test(context *fiber.Ctx) error {
	return context.JSON(fiber.Map{"message": "Server is alive!"})
}
