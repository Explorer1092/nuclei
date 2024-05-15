package utils

import (
	"errors"
	"io"
	"net/url"
	"strings"

<<<<<<< HEAD:v2/pkg/utils/utils.go
	"github.com/Explorer1092/nuclei/v2/pkg/catalog"
	"github.com/Explorer1092/nuclei/v2/pkg/catalog/config"
	"github.com/Explorer1092/nuclei/v2/pkg/utils/yaml"
=======
	"github.com/projectdiscovery/nuclei/v3/pkg/catalog"
>>>>>>> 7f556f8e33080a9eb2e52a9c14fa2e16f32f62c3:pkg/utils/utils.go
	"github.com/projectdiscovery/retryablehttp-go"
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

// ReadFromPathOrURL reads and returns the contents of a file or url.
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
