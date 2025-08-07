package upgrade

import (
	"fmt"

	"github.com/spf13/cobra"

	"src/plugins/command/base"
)

var CmdUpgrade = &cobra.Command{
	Use:   "upgrade",
	Short: "升级常用二进制文件",
	Run:   run,
}

func run(_ *cobra.Command, _ []string) {
	err := base.GoInstall(
		"google.golang.org/protobuf/cmd/protoc-gen-go@latest",
		"google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest",
		"github.com/google/gnostic/cmd/protoc-gen-openapi@latest",
		"github.com/favadi/protoc-go-inject-tag@latest",
		"github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2",
	)
	if err != nil {
		fmt.Println(err)
	}
}
