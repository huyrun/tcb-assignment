package main

import (
	"context"
	"log"
	"os"

	"tcb-assignment/internal/infra/app"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	ctx := context.Background()
	cli, cleanup, err := app.InitApplication(ctx)
	if err != nil {
		panic(err)
	}
	app.HandleSigterm(cleanup)

	err = cli.Commands().Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
