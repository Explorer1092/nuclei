package filter

import (
<<<<<<< HEAD
	"github.com/Explorer1092/nuclei/v3/pkg/catalog"
=======
<<<<<<< HEAD:v2/pkg/catalog/loader/filter/path_filter.go
	"github.com/Explorer1092/nuclei/v2/pkg/catalog"
=======
	"github.com/projectdiscovery/nuclei/v3/pkg/catalog"
>>>>>>> 419f08f61ce5ca2d3f0eae9fe36dc7c44c1f532a:pkg/catalog/loader/filter/path_filter.go
>>>>>>> projectdiscovery-main
)

// PathFilter is a path based template filter
type PathFilter struct {
	excludedTemplates          []string
	alwaysIncludedTemplatesMap map[string]struct{}
}

// PathFilterConfig contains configuration options for Path based templates Filter
type PathFilterConfig struct {
	IncludedTemplates []string
	ExcludedTemplates []string
}

// NewPathFilter creates a new path filter from provided config
func NewPathFilter(config *PathFilterConfig, catalogClient catalog.Catalog) *PathFilter {
	paths, _ := catalogClient.GetTemplatesPath(config.ExcludedTemplates)
	filter := &PathFilter{
		excludedTemplates:          paths,
		alwaysIncludedTemplatesMap: make(map[string]struct{}),
	}

	alwaysIncludeTemplates, _ := catalogClient.GetTemplatesPath(config.IncludedTemplates)
	for _, tpl := range alwaysIncludeTemplates {
		filter.alwaysIncludedTemplatesMap[tpl] = struct{}{}
	}
	return filter
}

// Match performs match for path filter on templates and returns final list
func (p *PathFilter) Match(templates []string) map[string]struct{} {
	templatesMap := make(map[string]struct{})
	for _, tpl := range templates {
		templatesMap[tpl] = struct{}{}
	}
	for _, template := range p.excludedTemplates {
		if _, ok := p.alwaysIncludedTemplatesMap[template]; ok {
			continue
		} else {
			delete(templatesMap, template)
		}
	}
	return templatesMap
}

// MatchIncluded returns true if the template was included explicitly
func (p *PathFilter) MatchIncluded(template string) bool {
	_, found := p.alwaysIncludedTemplatesMap[template]
	return found
}
