package commands

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/iand/logfmtr"
	"github.com/imperialhound/kli/internal/klient"
	"github.com/spf13/cobra"
)

var logLevel int

func Run(ctx context.Context, logger logr.Logger, client *klient.Klient) error {
	rootCmd, err := newRootCmd(ctx, logger, client)
	if err != nil {
		return err
	}

	err = rootCmd.Execute()
	if err != nil {
		return err
	}

	return nil
}

func newRootCmd(ctx context.Context, logger logr.Logger, client *klient.Klient) (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use:              "kli",
		Short:            "kli is used for day to day kubernetes operations and test",
		PersistentPreRun: setVerbosity,
	}

	rootCmd.PersistentFlags().IntVar(&logLevel, "log-level", 0, "verbosity of logs")

	return rootCmd, nil
}

func setVerbosity(cmd *cobra.Command, _ []string) {
	logfmtr.SetVerbosity(logLevel)
}
