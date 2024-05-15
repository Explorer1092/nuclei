package http

import (
<<<<<<< HEAD:v2/pkg/protocols/http/cluster.go
	"github.com/Explorer1092/nuclei/v2/pkg/protocols/common/compare"
=======
	sliceutil "github.com/projectdiscovery/utils/slice"
	"golang.org/x/exp/maps"
>>>>>>> 7f556f8e33080a9eb2e52a9c14fa2e16f32f62c3:pkg/protocols/http/cluster.go
)

// CanCluster returns true if the request can be clustered.
//
// This used by the clustering engine to decide whether two requests
// are similar enough to be considered one and can be checked by
// just adding the matcher/extractors for the request and the correct IDs.
func (request *Request) CanCluster(other *Request) bool {
	if len(request.Payloads) > 0 || len(request.Fuzzing) > 0 || len(request.Raw) > 0 || len(request.Body) > 0 || request.Unsafe || request.NeedsRequestCondition() || request.Name != "" {
		return false
	}
	if request.Method != other.Method ||
		request.MaxRedirects != other.MaxRedirects ||
		request.DisableCookie != other.DisableCookie ||
		request.Redirects != other.Redirects {
		return false
	}
	if !sliceutil.Equal(request.Path, other.Path) {
		return false
	}
	if !maps.Equal(request.Headers, other.Headers) {
		return false
	}
	return true
}
