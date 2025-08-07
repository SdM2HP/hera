package generate

import (
	"github.com/spf13/cobra"

	"src/plugins/command/generate/model"
)

var CmdGenerate = &cobra.Command{
	Use:   "generate",
	Short: "自动生成工具",
}

func init() {
	CmdGenerate.AddCommand(model.CmdModel)
}
