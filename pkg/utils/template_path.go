package utils

import (
	"strings"

<<<<<<< HEAD
	"github.com/Explorer1092/nuclei/v3/pkg/catalog/config"
=======
	"github.com/projectdiscovery/nuclei/v3/pkg/catalog/config"
	"github.com/projectdiscovery/nuclei/v3/pkg/keys"
>>>>>>> projectdiscovery-main
)

const (
	// TemplatesRepoURL is the URL for files in nuclei-templates repository
	TemplatesRepoURL = "https://cloud.projectdiscovery.io/public/"
)

// TemplatePathURL returns the Path and URL for the provided template
func TemplatePathURL(fullPath, templateId, templateVerifier string) (path string, url string) {
	configData := config.DefaultConfig
	if configData.TemplatesDirectory != "" && strings.HasPrefix(fullPath, configData.TemplatesDirectory) {
		path = strings.TrimPrefix(strings.TrimPrefix(fullPath, configData.TemplatesDirectory), "/")
	}
	if templateVerifier == keys.PDVerifier {
		url = TemplatesRepoURL + templateId
	}
	return
}
