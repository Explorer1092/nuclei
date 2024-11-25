package format

import (
	"strings"
	"testing"

<<<<<<< HEAD
	"github.com/Explorer1092/nuclei/v3/pkg/model"
	"github.com/Explorer1092/nuclei/v3/pkg/model/types/severity"
	"github.com/Explorer1092/nuclei/v3/pkg/model/types/stringslice"
	"github.com/Explorer1092/nuclei/v3/pkg/reporting/exporters/markdown/util"
	"github.com/stretchr/testify/require"
=======
<<<<<<< HEAD:v2/pkg/reporting/format/format_utils_test.go
<<<<<<< HEAD:v2/pkg/reporting/format/format_test.go
	"github.com/stretchr/testify/assert"

	"github.com/Explorer1092/nuclei/v2/pkg/model"
	"github.com/Explorer1092/nuclei/v2/pkg/model/types/severity"
	"github.com/Explorer1092/nuclei/v2/pkg/model/types/stringslice"
=======
	"github.com/projectdiscovery/nuclei/v2/pkg/model"
	"github.com/projectdiscovery/nuclei/v2/pkg/model/types/severity"
	"github.com/projectdiscovery/nuclei/v2/pkg/model/types/stringslice"
	"github.com/projectdiscovery/nuclei/v2/pkg/reporting/exporters/markdown/util"
>>>>>>> bb98eced070f4ae137b8cd2a7f887611bc1b9c93:v2/pkg/reporting/format/format_utils_test.go
=======
	"github.com/projectdiscovery/nuclei/v3/pkg/model"
	"github.com/projectdiscovery/nuclei/v3/pkg/model/types/severity"
	"github.com/projectdiscovery/nuclei/v3/pkg/model/types/stringslice"
	"github.com/projectdiscovery/nuclei/v3/pkg/reporting/exporters/markdown/util"
	"github.com/stretchr/testify/require"
>>>>>>> 419f08f61ce5ca2d3f0eae9fe36dc7c44c1f532a:pkg/reporting/format/format_utils_test.go
>>>>>>> projectdiscovery-main
)

func TestToMarkdownTableString(t *testing.T) {
	info := model.Info{
		Name:           "Test Template Name",
		Authors:        stringslice.StringSlice{Value: []string{"forgedhallpass", "ice3man"}},
		Description:    "Test description",
		SeverityHolder: severity.Holder{Severity: severity.High},
		Tags:           stringslice.StringSlice{Value: []string{"cve", "misc"}},
		Reference:      stringslice.NewRawStringSlice("reference1"),
		Metadata: map[string]interface{}{
			"customDynamicKey1": "customDynamicValue1",
			"customDynamicKey2": "customDynamicValue2",
		},
	}

	result := CreateTemplateInfoTable(&info, &util.MarkdownFormatter{})

	expectedOrderedAttributes := `| Key | Value |
| --- | --- |
| Name | Test Template Name |
| Authors | forgedhallpass, ice3man |
| Tags | cve, misc |
| Severity | high |
| Description | Test description |`

	expectedDynamicAttributes := []string{
		"| customDynamicKey1 | customDynamicValue1 |",
		"| customDynamicKey2 | customDynamicValue2 |",
		"", // the expected result ends in a new line (\n)
	}

	actualAttributeSlice := strings.Split(result, "\n")
	dynamicAttributeIndex := len(actualAttributeSlice) - len(expectedDynamicAttributes)
	require.Equal(t, strings.Split(expectedOrderedAttributes, "\n"), actualAttributeSlice[:dynamicAttributeIndex]) // the first part of the result is ordered
	require.ElementsMatch(t, expectedDynamicAttributes, actualAttributeSlice[dynamicAttributeIndex:])              // dynamic parameters are not ordered
}
