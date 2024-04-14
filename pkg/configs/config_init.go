package configs

func Init() {
	// 读取配置文件
	ViperInit()
	// 初始化redis
	// 注册验证器
	ValidatorInit("zh")
	// 初始化日志

}
