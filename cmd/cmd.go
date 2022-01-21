package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"github.com/xylonx/icc-core/internal/config"
	"github.com/xylonx/icc-core/internal/core"
	"github.com/xylonx/icc-core/internal/service"
)

var rootCmd = &cobra.Command{
	Use:   "icc-core",
	Short: "Image Collection Center - core service",
	PreRunE: func(cmd *cobra.Command, args []string) (err error) {
		err = config.Setup(cfgFile)
		if err != nil {
			return err
		}

		err = core.Setup()
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
	if err := service.StartService(); err != nil {
		return err
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM)

	<-sig

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	service.StopService(ctx)

	return nil
}
