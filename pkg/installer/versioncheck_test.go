package installer

import (
	"testing"

<<<<<<< HEAD:v2/internal/installer/versioncheck_test.go
	"github.com/Explorer1092/nuclei/v2/pkg/catalog/config"
=======
	"github.com/projectdiscovery/nuclei/v3/pkg/catalog/config"
>>>>>>> 7f556f8e33080a9eb2e52a9c14fa2e16f32f62c3:pkg/installer/versioncheck_test.go
	"github.com/projectdiscovery/utils/generic"
	"github.com/stretchr/testify/require"
)

func TestVersionCheck(t *testing.T) {
	err := NucleiVersionCheck()
	require.Nil(t, err)
	cfg := config.DefaultConfig
	if generic.EqualsAny("", cfg.LatestNucleiIgnoreHash, cfg.LatestNucleiVersion, cfg.LatestNucleiTemplatesVersion) {
		// all above values cannot be empty
		t.Errorf("something went wrong got empty response nuclei-version=%v templates-version=%v ignore-hash=%v", cfg.LatestNucleiVersion, cfg.LatestNucleiTemplatesVersion, cfg.LatestNucleiIgnoreHash)
	}
}
