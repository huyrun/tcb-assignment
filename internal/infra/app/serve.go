package app

import (
	tcb_assignment "tcb-assignment"

	"github.com/urfave/cli"
)

// Serve start application
func (a *ApplicationContext) Serve() cli.Command {
	return cli.Command{
		Name:  "serve",
		Usage: "Start Service",
		Action: func(c *cli.Context) error {
			tcb_assignment.VisualizeRbtreeMode = a.cfg.Visualize
			a.poolProducer.Subscribe(a.poolConsumer.Consumers())
			a.poolConsumer.Start()
			a.restSrv.MustStart()
			return nil
		},
	}
}
