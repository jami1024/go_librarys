package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {
	// 初始化 viper 实例
	v := viper.New()

	// 设置配置文件类型和路径
	v.SetConfigType("yaml")
	v.SetConfigFile("config.yaml")

	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file", err)
	}

	// 读取配置项
	dbHost := v.GetString("database.host")
	dbPort := v.GetInt("database.port")

	// 打印配置项
	fmt.Println("Database Host:", dbHost)
	fmt.Println("Database Port:", dbPort)
}
