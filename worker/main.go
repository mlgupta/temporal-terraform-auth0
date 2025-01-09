package main

import (
	"context"
	"log"
	"sync"

	"github.com/dbsentry/temporal-terraform-auth0/activities"
	"github.com/dbsentry/temporal-terraform-auth0/logger"
	"github.com/dbsentry/temporal-terraform-auth0/workflows"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Logger logger.Logger
	Client client.Client
}

func opts() fx.Option {
	return fx.Options(
		logger.Module,
		fx.Provide(
			NewWorker,
		),
		fx.Invoke(
			run,
		),
	)

}
func main() {
	fx.New(opts()).Run()
}

func run(p Params) {
	p.Logger.Info("Starting worker")
}

type WorkerParams struct {
	fx.In

	Lifecycle fx.Lifecycle

	logger.Logger
}

func NewWorker(p WorkerParams) (client.Client, error) {
	p.Logger.Info("Starting worker in NewWorker")

	// Create the client object just once per process
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}

	var wg sync.WaitGroup

	p.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			wg.Add(1)
			go func() {
				defer wg.Done()
				w := worker.New(c, workflows.TerraformTaskQueue, worker.Options{})
				w.RegisterWorkflow(workflows.TerraformWorkflow)
				w.RegisterActivity(activities.TerraformInitAuth0Activity)
				w.RegisterActivity(activities.TerraformApplyAuth0Activity)
				w.RegisterActivity(activities.TerraformOutputAuth0Activity)

				err := w.Run(worker.InterruptCh())
				if err != nil {
					log.Fatalln("failed to run worker here", err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			wg.Wait()
			p.Logger.Info("Stopping worker")
			c.Close()
			return nil
		},
	})

	return c, nil

}
