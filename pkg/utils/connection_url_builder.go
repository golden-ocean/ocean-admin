package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

func ConnectionURLBuilder(n string) (string, error) {
	var url string

	switch n {

	case "mysql":
		HOST := viper.GetString("datasource.host")
		PORT := viper.GetString("datasource.port")
		USERNAME := viper.GetString("datasource.username")
		PASSWORD := viper.GetString("datasource.password")
		CHARSET := viper.GetString("datasource.charset")
		NAME := viper.GetString("datasource.name")
		url = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
			USERNAME, PASSWORD, HOST, PORT, NAME, CHARSET,
		)

	case "redis":
		HOST := viper.GetString("redis.host")
		PORT := viper.GetString("redis.port")
		USERNAME := viper.GetString("redis.username")
		PASSWORD := viper.GetString("redis.password")
		DB := viper.GetString("redis.db")
		// url = fmt.Sprintf("%s:%s", HOST, PORT)
		// URL:        "redis://<user>:<pass>@127.0.0.1:6379/<db>",
		url = fmt.Sprintf("redis://%s:%s@%s:%s/%s", USERNAME, PASSWORD, HOST, PORT, DB)

	case "fiber":
		HOST := viper.GetString("server.host")
		PORT := viper.GetString("server.port")
		url = fmt.Sprintf("%s:%s", HOST, PORT)

	default:
		return "", fmt.Errorf("'%v' 数据库类型不支持", n)
	}

	return url, nil
}
