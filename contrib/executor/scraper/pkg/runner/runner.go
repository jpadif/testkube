package runner

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/kubeshop/testkube/pkg/executor/scraper"

	"github.com/pkg/errors"

	"github.com/kubeshop/testkube/pkg/executor/scraper/factory"

	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"github.com/kubeshop/testkube/pkg/envs"
	"github.com/kubeshop/testkube/pkg/executor/output"
	"github.com/kubeshop/testkube/pkg/executor/runner"
)

// NewRunner creates scraper runner
func NewRunner(ctx context.Context, params envs.Params) (*ScraperRunner, error) {
	var err error
	r := &ScraperRunner{
		Params: params,
	}

	r.Scraper, err = factory.TryGetScrapper(ctx, params)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// ScraperRunner prepares data for executor
type ScraperRunner struct {
	Params  envs.Params
	Scraper scraper.Scraper
}

var _ runner.Runner = &ScraperRunner{}

// Run prepares data for executor
func (r *ScraperRunner) Run(ctx context.Context, execution testkube.Execution) (result testkube.ExecutionResult, err error) {
	// check that the artifact dir exists
	if execution.ArtifactRequest == nil {
		return *result.Err(errors.Errorf("executor only support artifact based tests")), nil
	}

	_, err = os.Stat(execution.ArtifactRequest.VolumeMountPath)
	if errors.Is(err, os.ErrNotExist) {
		return result, err
	}

	if r.Params.ScrapperEnabled {
		directories := execution.ArtifactRequest.Dirs
		if len(directories) == 0 {
			directories = []string{"."}
		}

		for i := range directories {
			directories[i] = filepath.Join(execution.ArtifactRequest.VolumeMountPath, directories[i])
		}

		output.PrintLog(fmt.Sprintf("Scraping directories: %v", directories))

		if err := r.Scraper.Scrape(ctx, directories, execution); err != nil {
			return *result.Err(err), errors.Wrap(err, "error scraping artifacts from container executor")
		}
	}

	return result, nil
}

// GetType returns runner type
func (r *ScraperRunner) GetType() runner.Type {
	return runner.TypeFin
}
