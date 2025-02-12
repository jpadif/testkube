package main

import (
	"context"
	"os"

	"github.com/kubeshop/testkube/pkg/envs"

	"github.com/pkg/errors"

	"github.com/kubeshop/testkube/contrib/executor/scraper/pkg/runner"
	"github.com/kubeshop/testkube/pkg/executor/agent"
	"github.com/kubeshop/testkube/pkg/executor/output"
)

func main() {
	ctx := context.Background()
	params, err := envs.LoadTestkubeVariables()
	if err != nil {
		output.PrintError(os.Stderr, errors.Errorf("could not initialize Scraper Runner environment variables: %v", err))
		os.Exit(1)
	}
	r, err := runner.NewRunner(ctx, params)
	if err != nil {
		output.PrintError(os.Stderr, errors.Wrap(err, "could not initialize runner"))
		os.Exit(1)
	}
	agent.Run(ctx, r, os.Args)
}
