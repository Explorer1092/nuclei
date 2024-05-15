package colorizer

import (
	"fmt"

	"github.com/logrusorgru/aurora"

<<<<<<< HEAD:v2/internal/colorizer/colorizer.go
	"github.com/Explorer1092/nuclei/v2/pkg/model/types/severity"
=======
	"github.com/projectdiscovery/nuclei/v3/pkg/model/types/severity"
>>>>>>> 7f556f8e33080a9eb2e52a9c14fa2e16f32f62c3:internal/colorizer/colorizer.go
)

const (
	fgOrange uint8 = 208
)

func GetColor(colorizer aurora.Aurora, templateSeverity fmt.Stringer) string {
	var method func(arg interface{}) aurora.Value
	switch templateSeverity {
	case severity.Info:
		method = colorizer.Blue
	case severity.Low:
		method = colorizer.Green
	case severity.Medium:
		method = colorizer.Yellow
	case severity.High:
		method = func(stringValue interface{}) aurora.Value { return colorizer.Index(fgOrange, stringValue) }
	case severity.Critical:
		method = colorizer.Red
	default:
		method = colorizer.White
	}

	return method(templateSeverity.String()).String()
}

func New(colorizer aurora.Aurora) func(severity.Severity) string {
	return func(severity severity.Severity) string {
		return GetColor(colorizer, severity)
	}
}
