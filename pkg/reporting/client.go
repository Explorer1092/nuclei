package reporting

import (
<<<<<<< HEAD
	"github.com/Explorer1092/nuclei/v3/pkg/output"
=======
<<<<<<< HEAD:v2/pkg/reporting/client.go
	"github.com/Explorer1092/nuclei/v2/pkg/output"
=======
	"github.com/projectdiscovery/nuclei/v3/pkg/output"
>>>>>>> 419f08f61ce5ca2d3f0eae9fe36dc7c44c1f532a:pkg/reporting/client.go
>>>>>>> projectdiscovery-main
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
