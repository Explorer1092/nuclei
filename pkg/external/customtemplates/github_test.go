package customtemplates

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/Explorer1092/nuclei/v2/pkg/catalog/config"
	"github.com/Explorer1092/nuclei/v2/pkg/testutils"
	"github.com/projectdiscovery/gologger"
<<<<<<< HEAD:v2/pkg/external/customtemplates/github_test.go
=======
	"github.com/projectdiscovery/nuclei/v3/pkg/catalog/config"
	"github.com/projectdiscovery/nuclei/v3/pkg/testutils"
>>>>>>> 7f556f8e33080a9eb2e52a9c14fa2e16f32f62c3:pkg/external/customtemplates/github_test.go
	"github.com/stretchr/testify/require"
)

func TestDownloadCustomTemplatesFromGitHub(t *testing.T) {
	gologger.DefaultLogger.SetWriter(&testutils.NoopWriter{})

	templatesDirectory, err := os.MkdirTemp("", "template-custom-*")
	require.Nil(t, err, "could not create temp directory")
	defer os.RemoveAll(templatesDirectory)

	config.DefaultConfig.SetTemplatesDir(templatesDirectory)

	options := testutils.DefaultOptions
	options.GitHubTemplateRepo = []string{"projectdiscovery/nuclei-templates", "ehsandeep/nuclei-templates"}
	options.GitHubToken = os.Getenv("GITHUB_TOKEN")

	ctm, err := NewCustomTemplatesManager(options)
	require.Nil(t, err, "could not create custom templates manager")

	ctm.Download(context.Background())

	require.DirExists(t, filepath.Join(templatesDirectory, "github", "projectdiscovery", "nuclei-templates"), "cloned directory does not exists")
	require.DirExists(t, filepath.Join(templatesDirectory, "github", "ehsandeep", "nuclei-templates"), "cloned directory does not exists")
}
