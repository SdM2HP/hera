package model

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"src/plugins/database"
	"src/plugins/database/gen"
)

var (
	dsn     string
	server  string
	outpath string

	CmdModel = &cobra.Command{
		Use:   "model",
		Short: "生成数据库模型",
		Run:   run,
	}
)

func init() {
	// 命令行处理
	CmdModel.Flags().StringVarP(&dsn, "dsn", "d", "root:123456@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local", "connect to db")
	CmdModel.Flags().StringVarP(&server, "server", "s", "", "server name")
	CmdModel.Flags().StringVarP(&outpath, "outpath", "o", "./data/query", "output path")
}

func run(_ *cobra.Command, args []string) {
	db := database.New(database.WithDSN(dsn))
	outpathpath := outpath
	if server != "" {
		outpathpath = fmt.Sprintf("./data/%s/query", strings.ToLower(server))
	}
	generator := gen.New(db, gen.WithOutPath(outpathpath))
	generator.Execute()
}
