package main

import (
	"context"
	"os"

	"github.com/imperialhound/kli/internal/commands"
	"github.com/imperialhound/kli/internal/klient"
	"github.com/imperialhound/kli/internal/utils"
)

func main() {
	logger := utils.NewLogger()

	// Generate new kubernetes clientset
	client, err := klient.New(logger)
	if err != nil {
		logger.Error(err, "Failed to create kubernetes clientset")
		os.Exit(1)
	}

	ctx := context.Background()

	err = commands.Run(ctx, logger, client)
	if err != nil {
		logger.Error(err, "Fatal error occured")
	}
}
