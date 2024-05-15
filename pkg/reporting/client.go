package reporting

import (
<<<<<<< HEAD:v2/pkg/reporting/client.go
	"github.com/Explorer1092/nuclei/v2/pkg/output"
=======
	"github.com/projectdiscovery/nuclei/v3/pkg/output"
>>>>>>> 7f556f8e33080a9eb2e52a9c14fa2e16f32f62c3:pkg/reporting/client.go
)

// Client is a client for nuclei issue tracking module
type Client interface {
	RegisterTracker(tracker Tracker)
	RegisterExporter(exporter Exporter)
	Close()
	Clear()
	CreateIssue(event *output.ResultEvent) error
	CloseIssue(event *output.ResultEvent) error
	GetReportingOptions() *Options
}
