package runner

import (
	"context"
	"os"
	"testing"

	"github.com/kubeshop/testkube/pkg/utils/test"

	"github.com/kubeshop/testkube/pkg/envs"

	"github.com/stretchr/testify/assert"

	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
)

const repoURI = "https://github.com/kubeshop/testkube-executor-ginkgo.git"

func TestRun_Integration(t *testing.T) {
	test.IntegrationTest(t)
	t.Parallel()

	ctx := context.Background()

	t.Run("GinkgoRunner should run tests from a repo that pass", func(t *testing.T) {
		t.Parallel()

		tempDir, err := os.MkdirTemp("", "*")
		assert.NoErrorf(t, err, "could not create temp dir: %v", err)
		defer os.RemoveAll(tempDir)

		params := envs.Params{DataDir: tempDir}
		runner, err := NewGinkgoRunner(ctx, params)
		if err != nil {
			t.Fatalf("could not create runner: %v", err)
		}
		vars := make(map[string]testkube.Variable)
		variableOne := testkube.Variable{
			Name:  "GinkgoTestPackage",
			Value: "examples/e2e",
			Type_: testkube.VariableTypeBasic,
		}
		vars["GinkgoTestPackage"] = variableOne
		result, err := runner.Run(
			ctx,
			testkube.Execution{
				Content: &testkube.TestContent{
					Type_: string(testkube.TestContentTypeGitDir),
					Repository: &testkube.Repository{
						Type_:  "git",
						Uri:    repoURI,
						Branch: "main",
					},
				},
				Variables: vars,
			})

		assert.Equal(t, testkube.ExecutionStatusPassed, result.Status)
		assert.NoError(t, err)
	})

	t.Run("GinkgoRunner should run tests from a repo that fail", func(t *testing.T) {
		t.Skipf("check again is the examples/others test correct")
		t.Parallel()

		params := envs.Params{GitUsername: "testuser", GitToken: "testtoken"}
		runner, err := NewGinkgoRunner(ctx, params)
		if err != nil {
			t.Fail()
		}
		vars := make(map[string]testkube.Variable)
		variableOne := testkube.Variable{
			Name:  "GinkgoTestPackage",
			Value: "examples/other",
			Type_: testkube.VariableTypeBasic,
		}
		vars["GinkgoTestPackage"] = variableOne
		result, err := runner.Run(
			ctx,
			testkube.Execution{
				Content: &testkube.TestContent{
					Type_: string(testkube.TestContentTypeGitDir),
					Repository: &testkube.Repository{
						Type_:  "git",
						Uri:    "https://github.com/kubeshop/testkube-executor-ginkgo.git",
						Branch: "main",
					},
				},
				Variables: vars,
			})

		assert.Equal(t, testkube.ExecutionStatusFailed, result.Status)
		assert.NoError(t, err)
	})
}
