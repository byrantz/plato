package cmd

import (
	"github.com/byrantz/plato/client"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(clientCmd)
}

var clientCmd = &cobra.Command{
	Use: "client",
	Run: ClientHandle,
}

func ClientHandle(cmd *cobra.Command, args []string) {
	// 运行 "github.com/byrantz/plato/client" 中的 RunMain 函数
	client.RunMain()
}
