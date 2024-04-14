package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golden-ocean/ocean-admin/pkg/common/global"
	"github.com/golden-ocean/ocean-admin/pkg/configs"
	"github.com/golden-ocean/ocean-admin/pkg/middlewares"
	"github.com/golden-ocean/ocean-admin/pkg/routes"
	"github.com/golden-ocean/ocean-admin/pkg/utils"
	"github.com/golden-ocean/ocean-admin/platform/database"
	"github.com/spf13/viper"
)

func main() {
	// 配置初始化
	configs.Init()
	// 连接数据库
	global.DB = database.OpenDBConnection()
	// global.Storage = cache.OpenRedisConnection()
	// fiber 自身配置
	fiberConfig := configs.FiberConfig()
	// 创建实例
	app := fiber.New(fiberConfig)
	// 中间件
	middlewares.FiberMiddleware(app) // 注册 fiber内置中间件
	// 路由.
	//routes.SwaggerRoute(app)
	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)
	// routes.NotFoundRoute(app)

	// Start server (with or without graceful shutdown).
	if viper.GetString("server.status") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
