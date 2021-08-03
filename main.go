package main

import (
	"fmt"
	"github.com/milkb32/go-command/cmds"
	"github.com/milkb32/go-command/commands"
	"github.com/milkb32/go-command/config"
	"github.com/urfave/cli"
	"os"
	"time"
)

func main() {
	err := config.Init()
	if err != nil {
		fmt.Println("配置文件初始化出错:", err)
		return
	}

	start := time.Now()

	// 初始化app
	app := cli.NewApp()
	app.Name = "go-command"
	app.Usage = "用来执行各种命令"
	app.Description = "是一个命令行工具机，集成了所有的命令"

	// 命令注册
	commands.Register()

	app.Commands = cmds.GetCommands()

	err = app.Run(os.Args)
	if err != nil {
		fmt.Println("配置文件初始化出错:", err)
		return
	}

	fmt.Printf("commands: %v \r\nsuccess use %s\n", os.Args[1], time.Now().Sub(start))
}
