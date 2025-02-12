package scraper_test

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/kubeshop/testkube/pkg/utils/test"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/kubeshop/testkube/pkg/executor/scraper"
	"github.com/kubeshop/testkube/pkg/filesystem"
)

func TestFilesystemExtractor_Extract_Integration(t *testing.T) {
	test.IntegrationTest(t)
	t.Parallel()

	tempDir, err := os.MkdirTemp("", "test")
	require.NoError(t, err)

	defer os.RemoveAll(tempDir)

	err = os.Mkdir(filepath.Join(tempDir, "subdir"), os.ModePerm)
	require.NoError(t, err)

	file1 := filepath.Join(tempDir, "file1.txt")
	file2 := filepath.Join(tempDir, "file2.txt")
	file3 := filepath.Join(tempDir, "subdir", "file3.txt")

	err = os.WriteFile(file1, []byte("test1"), os.ModePerm)
	assert.NoError(t, err)

	err = os.WriteFile(file2, []byte("test2"), os.ModePerm)
	assert.NoError(t, err)

	err = os.WriteFile(file3, []byte("test3"), os.ModePerm)
	assert.NoError(t, err)

	processCallCount := 0
	processFn := func(ctx context.Context, object *scraper.Object) error {
		processCallCount++
		b, err := io.ReadAll(object.Data)
		if err != nil {
			t.Fatalf("error reading %s: %v", object.Name, err)
		}
		switch object.Name {
		case "file1.txt":

			assert.Equal(t, b, []byte("test1"))
		case "file2.txt":
			assert.Equal(t, b, []byte("test2"))
		case "subdir/file3.txt":
			assert.Equal(t, b, []byte("test3"))
		default:
			t.Fatalf("unexpected file: %s", object.Name)
		}

		return nil
	}

	extractor := scraper.NewRecursiveFilesystemExtractor(filesystem.NewOSFileSystem())
	err = extractor.Extract(context.Background(), []string{tempDir}, processFn)
	require.NoError(t, err)
	assert.Equal(t, processCallCount, 3)
}

func TestFilesystemExtractor_Extract_RelPath_Integration(t *testing.T) {
	test.IntegrationTest(t)
	t.Parallel()

	tempDir, err := os.MkdirTemp("", "test")
	require.NoError(t, err)

	defer os.RemoveAll(tempDir)

	err = os.Mkdir(filepath.Join(tempDir, "subdir"), os.ModePerm)
	require.NoError(t, err)

	file1 := filepath.Join(tempDir, "file1.txt")

	err = os.WriteFile(file1, []byte("test1"), os.ModePerm)
	assert.NoError(t, err)

	processCallCount := 0
	processFn := func(ctx context.Context, object *scraper.Object) error {
		processCallCount++
		b, err := io.ReadAll(object.Data)
		if err != nil {
			t.Fatalf("error reading %s: %v", object.Name, err)
		}
		assert.Equal(t, b, []byte("test1"))

		return nil
	}

	extractor := scraper.NewRecursiveFilesystemExtractor(filesystem.NewOSFileSystem())
	scrapeDirs := []string{filepath.Join(tempDir, "file1.txt")}
	err = extractor.Extract(context.Background(), scrapeDirs, processFn)
	require.NoError(t, err)
	assert.Equal(t, processCallCount, 1)
}
