//功能描述: 事件丰富策略匹配器
//Date: 2021/10/21
//author: lixiaoming
package main

import (
	"event-matcher/cmd"
	_ "event-matcher/conf"
	"github.com/urfave/cli/v2"
	"github.com/xm5646/log"
	"os"
)

func main() {
	authors := make([]*cli.Author, 0, 1)
	myself := &cli.Author{
		Name:  "李晓明",
		Email: "lixmsucc@163.com",
	}
	authors = append(authors, myself)
	app := &cli.App{
		Name:        "matcher",
		Usage:       "事件丰富匹配器",
		Authors:     authors,
		Version:     "1.0",
		Description: "用于对事件内容进行匹配后丰富事件字段",
		Commands: []*cli.Command{
			cmd.TestCMD,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Errorf(err, "命令执行失败")
	}
}
