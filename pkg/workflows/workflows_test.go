package workflows

import (
	"testing"

<<<<<<< HEAD:v2/pkg/workflows/workflows_test.go
	"github.com/Explorer1092/nuclei/v2/pkg/model/types/stringslice"
	"github.com/Explorer1092/nuclei/v2/pkg/operators"
=======
	"github.com/projectdiscovery/nuclei/v3/pkg/model/types/stringslice"
	"github.com/projectdiscovery/nuclei/v3/pkg/operators"
>>>>>>> 419f08f61ce5ca2d3f0eae9fe36dc7c44c1f532a:pkg/workflows/workflows_test.go
	"github.com/stretchr/testify/require"
)

func TestWorkflowMatchAndCompile(t *testing.T) {
	t.Run("name", func(t *testing.T) {
		matcher := &Matcher{Name: stringslice.StringSlice{Value: "sphinx"}}
		matched := matcher.Match(&operators.Result{Matches: map[string][]string{"sphinx": {}}, Extracts: map[string][]string{}})
		require.True(t, matched, "could not match value")
	})
	t.Run("name-negative", func(t *testing.T) {
		matcher := &Matcher{Name: stringslice.StringSlice{Value: "tomcat"}}
		matched := matcher.Match(&operators.Result{Matches: map[string][]string{"apache": {}}, Extracts: map[string][]string{}})
		require.False(t, matched, "could not match value")
	})
	t.Run("names-or", func(t *testing.T) {
		matcher := &Matcher{Name: stringslice.StringSlice{Value: []string{"sphinx", "elastic"}}, Condition: "or"}
		_ = matcher.Compile()
		matched := matcher.Match(&operators.Result{Matches: map[string][]string{"elastic": {}}, Extracts: map[string][]string{}})
		require.True(t, matched, "could not match value")
		matched = matcher.Match(&operators.Result{Matches: map[string][]string{"sphinx": {}}, Extracts: map[string][]string{}})
		require.True(t, matched, "could not match value")
		matched = matcher.Match(&operators.Result{Matches: map[string][]string{"random": {}}, Extracts: map[string][]string{}})
		require.False(t, matched, "could not match value")
	})
	t.Run("names-and", func(t *testing.T) {
		matcher := &Matcher{Name: stringslice.StringSlice{Value: []string{"sphinx", "elastic"}}, Condition: "and"}
		_ = matcher.Compile()
		matched := matcher.Match(&operators.Result{Matches: map[string][]string{"elastic": {}, "sphinx": {}}, Extracts: map[string][]string{}})
		require.True(t, matched, "could not match value")

		matched = matcher.Match(&operators.Result{Matches: map[string][]string{"sphinx": {}}, Extracts: map[string][]string{}})
		require.False(t, matched, "could not match value")
		matched = matcher.Match(&operators.Result{Matches: map[string][]string{"random": {}}, Extracts: map[string][]string{}})
		require.False(t, matched, "could not match value")
	})
}
