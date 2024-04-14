package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

func ViperInit() {
	viper.SetConfigName("application")    //读取配置文件名
	viper.SetConfigType("yml")            //配置文件类型,用来远程etcd获取配置信息格式
	viper.AddConfigPath("./pkg/configs/") //查找路径

	//配置读取并加载
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			//文件未找到错误提示
			fmt.Println("viper.ReadInConfig failed. err :", err)
		} else {
			//文件找到,加载错误提示
			panic(fmt.Errorf("fatal error config file: %s", err))
		}
	}
}
