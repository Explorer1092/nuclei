// Package runner executes the enumeration process.
package runner

import (
	"fmt"

	"github.com/Explorer1092/nuclei/v3/pkg/catalog/config"
	"github.com/projectdiscovery/gologger"
<<<<<<< HEAD
	pdcpauth "github.com/projectdiscovery/utils/auth/pdcp"
=======
<<<<<<< HEAD:v2/internal/runner/banner.go
	"github.com/Explorer1092/nuclei/v2/pkg/catalog/config"
=======
	"github.com/projectdiscovery/nuclei/v3/pkg/catalog/config"
	pdcpauth "github.com/projectdiscovery/utils/auth/pdcp"
>>>>>>> 419f08f61ce5ca2d3f0eae9fe36dc7c44c1f532a:internal/runner/banner.go
>>>>>>> projectdiscovery-main
	updateutils "github.com/projectdiscovery/utils/update"
)

var banner = fmt.Sprintf(`
                     __     _
   ____  __  _______/ /__  (_)
  / __ \/ / / / ___/ / _ \/ /
 / / / / /_/ / /__/ /  __/ /
/_/ /_/\__,_/\___/_/\___/_/   %s
`, config.Version)

// showBanner is used to show the banner to the user
func showBanner() {
	gologger.Print().Msgf("%s\n", banner)
	gologger.Print().Msgf("\t\tprojectdiscovery.io\n\n")
}

// NucleiToolUpdateCallback updates nuclei binary/tool to latest version
func NucleiToolUpdateCallback() {
	showBanner()
	updateutils.GetUpdateToolCallback(config.BinaryName, config.Version)()
}

// AuthWithPDCP is used to authenticate with PDCP
func AuthWithPDCP() {
	showBanner()
	pdcpauth.CheckNValidateCredentials(config.BinaryName)
}
