package utils

import (
	"errors"
	"fmt"
	"io"
	"net/url"
	"strings"

<<<<<<< HEAD
	"github.com/Explorer1092/nuclei/v3/pkg/catalog"
=======
<<<<<<< HEAD:v2/pkg/utils/utils.go
	"github.com/Explorer1092/nuclei/v2/pkg/catalog"
	"github.com/Explorer1092/nuclei/v2/pkg/catalog/config"
	"github.com/Explorer1092/nuclei/v2/pkg/utils/yaml"
=======
	"github.com/cespare/xxhash"
	"github.com/projectdiscovery/nuclei/v3/pkg/catalog"
>>>>>>> 419f08f61ce5ca2d3f0eae9fe36dc7c44c1f532a:pkg/utils/utils.go
>>>>>>> projectdiscovery-main
	"github.com/projectdiscovery/retryablehttp-go"
	mapsutil "github.com/projectdiscovery/utils/maps"
	"golang.org/x/exp/constraints"
)

func IsBlank(value string) bool {
	return strings.TrimSpace(value) == ""
}

func UnwrapError(err error) error {
	for { // get the last wrapped error
		unwrapped := errors.Unwrap(err)
		if unwrapped == nil {
			break
		}
		err = unwrapped
	}
	return err
}

// IsURL tests a string to determine if it is a well-structured url or not.
func IsURL(input string) bool {
	u, err := url.Parse(input)
	return err == nil && u.Scheme != "" && u.Host != ""
}

// ReaderFromPathOrURL reads and returns the contents of a file or url.
func ReaderFromPathOrURL(templatePath string, catalog catalog.Catalog) (io.ReadCloser, error) {
	if IsURL(templatePath) {
		resp, err := retryablehttp.DefaultClient().Get(templatePath)
		if err != nil {
			return nil, err
		}
		return resp.Body, nil
	} else {
		f, err := catalog.OpenFile(templatePath)
		if err != nil {
			return nil, err
		}
		return f, nil
	}
}

// StringSliceContains checks if a string slice contains a string.
func StringSliceContains(slice []string, item string) bool {
	for _, i := range slice {
		if strings.EqualFold(i, item) {
			return true
		}
	}
	return false
}

// MapHash generates a hash for any give map
func MapHash[K constraints.Ordered, V any](m map[K]V) uint64 {
	keys := mapsutil.GetSortedKeys(m)
	var sb strings.Builder
	for _, k := range keys {
		sb.WriteString(fmt.Sprintf("%v:%v\n", k, m[k]))
	}
	return xxhash.Sum64([]byte(sb.String()))
}
