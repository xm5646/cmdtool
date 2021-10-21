// Package conf
// 功能描述: 初始化日志配置
// Date: 2021/10/21
// author: lixiaoming
package conf

import "github.com/xm5646/log"

var CMDName = "matcher"

func init() {
	logInit()
}

func logInit() {
	cf := log.DefaultLagerDefinition()
	cf.LogFormatText = true
	cf.Writers = "file,stdout"
	cf.LoggerFile = "log/eventMatcher.log"
	cf.RollingPolicy = "size"
	cf.LogRotateSize = 20
	cf.LogBackupCount = 10
	cf.LoggerLevel = "DEBUG"
	log.InitWithConfig(cf)
}
