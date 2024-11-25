package main

import (
	"os"

<<<<<<< HEAD
	"github.com/Explorer1092/nuclei/v3/pkg/testutils"
=======
<<<<<<< HEAD:v2/cmd/integration-test/template-dir.go
	"github.com/Explorer1092/nuclei/v2/pkg/testutils"
=======
	"github.com/projectdiscovery/nuclei/v3/pkg/testutils"
>>>>>>> 419f08f61ce5ca2d3f0eae9fe36dc7c44c1f532a:cmd/integration-test/template-dir.go
>>>>>>> projectdiscovery-main
	errorutil "github.com/projectdiscovery/utils/errors"
)

var templatesDirTestCases = []TestCaseInfo{
	{Path: "protocols/dns/cname-fingerprint.yaml", TestCase: &templateDirWithTargetTest{}},
}

type templateDirWithTargetTest struct{}

// Execute executes a test case and returns an error if occurred
func (h *templateDirWithTargetTest) Execute(filePath string) error {
	tempdir, err := os.MkdirTemp("", "nuclei-update-dir-*")
	if err != nil {
		return errorutil.NewWithErr(err).Msgf("failed to create temp dir")
	}
	defer os.RemoveAll(tempdir)

	results, err := testutils.RunNucleiTemplateAndGetResults(filePath, "8x8exch02.8x8.com", debug, "-ud", tempdir)
	if err != nil {
		return err
	}

	return expectResultsCount(results, 1)
}
