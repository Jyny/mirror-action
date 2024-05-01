package main

import (
	"os"

	"github.com/caarlos0/env"
	"github.com/jyny/mirror-action/pkg/config"
	"github.com/jyny/mirror-action/pkg/logger"
	"github.com/jyny/mirror-action/pkg/mirror"
)

var (
	appLogger logger.Logger
	appConfig config.Config
)

func init() {
	appLogger = logger.New(os.Stdout, logger.InfoLevel)

	if err := env.Parse(&appConfig); err != nil {
		appLogger.Fatal("Error env.Parse():", "err", err)
	}

	if appConfig.Debug {
		appLogger.SetLevel(logger.DebugLevel)
	}
}

func main() {
	app, err := mirror.New(
		mirror.MirrorConfig{
			RemoteURL:     appConfig.SrcRemoteURL,
			SSHKey:        appConfig.SrcSShKey,
			HostKey:       appConfig.SrcKnownHost,
			IgnoreHostKey: appConfig.SrcIgnoreHostKey,
			Username:      appConfig.SrcUsername,
			Password:      appConfig.SrcPassword,
		},
		mirror.MirrorConfig{
			RemoteURL:     appConfig.DstRemoteURL,
			SSHKey:        appConfig.DstSShKey,
			HostKey:       appConfig.DstKnownHost,
			IgnoreHostKey: appConfig.DstIgnoreHostKey,
			Username:      appConfig.DstUsername,
			Password:      appConfig.DstPassword,
		},
		appConfig.RefSpec,
		appLogger,
	)
	if err != nil {
		appLogger.Fatal("Error mirror.New():", "err", err)
	}

	if err := app.Run(); err != nil {
		appLogger.Fatal("Error app.Run():", "err", err)
	}
}
