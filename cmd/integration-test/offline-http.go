package main

import (
	"fmt"

<<<<<<< HEAD
	"github.com/Explorer1092/nuclei/v3/pkg/testutils"
=======
<<<<<<< HEAD:v2/cmd/integration-test/offline-http.go
	"github.com/Explorer1092/nuclei/v2/pkg/testutils"
=======
	"github.com/projectdiscovery/nuclei/v3/pkg/testutils"
>>>>>>> 419f08f61ce5ca2d3f0eae9fe36dc7c44c1f532a:cmd/integration-test/offline-http.go
>>>>>>> projectdiscovery-main
)

var offlineHttpTestcases = []TestCaseInfo{
	{Path: "protocols/offlinehttp/rfc-req-resp.yaml", TestCase: &RfcRequestResponse{}},
	{Path: "protocols/offlinehttp/offline-allowed-paths.yaml", TestCase: &RequestResponseWithAllowedPaths{}},
	{Path: "protocols/offlinehttp/offline-raw.yaml", TestCase: &RawRequestResponse{}},
}

type RfcRequestResponse struct{}

// Execute executes a test case and returns an error if occurred
func (h *RfcRequestResponse) Execute(filePath string) error {
	results, err := testutils.RunNucleiTemplateAndGetResults(filePath, "protocols/offlinehttp/data/", debug, "-passive")
	if err != nil {
		return err
	}

	return expectResultsCount(results, 1)
}

type RequestResponseWithAllowedPaths struct{}

// Execute executes a test case and returns an error if occurred
func (h *RequestResponseWithAllowedPaths) Execute(filePath string) error {
	results, err := testutils.RunNucleiTemplateAndGetResults(filePath, "protocols/offlinehttp/data/", debug, "-passive")
	if err != nil {
		return err
	}

	return expectResultsCount(results, 1)
}

type RawRequestResponse struct{}

// Execute executes a test case and returns an error if occurred
func (h *RawRequestResponse) Execute(filePath string) error {
	_, err := testutils.RunNucleiTemplateAndGetResults(filePath, "protocols/offlinehttp/data/", debug, "-passive")
	if err == nil {
		return fmt.Errorf("incorrect result: no error (actual) vs error expected")
	}
	return nil
}
