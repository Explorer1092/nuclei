package dedupe

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

<<<<<<< HEAD
	"github.com/Explorer1092/nuclei/v3/pkg/output"
=======
<<<<<<< HEAD:v2/pkg/reporting/dedupe/dedupe_test.go
	"github.com/Explorer1092/nuclei/v2/pkg/output"
=======
	"github.com/projectdiscovery/nuclei/v3/pkg/output"
>>>>>>> 419f08f61ce5ca2d3f0eae9fe36dc7c44c1f532a:pkg/reporting/dedupe/dedupe_test.go
>>>>>>> projectdiscovery-main
)

func TestDedupeDuplicates(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "nuclei")
	require.Nil(t, err, "could not create temporary storage")
	defer os.RemoveAll(tempDir)

	storage, err := New(tempDir)
	require.Nil(t, err, "could not create duplicate storage")

	tests := []*output.ResultEvent{
		{TemplateID: "test", Host: "https://example.com"},
		{TemplateID: "test", Host: "https://example.com"},
	}
	first, err := storage.Index(tests[0])
	require.Nil(t, err, "could not index item")
	require.True(t, first, "could not index valid item")

	second, err := storage.Index(tests[1])
	require.Nil(t, err, "could not index item")
	require.False(t, second, "could index duplicate item")
}
