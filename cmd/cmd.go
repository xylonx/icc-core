package cmd

import (
	"github.com/spf13/cobra"

	"github.com/xylonx/go-template/internal/config"
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	PreRunE: func(cmd *cobra.Command, args []string) (err error) {
		err = config.Setup(cfgFile)
		if err != nil {
			return err
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return run()
	},
}

var cfgFile string

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "config.default.yaml", "specify config file path")
}

func Execute() error {
	return rootCmd.Execute()
}

func run() error {
	return nil
}
