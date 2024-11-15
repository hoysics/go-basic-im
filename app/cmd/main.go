package main

import (
	"fmt"
	"github.com/hoysics/basic-im/app/client"
	"github.com/spf13/cobra"
	"os"
)

var (
	ConfigPath string
)

var rootCmd = &cobra.Command{
	Use:   "basic-im",
	Short: "就是个IM",
	Run:   BasicIm,
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&ConfigPath, "config", "./basic.yaml", "config file (default is ./basic.yaml)")

	rootCmd.AddCommand(&cobra.Command{
		Use: "client",
		Run: ClientHandle,
	})
}

func initConfig() {
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func BasicIm(cmd *cobra.Command, args []string) {
	fmt.Println(`just a reliable IM`)
}
func ClientHandle(cmd *cobra.Command, args []string) {
	client.RunMain()
}
