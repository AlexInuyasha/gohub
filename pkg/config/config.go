package config

import (
	viperlib "github.com/spf13/viper" // 自定义包名，避免与内置 viper 实例冲突
)

// viper 库实例
var viper *viperlib.Viper

// ConfigFunc 动态加载配置信息
type ConfigFunc func() map[string]interface{}

// ConfigFuncs 先加载到此数组，loadConfig 再动态生成配置信息
var ConfigFuncs map[string]ConfigFunc

func init() {

	// 初始化 Viper 库
	viper = viperlib.New()
	// 设置配置类型
	viper.SetConfigType("env")
	// 环境变量配置文件查找的路径， 相对于main.go
	viper.AddConfigPath(".")
	// 设置环境变量前缀，用以区分Go的系统环境变量
	viper.SetEnvPrefix("appenv")
	// 读取环境变量
	viper.AutomaticEnv()

	make(map[string]ConfigFunc)
}
