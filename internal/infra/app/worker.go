package app

import "github.com/urfave/cli"

func (a *ApplicationContext) Consume() cli.Command {
	return cli.Command{
		Name:  "consume",
		Usage: "Start Service",
		Action: func(c *cli.Context) error {
			a.poolProducer.Subscribe(a.poolConsumer.Consumers())
			a.poolConsumer.Start()
			return nil
		},
	}
}
