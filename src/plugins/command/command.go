package command

import (
	"github.com/spf13/cobra"

	"src/plugins/command/generate"
	"src/plugins/command/upgrade"
)

var CmdTools = &cobra.Command{
	Use:   "tools",
	Short: "工具箱",
}

func init() {
	CmdTools.AddCommand(upgrade.CmdUpgrade)
	CmdTools.AddCommand(generate.CmdGenerate)
}
