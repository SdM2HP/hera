package command

import (
	"github.com/spf13/cobra"

	"src/plugins/command/generate"
	"src/plugins/command/upgrade"
)

var CmdTool = &cobra.Command{
	Use:   "tool",
	Short: "工具箱",
}

func init() {
	CmdTool.AddCommand(upgrade.CmdUpgrade)
	CmdTool.AddCommand(generate.CmdGenerate)
}
