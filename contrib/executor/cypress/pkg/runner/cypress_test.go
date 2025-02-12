package runner

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/kubeshop/testkube/pkg/envs"

	cp "github.com/otiai10/copy"
	"github.com/stretchr/testify/assert"

	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
)

func TestRun(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	// setup
	tempDir, err := os.MkdirTemp("", "*")
	assert.NoErrorf(t, err, "failed to create temp dir: %v", err)
	defer os.RemoveAll(tempDir)
	repoDir := filepath.Join(tempDir, "repo")
	assert.NoError(t, os.Mkdir(repoDir, 0755))
	_ = cp.Copy("../../examples", repoDir)

	params := envs.Params{DataDir: tempDir}
	runner, err := NewCypressRunner(ctx, "npm", params)
	if err != nil {
		t.Fail()
	}

	repoURI := "https://github.com/kubeshop/testkube-executor-cypress.git"
	result, err := runner.Run(
		ctx,
		testkube.Execution{
			Content: &testkube.TestContent{
				Type_: string(testkube.TestContentTypeGitDir),
				Repository: &testkube.Repository{
					Type_:  "git",
					Uri:    repoURI,
					Branch: "jacek/feature/json-output",
					Path:   "",
				},
			},
		})

	assert.NoErrorf(t, err, "Cypress Test Failed: ResultErr: %v, Err: %v ", result.ErrorMessage, err)
	fmt.Printf("RESULT: %+v\n", result)
}

func TestRunErrors(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	t.Run("no RUNNER_DATADIR", func(t *testing.T) {
		t.Parallel()

		params := envs.Params{DataDir: "/unknown"}
		runner, err := NewCypressRunner(ctx, "yarn", params)
		if err != nil {
			t.Fail()
		}

		execution := testkube.NewQueuedExecution()

		// when
		_, err = runner.Run(ctx, *execution)

		// then
		assert.Error(t, err)
	})
}
