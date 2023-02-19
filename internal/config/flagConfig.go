package config

import "flag"

type flagConfig struct {
	ConfigDir  string
	BasicDir   string
	BatchModel bool
}

var FlagConfig flagConfig

func initFlagConfig() {
	flag.StringVar(&FlagConfig.ConfigDir, "c", "", "用于匹配的config文件名")
	flag.StringVar(&FlagConfig.BasicDir, "d", "", "匹配的Dir")
	flag.BoolVar(&FlagConfig.BatchModel, "b", false, "批量处理Dir")
	flag.Parse()
}
