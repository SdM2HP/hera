package main

import (
	"context"

	"src/engine"
	"src/internal/config"
	"src/internal/server"
	"src/plugins/command"

	"github.com/spf13/cobra"
)

const version = "v0.0.1"

var (
	filename string
)

var rootCmd = &cobra.Command{
	Use:     "hera",
	Version: version,
	Run:     run,
}

func init() {
	// 命令行处理
	rootCmd.Flags().StringVarP(&filename, "config", "c", "./config/config.yaml", "配置文件路径")

	// 添加命令
	rootCmd.AddCommand(command.CmdTools)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func run(cmd *cobra.Command, args []string) {
	// 初始化配置
	if err := config.Conf.Setup(filename); err != nil {
		panic(err)
	}

	// 实例化应用
	app := engine.New(
		engine.WithContext(context.Background()),
		engine.WithServer(server.NewServer(config.Conf.Server.Http)),
	)

	// 启动应用
	if err := app.Run(); err != nil {
		panic(err)
	}
}
