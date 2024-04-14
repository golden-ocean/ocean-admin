package configs

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golden-ocean/ocean-admin/pkg/exceptions"
)

func FiberConfig() fiber.Config {

	return fiber.Config{
		ErrorHandler:              exceptions.Handler,
		AppName:                   "ocean-admin",     // APP名称
		BodyLimit:                 4 * 1024 * 1024,   // 请求体大小
		CaseSensitive:             false,             // 路由区分大小写
		CompressedFileSuffix:      ".ocean-admin.gz", // 压缩文件后缀
		Concurrency:               256 * 1024,        // 并发数
		DisableDefaultContentType: false,             // 排除默认响应头
		DisableStartupMessage:     false,             // 关闭调试信息
		ETag:                      false,             // Etag标头启用
		EnablePrintRoutes:         false,             // 打印路由
		Immutable:                 false,             // 上下文不可变
		//JSONDecoder:               sonic.Unmarshal,                 // json 解码
		//JSONEncoder:               sonic.Marshal,                   // json 编码
		Prefork:       false,                           // 同一端口多进程
		ServerHeader:  "",                              // 服务报头
		StrictRouting: false,                           // 严格路由(是否区分 /foo 和 /foo/
		UnescapePath:  false,                           // 转换路由编码字符
		ReadTimeout:   time.Second * time.Duration(60), // 读取请求允许的时间
	}
}
