package utils

import (
	"strings"

<<<<<<< HEAD:v2/pkg/utils/template_path.go
	"github.com/Explorer1092/nuclei/v2/pkg/catalog/config"
=======
	"github.com/projectdiscovery/nuclei/v3/pkg/catalog/config"
>>>>>>> 7f556f8e33080a9eb2e52a9c14fa2e16f32f62c3:pkg/utils/template_path.go
)

const (
	// TemplatesRepoURL is the URL for files in nuclei-templates repository
<<<<<<< HEAD:v2/pkg/utils/template_path.go
	TemplatesRepoURL = "https://github.com/Explorer1092/nuclei-templates/blob/main/"
=======
	TemplatesRepoURL = "https://cloud.projectdiscovery.io/public/"
>>>>>>> 7f556f8e33080a9eb2e52a9c14fa2e16f32f62c3:pkg/utils/template_path.go
)

// TemplatePathURL returns the Path and URL for the provided template
func TemplatePathURL(fullPath, templateId string) (string, string) {
	var templateDirectory string
	configData := config.DefaultConfig
	if configData.TemplatesDirectory != "" && strings.HasPrefix(fullPath, configData.TemplatesDirectory) {
		templateDirectory = configData.TemplatesDirectory
	} else {
		return "", ""
	}

	finalPath := strings.TrimPrefix(strings.TrimPrefix(fullPath, templateDirectory), "/")
	templateURL := TemplatesRepoURL + templateId
	return finalPath, templateURL
}
