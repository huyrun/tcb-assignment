package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"tcb-assignment/internal/consumers"
	"tcb-assignment/internal/infra"
	"tcb-assignment/internal/pkg/message-queue/producer"

	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

type ApplicationContext struct {
	ctx          context.Context
	cfg          *infra.AppConfig
	restSrv      *infra.RestService
	poolProducer producer.Producer
	poolConsumer consumers.PoolConsumer
}

var ApplicationSet = wire.NewSet(
	infra.ProvideConfig,
	infra.ProvideRestAPIHandler,
	infra.ProvideRestService,
	infra.ProvideAuthService,
	infra.ProvideCacheService,
	infra.ProvidePoolService,
	infra.ProvidePoolRepo,
	infra.ProvidePoolProducer,
	infra.ProvidePoolConsumer,
)

func (a *ApplicationContext) Commands() *cli.App {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		a.Serve(),
		a.Consume(),
	}

	return app
}

// HandleSigterm -- Handles Ctrl+C or most other means of "controlled" shutdown gracefully.
// Invokes the supplied func before exiting.
func HandleSigterm(handleExit func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGTRAP)
	go func() {
		<-c
		logrus.Infof("Handler shutdown signal in main process...")
		handleExit()
		os.Exit(1)
	}()
}
