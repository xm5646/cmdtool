// Package command
// 功能描述:
// Date: 2021/10/21
// author: lixiaoming
package cmd

import (
	"event-matcher/conf"
	"fmt"
	"github.com/urfave/cli/v2"
)

var TestCMD = &cli.Command{
	Name:    "test",
	Aliases: []string{"test"},
	Usage:   fmt.Sprintf("./%s test", conf.CMDName),
	Action: func(c *cli.Context) error {
		fmt.Println("cmd test")
		return nil
	},
}
